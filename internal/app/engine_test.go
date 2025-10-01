package app

import (
	"botory_project/internal/domain"
	"errors"
	"testing"
)

// --- 1. 테스트를 위한 '가짜 API 클라이언트' 만들기 ---

// MockAPIClient는 테스트에서 실제 HTTP 요청을 보내는 대신,
// 우리가 원하는 결과를 반환하도록 조종할 수 있는 가짜 클라이언트입니다.
// domain.APIClient 인터페이스를 똑같이 구현합니다.
type MockAPIClient struct {
	// 이 함수 변수에 우리가 원하는 동작을 테스트 케이스마다 바꿔서 넣어줄 것입니다.
	FetchFunc func(url string) (string, error)
}

// Fetch 메서드는 MockAPIClient가 domain.APIClient 인터페이스를 만족시키기 위해 필요합니다.
// 실제 동작은 FetchFunc에 위임합니다.
func (m *MockAPIClient) Fetch(url string) (string, error) {
	if m.FetchFunc != nil {
		return m.FetchFunc(url)
	}
	return "", errors.New("FetchFunc가 설정되지 않았습니다")
}

// --- 2. 실제 테스트 코드 작성 ---

func TestBotEngine_GetResponses(t *testing.T) {
	// --- 테스트 준비 (Arrange) ---

	// 가짜 API 클라이언트 인스턴스를 생성합니다.
	mockAPIClient := &MockAPIClient{}

	// 테스트에 사용할 설정 데이터를 정의합니다. API 호출용 dialog를 추가합니다.
	mockConfig := &domain.BotConfig{
		BotName:         "테스트봇",
		DefaultResponse: "이해하지 못했어요.",
		Dialogs: []domain.Dialog{
			{Keyword: "안녕", Response: "반가워요"}, // MatchType 생략 (기본값 exact)
			{Keyword: "가격", Response: "가격 문의는 1588-XXXX 입니다.", MatchType: "contains"},
			{Keyword: "농담", ApiCallURL: "https://api.example.com/joke"},                           // API 호출용
			{Keyword: "날씨", ApiCallURL: "https://api.example.com/weather", MatchType: "contains"}, // contains + API 호출
		},
	}

	// 이제 NewBotEngine에 가짜 API 클라이언트를 함께 전달합니다.
	engine := NewBotEngine(mockConfig, mockAPIClient)

	// --- 테스트 케이스 정의 ---
	testCases := []struct {
		name             string
		inputMessage     string
		expectedResponse string
		mockAPIResponse  string // 이 케이스에서 가짜 API가 어떤 응답을 할지 정의
		expectAPICall    bool   // 이 케이스에서 API 호출이 예상되는지 여부
	}{
		{
			name:             "정확히 일치하는 경우 (단순 응답)",
			inputMessage:     "안녕",
			expectedResponse: "반가워요",
			expectAPICall:    false,
		},
		{
			name:             "포함하는 경우 (단순 응답)",
			inputMessage:     "상품 가격이 얼마인가요?",
			expectedResponse: "가격 문의는 1588-XXXX 입니다.",
			expectAPICall:    false,
		},
		{
			name:             "정확히 일치하는 경우 (API 호출 성공)",
			inputMessage:     "농담",
			expectedResponse: "개발자가 가장 좋아하는 소금은? 'Ctrl+S' 소금!",
			mockAPIResponse:  "개발자가 가장 좋아하는 소금은? 'Ctrl+S' 소금!",
			expectAPICall:    true,
		},
		{
			name:             "포함하는 경우 (API 호출 성공)",
			inputMessage:     "오늘 서울 날씨 알려줘",
			expectedResponse: "오늘 서울은 맑습니다.",
			mockAPIResponse:  "오늘 서울은 맑습니다.",
			expectAPICall:    true,
		},
		{
			name:             "일치하는 키워드가 없을 경우 (기본 응답)",
			inputMessage:     "이건 없는 키워드야",
			expectedResponse: "이해하지 못했어요.",
			expectAPICall:    false,
		},
	}

	// --- 테스트 실행 ---
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// 각 테스트 케이스를 실행하기 전에, 가짜 API의 동작을 설정합니다.
			if tc.expectAPICall {
				mockAPIClient.FetchFunc = func(url string) (string, error) {
					// API 호출이 예상되면, 미리 정의된 mockAPIResponse를 반환하도록 설정
					return tc.mockAPIResponse, nil
				}
			}

			// 엔진의 GetResponses 메서드를 실행합니다.
			actualResponse, _ := engine.GetResponses(tc.inputMessage)

			// 실제 응답과 기대한 응답을 비교합니다.
			if actualResponse != tc.expectedResponse {
				t.Errorf("기대값: '%s', 실제값: '%s'", tc.expectedResponse, actualResponse)
			}
		})
	}
}
