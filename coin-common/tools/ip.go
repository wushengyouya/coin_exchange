package tools

import (
	"net"
	"net/http"
)

// 获取远程请求的ip
func GetRemoteClientIp(r *http.Request) string {
	remoteIp := r.RemoteAddr
	if ip := r.Header.Get("X-Real-IP"); ip != "" {
		remoteIp = ip
	} else if ip = r.Header.Get("X-Forwarede-For"); ip != "" {
		remoteIp = ip
	} else {
		remoteIp, _, _ = net.SplitHostPort(remoteIp)
	}
	if remoteIp == "::1" {
		remoteIp = "127.0.0.1"
	}
	return remoteIp
}
