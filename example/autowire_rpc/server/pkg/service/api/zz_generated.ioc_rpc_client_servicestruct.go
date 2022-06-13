// Code generated by iocli

package api

import (
	"github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/normal"
	"github.com/alibaba/ioc-golang/example/autowire_rpc/server/pkg/dto"
	"github.com/alibaba/ioc-golang/extension/autowire/rpc/rpc_client"
)

func init() {
	rpc_client.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceStructIOCRPCClient{}
		},
	})
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &serviceStructIOCRPCClient_{}
		},
	})
}

type serviceStructIOCRPCClient_ struct {
	GetUser_ func(name string, age int) (*dto.User, error)
}

func (s *serviceStructIOCRPCClient_) GetUser(name string, age int) (*dto.User, error) {
	return s.GetUser_(name, age)
}

type ServiceStructIOCRPCClient interface {
	GetUser(name string, age int) (*dto.User, error)
}

type serviceStructIOCRPCClient struct {
	GetUser func(name string, age int) (*dto.User, error)
}