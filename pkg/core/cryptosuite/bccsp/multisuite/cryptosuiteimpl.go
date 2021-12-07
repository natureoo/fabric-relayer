/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package multisuite

import (
	"github.com/pkg/errors"
	"github.com/polynetwork/fabric-relayer/pkg/common/providers/core"
	"github.com/polynetwork/fabric-relayer/pkg/core/cryptosuite/bccsp/pkcs11"
	"github.com/polynetwork/fabric-relayer/pkg/core/cryptosuite/bccsp/sw"
)

//GetSuiteByConfig returns cryptosuite adaptor for bccsp loaded according to given config
func GetSuiteByConfig(config core.CryptoSuiteConfig) (core.CryptoSuite, error) {
	switch config.SecurityProvider() {
	case "sw":
		return sw.GetSuiteByConfig(config)
	case "pkcs11":
		return pkcs11.GetSuiteByConfig(config)
	}

	return nil, errors.Errorf("Unsupported security provider requested: %s", config.SecurityProvider())
}
