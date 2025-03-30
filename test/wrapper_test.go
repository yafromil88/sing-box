package main

import (
	"testing"

	C "github.com/yafromil88/sing-box/constant"
	"github.com/yafromil88/sing-box/option"

	"github.com/stretchr/testify/require"
)

func TestOptionsWrapper(t *testing.T) {
	inbound := option.Inbound{
		Type: C.TypeHTTP,
		Options: &option.HTTPMixedInboundOptions{
			InboundTLSOptionsContainer: option.InboundTLSOptionsContainer{
				TLS: &option.InboundTLSOptions{
					Enabled: true,
				},
			},
		},
	}
	tlsOptionsWrapper, loaded := inbound.Options.(option.InboundTLSOptionsWrapper)
	require.True(t, loaded, "find inbound tls options")
	tlsOptions := tlsOptionsWrapper.TakeInboundTLSOptions()
	require.NotNil(t, tlsOptions, "find inbound tls options")
	tlsOptions.Enabled = false
	tlsOptionsWrapper.ReplaceInboundTLSOptions(tlsOptions)
	require.False(t, inbound.Options.(*option.HTTPMixedInboundOptions).TLS.Enabled, "replace tls enabled")
}
