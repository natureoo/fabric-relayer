/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package balanced

import (
	"github.com/polynetwork/fabric-relayer/pkg/common/logging"
	"github.com/polynetwork/fabric-relayer/pkg/common/options"
	"github.com/polynetwork/fabric-relayer/pkg/common/providers/context"
	"github.com/polynetwork/fabric-relayer/pkg/common/providers/fab"
	"github.com/polynetwork/fabric-relayer/pkg/fab/events/client/peerresolver"
	"github.com/polynetwork/fabric-relayer/pkg/fab/events/service"
)

var logger = logging.NewLogger("fabsdk/fab")

// PeerResolver is a peer resolver that chooses peers using the provided load balancer.
type PeerResolver struct {
	*params
}

// NewResolver returns a new "balanced" peer resolver provider.
func NewResolver() peerresolver.Provider {
	return func(ed service.Dispatcher, context context.Client, channelID string, opts ...options.Opt) peerresolver.Resolver {
		return New(ed, context, channelID, opts...)
	}
}

// New returns a new "balanced" peer resolver.
func New(dispatcher service.Dispatcher, context context.Client, channelID string, opts ...options.Opt) *PeerResolver {
	params := defaultParams(context, channelID)
	options.Apply(params, opts)

	logger.Debugf("Creating new balanced peer resolver")

	return &PeerResolver{
		params: params,
	}
}

// Resolve returns a peer usig the configured load balancer.
func (r *PeerResolver) Resolve(peers []fab.Peer) (fab.Peer, error) {
	return r.loadBalancePolicy.Choose(peers)
}

// ShouldDisconnect always returns false (will not disconnect a connected peer)
func (r *PeerResolver) ShouldDisconnect(peers []fab.Peer, connectedPeer fab.Peer) bool {
	return false
}
