package main

import (
	"botory_project/internal/config" // yaml파일에서 Botconfig구조체로 변환시켜주는 함수
	"fmt"
	"log"
)

func main() {
	// 설정 파일 로드
	botConfig, err := config.Load("configs/chatbot.yaml")
	if err != nil {
		log.Fatalf("설정 파일을 불러오는 데 실패했습니다: %v", err)
	}

	// 파싱된 결과 출력
	fmt.Println("챗봇 이름:", botConfig.BotName)
	fmt.Println("기본 응답:", botConfig.DefaultResponse)
	fmt.Println("--- 대화 목록 ---")
	for _, dialog := range botConfig.Dialogs {
		fmt.Printf("키워드: '%s' -> 응답: '%s'\n", dialog.Keyword, dialog.Response)
	}
}
