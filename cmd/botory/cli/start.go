package cli

import (
	"botory_project/internal/adapter"
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
		botConfig, err := config.Load("botory_project/configs/chatbot.yaml")
		if err != nil {
			log.Fatalf("설정 파일을 불러오는 데 실패했습니다: %v", err)
		}

		// 2. 외부 통신을 위한 어댑터 생성
		apiClient := adapter.NewHTTPClient()

		// 3. 챗봇 엔진 생성 시, 어댑터를 주입
		engine := app.NewBotEngine(botConfig, apiClient)

		// 4. HTTP 핸들러 등록
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

		// 에러를 더 이상 무시하지 않고 'err' 변수에 받습니다.
		response, err := engine.GetResponses(req.Message)
		if err != nil {
			// 만약 에러가 발생했다면, 터미널(서버 창)에 에러 내용을 출력합니다.
			log.Printf("!!! 응답 생성 중 에러 발생: %v", err)
			// 사용자에게는 간단한 에러 메시지를 보냅니다.
			http.Error(w, "챗봇 응답 생성에 실패했습니다.", http.StatusInternalServerError)
			return
		}

		resp := ChatResponse{
			Response: response,
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "응답 생성에 실패했습니다.", http.StatusInternalServerError)
		}
	}
}
