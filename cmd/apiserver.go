package main

import (
	"coxwave-load-balancer-golang/app/service"
	"coxwave-load-balancer-golang/domain/load_balancer"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	heap := load_balancer.NodeHeap{
		{
			Address:       "node1",
			BPM:           20,
			RPM:           10,
		},
		{
			Address:       "node2",
			BPM:           200,
			RPM:           20,
		},
	}

	svc := service.NewLoadBalancer(heap)

	r.POST("/load_balance", func(c *gin.Context) {
		var r *load_balancer.Request

		if err := c.ShouldBindJSON(&r); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		node, err := svc.Route(r)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("Current Node: %s, UsedByteSize: %d, RequestCounts: %d", node.Address, node.UsedByteSize, node.RequestCounts),
		})
		return
	})

	r.Run()
}