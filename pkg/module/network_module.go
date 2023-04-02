package module

import "github.com/DevellSoftware/metis-engine/pkg/network"

type NetworkModule struct {
	network *network.Network
}

func NewNetworkModule(network *network.Network) *NetworkModule {
	return &NetworkModule{
		network: network,
	}
}

func (m *NetworkModule) Forward() {
	m.network.Predict(nil)
}

func (m *NetworkModule) Backward() {
	m.network.Train(nil, nil, 0)
}
