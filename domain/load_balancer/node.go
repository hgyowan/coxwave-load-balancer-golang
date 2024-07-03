package load_balancer

import (
	"sync"
	"time"
)

type NodeHeap []*Node

type Node struct {
	// 노드 주소
	Address string
	// 분당 HTTP 바이트 수 제한
	BPM int32
	// 현 노드에서 사용한 총 바이트량
	UsedByteSize int32
	// 분당 요청 수 제한
	RPM int32
	// 현 노드에 들어온 총 요청 수량
	RequestCounts int32
	// 해당 노드 Rate Limit 초기화 시간 (분당 1번)
	ResetTime time.Time
	// 노드가 Heap 에 위치하고 있는 index
	Index int
	Mu    sync.Mutex
}
