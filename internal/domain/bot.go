package domain

// BotConfig는 chatbot.yaml 파일의 전체 구조를 나타냅니다.
type BotConfig struct {
	BotName         string   `yaml:"botName"`
	DefaultResponse string   `yaml:"defaultResponse"`
	Dialogs         []Dialog `yaml:"dialogs"`
}

// Dialog는 하나의 키워드와 응답 쌍을 나타냅니다.
type Dialog struct {
	Keyword  string `yaml:"keyword"`
	Response string `yaml:"response"`
}
