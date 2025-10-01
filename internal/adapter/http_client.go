package adapter

import (
	"io"
	"net/http"
)

// HTTPClient는 net/http를 사용하여 APIClient 인터페이스를 구현합니다.
type HTTPClient struct{}

// NewHTTPClient는 새로운 HTTPClient를 생성합니다.
func NewHTTPClient() *HTTPClient {
	return &HTTPClient{}
}

// Fetch는 주어진 URL로 GET 요청을 보내고 응답 본문을 문자열로 반환합니다.
func (c *HTTPClient) Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
