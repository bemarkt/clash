package rules

import (
	"net"

	C "github.com/Dreamacro/clash/constant"
)

type HostIPCIDR struct {
	ipnet 	*net.IPNet
	adapter string
}

func (h *HostIPCIDR) RuleType() C.RuleType {
	return C.HostIPCIDR
}

func (h *HostIPCIDR) IsMatch(metadata *C.Metadata) bool {
	host := metadata.Host
	ip := net.ParseIP(host)
	return ip != nil && h.ipnet.Contains(ip)
}

func (h *HostIPCIDR) Adapter() string {
	return h.adapter
}

func (h *HostIPCIDR) Payload() string {
	return h.ipnet.String()
}

func NewHostIPCIDR(s string, adapter string) *HostIPCIDR {
	_, ipnet, err := net.ParseCIDR(s)
	if err != nil {
		return nil
	}
	return &HostIPCIDR{
		ipnet:      ipnet,
		adapter:    adapter,
	}
}