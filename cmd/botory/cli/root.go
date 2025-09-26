package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// rootCmd는 'botory'라는 기본 명령어를 나타냅니다.
var rootCmd = &cobra.Command{
	Use:   "botory",
	Short: "Botory는 YAML 설정으로 챗봇을 만드는 도구입니다.",
	Long: `Botory는 개발자가 아닌 사용자가 간단한 YAML 설정 파일만으로
쉽게 자신만의 챗봇을 만들어 API 서버로 실행할 수 있도록 돕는 CLI 도구입니다.`,
}

// Execute는 모든 CLI 명령어의 진입점 역할을 합니다. main.go에서 호출됩니다.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "명령어 실행 중 오류가 발생했습니다: '%s'", err)
		os.Exit(1)
	}
}

// init 함수는 패키지가 로드될 때 실행됩니다.
// startCmd를 rootCmd의 하위 명령어로 등록합니다.
func init() {
	rootCmd.AddCommand(startCmd)
}
