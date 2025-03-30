//go:build !with_dhcp

package include

import (
	"context"

	"github.com/yafromil88/sing-box/adapter"
	C "github.com/yafromil88/sing-box/constant"
	"github.com/yafromil88/sing-box/dns"
	"github.com/yafromil88/sing-box/log"
	"github.com/yafromil88/sing-box/option"
	E "github.com/sagernet/sing/common/exceptions"
)

func registerDHCPTransport(registry *dns.TransportRegistry) {
	dns.RegisterTransport[option.DHCPDNSServerOptions](registry, C.DNSTypeDHCP, func(ctx context.Context, logger log.ContextLogger, tag string, options option.DHCPDNSServerOptions) (adapter.DNSTransport, error) {
		return nil, E.New(`DHCP is not included in this build, rebuild with -tags with_dhcp`)
	})
}
