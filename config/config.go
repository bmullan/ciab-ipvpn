package config

import (
	"encoding/json"
	"os"
	"strings"
	"sync/atomic"

	"github.com/caarlos0/env"
)

var configInstance atomic.Value

type config struct {
	DataDirectory             string `env:"IPVPN_DATADIR" envDefault:"${HOME}/.ipvpn"`
	NetworkSubnet             string `env:"IPVPN_NETWORK_SUBNET" envDefault:"172.16.0.0/12"`
	DumpVPNCommunications     bool   `env:"IPVPN_NETWORK_DUMP_VPN"`
	DumpNetworkCommunications bool   `env:"IPVPN_NETWORK_DUMP_MESH"`
	DumpConfiguration         bool   `env:"IPVPN_DUMP_CONFIG"`
	DisableWGThroughIPFS      bool   `env:"IPVPN_DISABLE_WG_THROUGH_IPFS"`
	DisableWGThroughTunnel    bool   `env:"IPVPN_DISABLE_WG_THROUGH_TUNNEL"`
	DisableWGThroughDirect    bool   `env:"IPVPN_DISABLE_WG_THROUGH_DIRECT"`
	PprofNetAddress           string `env:"IPVPN_PPROF_NET_ADDRESS"`
}

func (cfg config) String() string {
	jsonBytes, err := json.MarshalIndent(cfg, "", "\t")
	panicIf(err)
	return string(jsonBytes)
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func init() {
	cfg := &config{}
	panicIf(env.Parse(cfg))
	homedir, _ := os.UserHomeDir()
	cfg.DataDirectory = strings.Replace(cfg.DataDirectory, `${HOME}`, homedir, -1)
	configInstance.Store(cfg)
}

func Get() config {
	return *configInstance.Load().(*config)
}
