/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fabsdk

import (
	"github.com/polynetwork/fabric-relayer/pkg/core/logging/api"
	"github.com/polynetwork/fabric-relayer/pkg/core/logging/modlog"
	sdkApi "github.com/polynetwork/fabric-relayer/pkg/fabsdk/api"
	"github.com/polynetwork/fabric-relayer/pkg/fabsdk/factory/defcore"
	"github.com/polynetwork/fabric-relayer/pkg/fabsdk/factory/defmsp"
	"github.com/polynetwork/fabric-relayer/pkg/fabsdk/factory/defsvc"
)

type defPkgSuite struct{}

func (ps *defPkgSuite) Core() (sdkApi.CoreProviderFactory, error) {
	return defcore.NewProviderFactory(), nil
}

func (ps *defPkgSuite) MSP() (sdkApi.MSPProviderFactory, error) {
	return defmsp.NewProviderFactory(), nil
}

func (ps *defPkgSuite) Service() (sdkApi.ServiceProviderFactory, error) {
	return defsvc.NewProviderFactory(), nil
}

func (ps *defPkgSuite) Logger() (api.LoggerProvider, error) {
	return modlog.LoggerProvider(), nil
}
