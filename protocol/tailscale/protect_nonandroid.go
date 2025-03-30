//go:build !android

package tailscale

import "github.com/yafromil88/sing-box/experimental/libbox/platform"

func setAndroidProtectFunc(platformInterface platform.Interface) {
}
