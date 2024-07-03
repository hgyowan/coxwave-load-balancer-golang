package service

import (
	"coxwave-load-balancer-golang/domain"
	"coxwave-load-balancer-golang/domain/load_balancer"
	"sync"
)

type loadBalancer struct {
	nodes load_balancer.NodeHeap
	mu    sync.Mutex
}

func (l *loadBalancer) Route(r *load_balancer.Request) (*load_balancer.Node, error){
	panic("implement me")
}

func NewLoadBalancer(nodes load_balancer.NodeHeap) domain.LoadBalancerService {
	return &loadBalancer{
		nodes: nodes,
	}
}