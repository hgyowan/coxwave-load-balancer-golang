package service

import (
	"container/heap"
	"coxwave-load-balancer-golang/domain"
	"coxwave-load-balancer-golang/domain/load_balancer"
	"errors"
	"sync"
)

type loadBalancer struct {
	nodes load_balancer.NodeHeap
	mu    sync.Mutex
}

func (l *loadBalancer) Route(r *load_balancer.Request) (*load_balancer.Node, error){
	l.mu.Lock() // NodeHeap 동시 접근 제어
	defer l.mu.Unlock()

	// 요청량이 가장 적은 노드중 요청을 수용할 수 있는 노드를 찾는다
	for len(l.nodes) > 0 {
		node := heap.Pop(&l.nodes).(*load_balancer.Node)
		node.Mu.Lock() // 각 노드의 Rate Limit 값 동시 접근 제어
		if node.HealthCheck(r) {
			node.Work(r)
			node.Mu.Unlock()
			heap.Push(&l.nodes, node)
			return node, nil
		}
		node.Mu.Unlock()
	}

	return nil, errors.New("there is no available node to handle the request")
}

func NewLoadBalancer(nodes load_balancer.NodeHeap) domain.LoadBalancerService {
	return &loadBalancer{
		nodes: nodes,
	}
}