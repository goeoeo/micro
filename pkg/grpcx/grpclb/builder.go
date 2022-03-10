package grpclb

import (
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/balancer/roundrobin"
	"google.golang.org/grpc/resolver"

	"git.internal.yunify.com/benchmark/benchmark/pkg/infra/registry"
)

type Resolver struct {
	exit chan bool

	serviceName string            //服务名称
	scheme      string            //自定义协议类型
	registry    registry.Registry //注册中心
	cc          resolver.ClientConn
	lb          string //负载均衡策略
}

func NewResolver(registry registry.Registry, serviceName string) *Resolver {
	r := &Resolver{
		serviceName: serviceName,
		scheme:      "customerDsn",
		lb:          fmt.Sprintf(`{"loadBalancingPolicy": "%s"}`, roundrobin.Name),
		registry:    registry,
		exit:        make(chan bool),
	}

	return r
}

func (r *Resolver) Endpoint() string {
	return fmt.Sprintf("%s:///%s", r.scheme, r.serviceName)
}

//为给定目标创建一个新的resolver grpc.Dail的时候执行
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (rl resolver.Resolver, err error) {

	r.cc = cc

	if r.serviceName != target.Endpoint {
		return nil, fmt.Errorf("serviceName err[%s,%s]", r.serviceName, target.Endpoint)
	}

	go r.watch()

	if err = r.updateState(); err != nil {
		return
	}

	return r, nil
}

//监听注册中心服务列表变更
func (r *Resolver) watch() {
	var attempts int

	for {

		// watch for changes
		w, err := r.registry.Watch()
		if err != nil {
			attempts++
			logrus.Errorf("error watching endpoints: %v", err)
			time.Sleep(time.Duration(attempts) * time.Second)
			continue
		}

		ch := make(chan bool)

		go func() {
			select {
			case <-ch:
				w.Stop()
			case <-r.exit:
				w.Stop()
			}
		}()

		// reset if we get here
		attempts = 0

		for {
			// process next event
			_, err := w.Next()
			if err != nil {
				logrus.Errorf("error getting next endoint: %v", err)
				close(ch)
				break
			}

			if err = r.updateState(); err != nil {
				logrus.Warningf("Resolver updateState err:%v", err)
			}
		}
	}
}

func (r *Resolver) updateState() (err error) {
	var (
		services []*registry.Service
		state    resolver.State
		address  []string
	)
	if services, err = r.registry.GetService(r.serviceName); err != nil {
		return
	}

	for _, v := range services {
		for _, v1 := range v.Nodes {
			address = append(address, v1.Address)
			state.Addresses = append(state.Addresses, resolver.Address{
				Addr:       v1.Address,
				ServerName: r.serviceName,
			})
		}
	}

	if len(state.Addresses) == 0 {
		r.cc.ReportError(fmt.Errorf("service address is empty"))
		return
	}

	if err = r.cc.UpdateState(state); err != nil {
		return
	}

	logrus.Infof("update service address success [%s,%s]", r.serviceName, address)

	return
}

func (r *Resolver) ResolveNow(o resolver.ResolveNowOptions) {
}

func (r *Resolver) Close() {
}

func (r *Resolver) Scheme() string {
	return r.scheme
}

func (r *Resolver) DailOptions() []grpc.DialOption {
	return []grpc.DialOption{
		grpc.WithResolvers(r),
		grpc.WithDefaultServiceConfig(r.lb), //加入负载均衡策略
	}
}
