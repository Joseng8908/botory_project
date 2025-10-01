package app

import (
	"botory_project/internal/domain"
	"strings"
)

type BotEngine struct {
	config    *domain.BotConfig
	apiClient domain.APIClient
}

// 새로운 BotEngine을 만드는 생성자
func NewBotEngine(config *domain.BotConfig, apiClient domain.APIClient) *BotEngine {
	return &BotEngine{
		config:    config,
		apiClient: apiClient,
	}
}

// 사용자 메세지를 받고, 그 메세지에서 적절한 키워드를 찾으면 그에 맞는 응답 반환
func (e *BotEngine) GetResponses(message string) (string, error) {
	for _, dialog := range e.config.Dialogs {

		// 1단계: 먼저 메시지가 현재 dialog에 해당하는지 '매치 여부'를 판단합니다.
		isMatch := false
		if dialog.MatchType == "contains" {
			if strings.Contains(message, dialog.Keyword) {
				isMatch = true
			}
		} else { // 'exact' 또는 matchType이 생략된 경우
			if dialog.Keyword == message {
				isMatch = true
			}
		}

		// 2단계: 만약 매치되는 dialog를 찾았다면, 어떤 응답을 할지 결정합니다.
		if isMatch {
			// 이 dialog가 API 호출용인지 확인합니다.
			if dialog.ApiCallURL != "" {
				// API 호출용이라면, apiClient를 사용해 결과를 가져옵니다.
				return e.apiClient.Fetch(dialog.ApiCallURL)
			}

			// API 호출용이 아니라면, 단순 텍스트 응답을 반환합니다.
			return dialog.Response, nil
		}
	}

	// 3단계: for문이 끝날 때까지 어떤 dialog와도 매치되지 않았다면, 기본 응답을 반환합니다.
	return e.config.DefaultResponse, nil
}
