package domain

import "coxwave-load-balancer-golang/domain/load_balancer"

type LoadBalancerService interface {
	Route(r *load_balancer.Request)
}