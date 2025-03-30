//go:build with_dhcp

package include

import (
	"github.com/yafromil88/sing-box/dns"
	"github.com/yafromil88/sing-box/dns/transport/dhcp"
)

func registerDHCPTransport(registry *dns.TransportRegistry) {
	dhcp.RegisterTransport(registry)
}
