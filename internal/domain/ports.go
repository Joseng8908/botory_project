package domain

// APIClient는 외부 API를 호출하는 기능에 대한 명세입니다.
type APIClient interface {
	Fetch(url string) (string, error)
}
