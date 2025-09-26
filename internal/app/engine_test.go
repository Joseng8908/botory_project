package app

import (
	"botory_project/internal/domain"
	"testing"
)

// TestBotEngine_GetResponses는 BotEngine의 GetResponses 메서드를 테스트합니다.
func TestBotEngine_GetResponses(t *testing.T) {
	// --- 테스트 준비 (Arrange) ---
	// 실제 YAML 파일을 읽는 대신, 테스트용 가짜(mock) 설정 데이터를 직접 만듭니다.
	// 이렇게 하면 테스트가 외부 파일에 의존하지 않아 더 빠르고 안정적이 됩니다.
	mockConfig := &domain.BotConfig{
		BotName:         "테스트봇",
		DefaultResponse: "이해하지 못했어요.",
		Dialogs: []domain.Dialog{
			{Keyword: "테스트", Response: "성공", MatchType: "exact"},
			{Keyword: "안녕", Response: "반가워요", MatchType: "exact"},
			{Keyword: "가격", Response: "~$입니다", MatchType: "contains"},
		},
	}

	// 가짜 설정 데이터로 챗봇 엔진을 생성합니다.
	engine := NewBotEngine(mockConfig)

	// --- 테스트 케이스 정의 (Table-Driven Test) ---
	// 여러 시나리오를 한 번에 테스트하기 위해 테스트 케이스들을 구조체 슬라이스로 정의합니다.
	testCases := []struct {
		name             string // 테스트 케이스 이름
		inputMessage     string // 입력 메시지
		expectedResponse string // 기대되는 응답
	}{
		{
			name:             "일치하는 키워드가 있을 경우",
			inputMessage:     "안녕",
			expectedResponse: "반가워요",
		},
		{
			name:             "일치하는 다른 키워드가 있을 경우",
			inputMessage:     "테스트",
			expectedResponse: "성공",
		},
		{
			name:             "일치하는 키워드가 없을 경우 (기본 응답)",
			inputMessage:     "이건 없는 키워드야",
			expectedResponse: "이해하지 못했어요.",
		},
		{
			name:             "빈 문자열이 입력될 경우 (기본 응답)",
			inputMessage:     "",
			expectedResponse: "이해하지 못했어요.",
		},
		{
			name:             "키워드가 포함되어있을 경우",
			inputMessage:     "가격",
			expectedResponse: "~$입니다",
		},
	}

	// --- 테스트 실행 (Act & Assert) ---
	for _, tc := range testCases {
		// t.Run을 사용하면 각 테스트 케이스를 독립적인 서브테스트로 실행하고 결과를 명확하게 볼 수 있습니다.
		t.Run(tc.name, func(t *testing.T) {
			// 엔진에 메시지를 보내 실제 응답을 받습니다.
			// GetResponses는 (string, error)를 반환하므로, 지금은 에러를 무시합니다.
			actualResponse, _ := engine.GetResponses(tc.inputMessage)

			// 실제 응답과 기대한 응답이 다른 경우, 테스트를 실패 처리하고 에러 메시지를 출력합니다.
			if actualResponse != tc.expectedResponse {
				t.Errorf("기대값: '%s', 실제값: '%s'", tc.expectedResponse, actualResponse)
			}
		})
	}
}
