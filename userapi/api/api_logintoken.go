// Copyright 2024 New Vector Ltd.
// Copyright 2021 The Matrix.org Foundation C.I.C.
//
// SPDX-License-Identifier: AGPL-3.0-only
// Please see LICENSE in the repository root for full details.

package api

import (
	"context"
	"time"
)

// DefaultLoginTokenLifetime determines how old a valid token may be.
//
// NOTSPEC: The current spec says "SHOULD be limited to around five
// seconds". Since TCP retries are on the order of 3 s, 5 s sounds very low.
// Synapse uses 2 min (https://github.com/matrix-org/synapse/blob/78d5f91de1a9baf4dbb0a794cb49a799f29f7a38/synapse/handlers/auth.py#L1323-L1325).
const DefaultLoginTokenLifetime = 2 * time.Minute

type LoginTokenInternalAPI interface {
	// PerformLoginTokenCreation creates a new login token and associates it with the provided data.
	PerformLoginTokenCreation(ctx context.Context, req *PerformLoginTokenCreationRequest, res *PerformLoginTokenCreationResponse) error

	// PerformLoginTokenDeletion ensures the token doesn't exist. Success
	// is returned even if the token didn't exist, or was already deleted.
	PerformLoginTokenDeletion(ctx context.Context, req *PerformLoginTokenDeletionRequest, res *PerformLoginTokenDeletionResponse) error

	// QueryLoginToken returns the data associated with a login token. If
	// the token is not valid, success is returned, but res.Data == nil.
	QueryLoginToken(ctx context.Context, req *QueryLoginTokenRequest, res *QueryLoginTokenResponse) error
}

// LoginTokenData is the data that can be retrieved given a login token. This is
// provided by the calling code.
type LoginTokenData struct {
	// UserID is the full mxid of the user.
	UserID string
}

// LoginTokenMetadata contains metadata created and maintained by the User API.
type LoginTokenMetadata struct {
	Token      string
	Expiration time.Time
}

type PerformLoginTokenCreationRequest struct {
	Data LoginTokenData
}

type PerformLoginTokenCreationResponse struct {
	Metadata LoginTokenMetadata
}

type PerformLoginTokenDeletionRequest struct {
	Token string
}

type PerformLoginTokenDeletionResponse struct{}

type QueryLoginTokenRequest struct {
	Token string
}

type QueryLoginTokenResponse struct {
	// Data is nil if the token was invalid.
	Data *LoginTokenData
}
