package v2raygrpclite

import (
	"context"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/yafromil88/sing-box/adapter"
	"github.com/yafromil88/sing-box/common/tls"
	"github.com/yafromil88/sing-box/option"
	"github.com/yafromil88/sing-box/transport/v2rayhttp"
	E "github.com/sagernet/sing/common/exceptions"
	M "github.com/sagernet/sing/common/metadata"
	N "github.com/sagernet/sing/common/network"

	"golang.org/x/net/http2"
)

var _ adapter.V2RayClientTransport = (*Client)(nil)

var defaultClientHeader = http.Header{
	"Content-Type": []string{"application/grpc"},
	"User-Agent":   []string{"grpc-go/1.48.0"},
	"TE":           []string{"trailers"},
}

type Client struct {
	ctx        context.Context
	dialer     N.Dialer
	serverAddr M.Socksaddr
	transport  *http2.Transport
	options    option.V2RayGRPCOptions
	url        *url.URL
	host       string
}

func NewClient(ctx context.Context, dialer N.Dialer, serverAddr M.Socksaddr, options option.V2RayGRPCOptions, tlsConfig tls.Config) adapter.V2RayClientTransport {
	var host string
	if tlsConfig != nil && tlsConfig.ServerName() != "" {
		host = M.ParseSocksaddrHostPort(tlsConfig.ServerName(), serverAddr.Port).String()
	} else {
		host = serverAddr.String()
	}
	client := &Client{
		ctx:        ctx,
		dialer:     dialer,
		serverAddr: serverAddr,
		options:    options,
		transport: &http2.Transport{
			ReadIdleTimeout:    time.Duration(options.IdleTimeout),
			PingTimeout:        time.Duration(options.PingTimeout),
			DisableCompression: true,
		},
		url: &url.URL{
			Scheme:  "https",
			Host:    serverAddr.String(),
			Path:    "/" + options.ServiceName + "/Tun",
			RawPath: "/" + url.PathEscape(options.ServiceName) + "/Tun",
		},
		host: host,
	}

	if tlsConfig == nil {
		client.transport.DialTLSContext = func(ctx context.Context, network, addr string, cfg *tls.STDConfig) (net.Conn, error) {
			return dialer.DialContext(ctx, network, M.ParseSocksaddr(addr))
		}
	} else {
		if len(tlsConfig.NextProtos()) == 0 {
			tlsConfig.SetNextProtos([]string{http2.NextProtoTLS})
		}
		client.transport.DialTLSContext = func(ctx context.Context, network, addr string, cfg *tls.STDConfig) (net.Conn, error) {
			conn, err := dialer.DialContext(ctx, network, M.ParseSocksaddr(addr))
			if err != nil {
				return nil, err
			}
			return tls.ClientHandshake(ctx, conn, tlsConfig)
		}
	}

	return client
}

func (c *Client) DialContext(ctx context.Context) (net.Conn, error) {
	pipeInReader, pipeInWriter := io.Pipe()
	request := &http.Request{
		Method: http.MethodPost,
		Body:   pipeInReader,
		URL:    c.url,
		Header: defaultClientHeader,
		Host:   c.host,
	}
	request = request.WithContext(ctx)
	conn := newLateGunConn(pipeInWriter)
	go func() {
		response, err := c.transport.RoundTrip(request)
		if err != nil {
			conn.setup(nil, err)
		} else if response.StatusCode != 200 {
			response.Body.Close()
			conn.setup(nil, E.New("v2ray-grpc: unexpected status: ", response.Status))
		} else {
			conn.setup(response.Body, nil)
		}
	}()
	return conn, nil
}

func (c *Client) Close() error {
	v2rayhttp.ResetTransport(c.transport)
	return nil
}
