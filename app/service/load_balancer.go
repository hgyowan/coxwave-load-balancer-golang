package service

import (
	"coxwave-load-balancer-golang/domain"
	"coxwave-load-balancer-golang/domain/load_balancer"
	"sync"
)

type loadBalancer struct {
	nodes load_balancer.Nodes
	mu    sync.Mutex
}

func (l *loadBalancer) Route(r *load_balancer.Request) {
	panic("implement me")
}

func NewLoadBalancer(nodes load_balancer.Nodes) domain.LoadBalancerService {
	return &loadBalancer{
		nodes: nodes,
	}
}