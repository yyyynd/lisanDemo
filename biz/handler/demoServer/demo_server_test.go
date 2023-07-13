package demoServer

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/cloudwego/hertz/pkg/common/ut"
	"log"
	"testing"
)

func TestStuInformation(t *testing.T) {
	h := server.Default()
	h.GET("/treeStructure", TreeStructure)
	w := ut.PerformRequest(h.Engine, "GET", "/treeStructure",
		&ut.Body{}, ut.Header{"Connection", "close"})
	resp := w.Result()
	log.Println(resp.Body())
}
