/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package msp

import (
	"github.com/pkg/errors"
	"github.com/polynetwork/fabric-relayer/pkg/common/providers/msp"
)

var (
	// ErrUserNotFound indicates the user was not found
	ErrUserNotFound = errors.New("user not found")
)

// IdentityManager provides management of identities in a Fabric network
type IdentityManager interface {
	GetSigningIdentity(name string) (msp.SigningIdentity, error)
	CreateSigningIdentity(ops ...msp.SigningIdentityOption) (msp.SigningIdentity, error)
}
