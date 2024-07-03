package load_balancer

import (
	"sync"
	"time"
)

// NodeHeap
// 현재 가장 적은 요청을 수행 하고 있는 Node 를 찾기 위한 min heap 구현
// https://go.dev/src/container/heap/example_pq_test.go 참고 하여 heap interface 구현
type NodeHeap []*Node

func (nh NodeHeap) Len() int { return len(nh) }

// Less
// 적은 요청량 일수록 우선순위 높음
func (nh NodeHeap) Less(i, j int) bool {
	return nh[i].RequestCounts < nh[j].RequestCounts
}

// Swap
// 값, 인덱스 순서 변경
func (nh NodeHeap) Swap(i, j int) {
	nh[i], nh[j] = nh[j], nh[i]
	nh[i].Index = i
	nh[j].Index = j
}

// Push
// 기존 heap 맨 뒤에 새로운 값 추가
// 이후 내부적으로 up 함수 수행
func (nh *NodeHeap) Push(x interface{}) {
	n := len(*nh)
	item := x.(*Node)
	item.Index = n
	*nh = append(*nh, item)
}

// Pop
// heap 의 마지막 값 추출 후 heap 길이 감소
// 이후 내부적으로 down 함수 수행
func (nh *NodeHeap) Pop() interface{} {
	old := *nh
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.Index = -1
	*nh = old[0 : n-1]
	return item
}

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
