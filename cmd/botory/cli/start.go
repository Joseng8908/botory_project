package cli

import (
	"botory_project/internal/app"
	"botory_project/internal/config"
	"encoding/json"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

// --- API 서버를 위한 요청/응답 구조체 정의 ---

type ChatRequest struct {
	Message string `json:"message"`
}

type ChatResponse struct {
	Response string `json:"response"`
}

// --- 'start' 명령어 정의 ---

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "YAML 설정 파일을 기반으로 챗봇 서버를 시작합니다.",
	Run: func(cmd *cobra.Command, args []string) {
		// 1. 설정 파일 로드
		// 참고: 나중에는 이 경로를 플래그로 받아 동적으로 처리할 수 있습니다.
		botConfig, err := config.Load("configs/chatbot.yaml")
		if err != nil {
			log.Fatalf("설정 파일을 불러오는 데 실패했습니다: %v", err)
		}

		// 2. 챗봇 엔진 생성
		engine := app.NewBotEngine(botConfig)

		// 3. HTTP 핸들러 등록
		http.HandleFunc("/chat", chatHandler(engine))

		// 4. 웹 서버 시작
		log.Println("챗봇 서버 시작. http://localhost:8080/chat 에서 POST 요청을 기다립니다.")
		log.Printf("'%s' 챗봇이 당신을 기다립니다.", botConfig.BotName)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatalf("서버 시작에 실패했습니다: %v", err)
		}
	},
}

// --- HTTP 핸들러 함수 ---

func chatHandler(engine *app.BotEngine) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "POST 요청만 지원합니다.", http.StatusMethodNotAllowed)
			return
		}

		var req ChatRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "잘못된 요청 형식입니다.", http.StatusBadRequest)
			return
		}

		// 엔진을 사용해 응답 생성 (GetResponses 함수 사용)
		response, _ := engine.GetResponses(req.Message)

		resp := ChatResponse{
			Response: response,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "응답 생성에 실패했습니다.", http.StatusInternalServerError)
		}
	}
}
