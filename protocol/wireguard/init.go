package wireguard

import (
	"github.com/yafromil88/sing-box/common/dialer"
	"github.com/sagernet/wireguard-go/conn"
)

func init() {
	dialer.WgControlFns = conn.ControlFns
}
