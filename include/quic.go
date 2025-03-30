//go:build with_quic

package include

import (
	"github.com/yafromil88/sing-box/adapter/inbound"
	"github.com/yafromil88/sing-box/adapter/outbound"
	"github.com/yafromil88/sing-box/dns"
	"github.com/yafromil88/sing-box/dns/transport/quic"
	"github.com/yafromil88/sing-box/protocol/hysteria"
	"github.com/yafromil88/sing-box/protocol/hysteria2"
	_ "github.com/yafromil88/sing-box/protocol/naive/quic"
	"github.com/yafromil88/sing-box/protocol/tuic"
	_ "github.com/yafromil88/sing-box/transport/v2rayquic"
)

func registerQUICInbounds(registry *inbound.Registry) {
	hysteria.RegisterInbound(registry)
	tuic.RegisterInbound(registry)
	hysteria2.RegisterInbound(registry)
}

func registerQUICOutbounds(registry *outbound.Registry) {
	hysteria.RegisterOutbound(registry)
	tuic.RegisterOutbound(registry)
	hysteria2.RegisterOutbound(registry)
}

func registerQUICTransports(registry *dns.TransportRegistry) {
	quic.RegisterTransport(registry)
	quic.RegisterHTTP3Transport(registry)
}
