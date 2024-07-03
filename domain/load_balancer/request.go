package load_balancer

type Request struct {
	// 요청 HTTP Body 크기라고 가정
	BodySize int32 `json:"body_size"`
}