package proxy

import "testing"

func TestProxy(t *testing.T) {
	proxyIml := createProxy()
	ret := proxyIml.solveRequest()
	if ret != "[solve request complete]" {
		t.Error("proxy error")
	}
}