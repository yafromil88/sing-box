//go:build with_quic

package v2rayquic

import "github.com/yafromil88/sing-box/transport/v2ray"

func init() {
	v2ray.RegisterQUICConstructor(NewServer, NewClient)
}
