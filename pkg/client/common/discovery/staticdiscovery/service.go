/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package staticdiscovery

import (
	"github.com/pkg/errors"
	"github.com/polynetwork/fabric-relayer/pkg/common/providers/fab"
)

// DiscoveryService implements a static discovery service
type DiscoveryService struct {
	peers []fab.Peer
}

// NewService creates a static discovery service
func NewService(config fab.EndpointConfig, peerCreator peerCreator, channelID string) (*DiscoveryService, error) {
	if channelID == "" {
		return nil, errors.New("channel ID must be provided")
	}

	// Use configured channel peers
	chPeers := config.ChannelPeers(channelID)
	if len(chPeers) == 0 {
		return nil, errors.Errorf("no channel peers configured for channel [%s]", channelID)
	}

	peers := []fab.Peer{}
	for _, p := range chPeers {
		newPeer, err := peerCreator.CreatePeerFromConfig(&p.NetworkPeer)
		if err != nil || newPeer == nil {
			return nil, errors.WithMessage(err, "NewPeer failed")
		}

		peers = append(peers, newPeer)
	}

	return &DiscoveryService{
		peers: peers,
	}, nil
}

// GetPeers is used to get peers
func (ds *DiscoveryService) GetPeers() ([]fab.Peer, error) {
	return ds.peers, nil
}
