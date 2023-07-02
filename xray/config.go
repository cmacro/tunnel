package xray

import (
	"bytes"
	"encoding/json"
)

type Config struct {
	LogConfig       json.RawMessage `json:"log"`
	RouterConfig    json.RawMessage `json:"routing"`
	DNSConfig       json.RawMessage `json:"dns"`
	InboundConfigs  []InboundConfig `json:"inbounds"`
	OutboundConfigs json.RawMessage `json:"outbounds"`
	Transport       json.RawMessage `json:"transport"`
	Policy          json.RawMessage `json:"policy"`
	API             json.RawMessage `json:"api"`
	Stats           json.RawMessage `json:"stats"`
	Reverse         json.RawMessage `json:"reverse"`
	FakeDNS         json.RawMessage `json:"fakeDns"`
}

func (c *Config) Equals(other *Config) bool {
	if len(c.InboundConfigs) != len(other.InboundConfigs) {
		return false
	}
	for i, inbound := range c.InboundConfigs {
		if !inbound.Equals(&other.InboundConfigs[i]) {
			return false
		}
	}
	if !bytes.Equal(c.LogConfig, other.LogConfig) {
		return false
	}
	if !bytes.Equal(c.RouterConfig, other.RouterConfig) {
		return false
	}
	if !bytes.Equal(c.DNSConfig, other.DNSConfig) {
		return false
	}
	if !bytes.Equal(c.OutboundConfigs, other.OutboundConfigs) {
		return false
	}
	if !bytes.Equal(c.Transport, other.Transport) {
		return false
	}
	if !bytes.Equal(c.Policy, other.Policy) {
		return false
	}
	if !bytes.Equal(c.API, other.API) {
		return false
	}
	if !bytes.Equal(c.Stats, other.Stats) {
		return false
	}
	if !bytes.Equal(c.Reverse, other.Reverse) {
		return false
	}
	if !bytes.Equal(c.FakeDNS, other.FakeDNS) {
		return false
	}
	return true
}
