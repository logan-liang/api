package svc

import (
	"github.com/logan-liang/api/internal/config"

	"github.com/logan-liang/rpc/demo"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config  config.Config
	TestRpc demo.Demo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		TestRpc: demo.NewDemo(zrpc.MustNewClient(c.TestRpc)),
	}
}
