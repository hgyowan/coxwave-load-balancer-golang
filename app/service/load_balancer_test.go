package service

import (
	"coxwave-load-balancer-golang/domain"
	"coxwave-load-balancer-golang/domain/load_balancer"
	"github.com/stretchr/testify/require"
	"sync"
	"testing"
)

type testModel struct {
	nodes load_balancer.NodeHeap
	svc domain.LoadBalancerService
}

func beforeEach(t *testing.T) *testModel {
	nodes := load_balancer.NodeHeap{
		{
			Address:       "node1",
			BPM:           100,
			RPM:           10,
		},
		{
			Address:       "node2",
			BPM:           200,
			RPM:           20,
		},
	}

	return &testModel{svc: NewLoadBalancer(nodes), nodes: nodes}
}

func TestService_Route(t *testing.T) {
	type testCase struct {
		name string
		requests []*load_balancer.Request
		result string
	}

	testCases := []*testCase{
		{
			name:     "default",
			requests: []*load_balancer.Request{
				{
					BodySize: 10,
				},
				{
					BodySize: 10,
				},
				{
					BodySize: 10,
				},
			},
		},
		{
			name:     "health check error (size)",
			requests: []*load_balancer.Request{
				{
					BodySize: 100,
				},
				{
					BodySize: 100,
				},
				{
					BodySize: 200,
				},
			},
		},
	}

	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			b := beforeEach(t)
			var wg sync.WaitGroup
			for j := range testCases[i].requests {
				wg.Add(1)
				j := j
				go func() {
					defer wg.Done()
					_, err := b.svc.Route(testCases[i].requests[j])
					require.NoError(t, err)
				}()
			}
			wg.Wait()

			for _, n := range b.nodes {
				t.Log(n)
			}
		})
	}
}