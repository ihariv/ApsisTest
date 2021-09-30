package main

import "github.com/gin-gonic/gin"

type (
	collectGin struct{
		*gin.Engine
		Apps []Node
	}
	Node interface{
		Route(router *gin.Engine)
	}
)

func initCollect(router *gin.Engine) (collect collectGin) {
	collect = collectGin{
		Engine: router,
		Apps: []Node{},
	}
	return
}

func (cg *collectGin) AddNodes(nodes ...Node){
	for _, node := range nodes{
		node.Route(cg.Engine)
		cg.Apps = append(cg.Apps, node)
	}
}
