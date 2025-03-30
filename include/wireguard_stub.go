//go:build !with_wireguard

package include

import (
	"context"

	"github.com/yafromil88/sing-box/adapter"
	"github.com/yafromil88/sing-box/adapter/endpoint"
	"github.com/yafromil88/sing-box/adapter/outbound"
	C "github.com/yafromil88/sing-box/constant"
	"github.com/yafromil88/sing-box/log"
	"github.com/yafromil88/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func registerWireGuardOutbound(registry *outbound.Registry) {
	outbound.Register[option.LegacyWireGuardOutboundOptions](registry, C.TypeWireGuard, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.LegacyWireGuardOutboundOptions) (adapter.Outbound, error) {
		return nil, E.New(`WireGuard is not included in this build, rebuild with -tags with_wireguard`)
	})
}

func registerWireGuardEndpoint(registry *endpoint.Registry) {
	endpoint.Register[option.WireGuardEndpointOptions](registry, C.TypeWireGuard, func(ctx context.Context, router adapter.Router, logger log.ContextLogger, tag string, options option.WireGuardEndpointOptions) (adapter.Endpoint, error) {
		return nil, E.New(`WireGuard is not included in this build, rebuild with -tags with_wireguard`)
	})
}
