/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package endpoint

import (
	"crypto/x509"

	"github.com/polynetwork/fabric-relayer/pkg/common/options"
	"github.com/polynetwork/fabric-relayer/pkg/common/providers/fab"
	"github.com/polynetwork/fabric-relayer/pkg/fab/comm"
)

// EventEndpoint extends a Peer endpoint
type EventEndpoint struct {
	Certificate *x509.Certificate
	fab.Peer
	opts []options.Opt
}

// Opts returns additional options for the event connection
func (e *EventEndpoint) Opts() []options.Opt {
	return e.opts
}

// BlockHeight returns the block height of the peer. If the peer doesn't contain any state info then 0 is returned.
func (e *EventEndpoint) BlockHeight() uint64 {
	peerState, ok := e.Peer.(fab.PeerState)
	if !ok {
		return 0
	}
	return peerState.BlockHeight()
}

// FromPeerConfig creates a new EventEndpoint from the given config
func FromPeerConfig(config fab.EndpointConfig, peer fab.Peer, peerCfg *fab.PeerConfig) *EventEndpoint {
	opts := comm.OptsFromPeerConfig(peerCfg)
	opts = append(opts, comm.WithConnectTimeout(config.Timeout(fab.PeerConnection)))

	return &EventEndpoint{
		Peer: peer,
		opts: opts,
	}
}
