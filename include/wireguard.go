//go:build with_wireguard

package include

import (
	"github.com/yafromil88/sing-box/adapter/endpoint"
	"github.com/yafromil88/sing-box/adapter/outbound"
	"github.com/yafromil88/sing-box/protocol/wireguard"
)

func registerWireGuardOutbound(registry *outbound.Registry) {
	wireguard.RegisterOutbound(registry)
}

func registerWireGuardEndpoint(registry *endpoint.Registry) {
	wireguard.RegisterEndpoint(registry)
}
