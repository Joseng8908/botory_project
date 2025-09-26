package app

import (
	"botory_project/internal/domain"
	"strings"
)

type BotEngine struct {
	config *domain.BotConfig
}

// 새로운 BotEngine을 만드는 생성자
func NewBotEngine(config *domain.BotConfig) *BotEngine {
	return &BotEngine{config: config}
}

// 사용자 메세지를 받고, 그 메세지에서 적절한 키워드를 찾으면 그에 맞는 응답 반환
func (e *BotEngine) GetResponses(message string) (string, error) {
	for _, dialog := range e.config.Dialogs {
		if dialog.MatchType == "contains" { // 문장에 단어가 포함되어 있을때
			if strings.Contains(message, dialog.Keyword) {
				return dialog.Response, nil
			}
		} else { // 기본값은 'exact'
			if dialog.Keyword == message { // 완전히 일치할 때
				return dialog.Response, nil
			}
		}
	}
	return e.config.DefaultResponse, nil
}
