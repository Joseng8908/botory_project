package config

import (
	"botory_project/internal/domain"
	"os"

	"gopkg.in/yaml.v3"
)

// Load는 지정된 경로의 YAML 파일을 읽어 BotConfig 구조체로 변환합니다.
func Load(path string) (*domain.BotConfig, error) {
	// 파일을 읽습니다.
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// 읽어온 데이터를 BotConfig 구조체에 파싱(언마셜링)합니다.
	var config domain.BotConfig
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
