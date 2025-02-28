/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package dynamicdiscovery

import (
	"github.com/pkg/errors"
	"github.com/polynetwork/fabric-relayer/pkg/common/logging"
	coptions "github.com/polynetwork/fabric-relayer/pkg/common/options"
	"github.com/polynetwork/fabric-relayer/pkg/common/providers/fab"
	"github.com/polynetwork/fabric-relayer/pkg/util/concurrent/lazycache"
)

var logger = logging.NewLogger("fabsdk/client")

// LocalProvider implements a local Dynamic Discovery LocalProvider that queries
// Fabric's Discovery service for information about the peers that
// are in the local MSP.
type LocalProvider struct {
	cache *lazycache.Cache
}

// NewLocalProvider creates a new local dynamic discovery provider
func NewLocalProvider(config fab.EndpointConfig, opts ...coptions.Opt) *LocalProvider {
	return &LocalProvider{
		cache: lazycache.New("Local_Discovery_Service_Cache", func(key lazycache.Key) (interface{}, error) {
			return newLocalService(config, key.String(), opts...), nil
		}),
	}
}

// CreateLocalDiscoveryService returns a local discovery service
func (p *LocalProvider) CreateLocalDiscoveryService(mspID string) (fab.DiscoveryService, error) {
	ref, err := p.cache.Get(lazycache.NewStringKey(mspID))
	if err != nil {
		return nil, errors.WithMessage(err, "failed to get local discovery service from cache")
	}
	return ref.(fab.DiscoveryService), nil
}

// Close will close the cache and all services contained by the cache.
func (p *LocalProvider) Close() {
	logger.Debug("Closing local provider cache")
	p.cache.Close()
}

// CloseContext frees resources and caches for the given context.
func (p *LocalProvider) CloseContext(ctx fab.ClientContext) {
	mspID := ctx.Identifier().MSPID
	logger.Debugf("Closing local discovery service for MSP [%s]", mspID)
	p.cache.Delete(lazycache.NewStringKey(mspID))
}
