//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by iocli

package protocol_impl

import (
	autowire "github.com/alibaba/ioc-golang/autowire"
	"github.com/alibaba/ioc-golang/autowire/normal"
)

func init() {
	normal.RegisterStructDescriptor(&autowire.StructDescriptor{
		Factory: func() interface{} {
			return &IOCProtocol{}
		},
		ParamFactory: func() interface{} {
			var _ paramInterface = &Param{}
			return &Param{}
		},
		ConstructFunc: func(i interface{}, p interface{}) (interface{}, error) {
			param := p.(paramInterface)
			impl := i.(*IOCProtocol)
			return param.Init(impl)
		},
	})
}

type paramInterface interface {
	Init(impl *IOCProtocol) (*IOCProtocol, error)
}
