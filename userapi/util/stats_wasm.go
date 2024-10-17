// Copyright 2024 New Vector Ltd.
// Copyright 2022 The Matrix.org Foundation C.I.C.
//
// SPDX-License-Identifier: AGPL-3.0-only
// Please see LICENSE in the repository root for full details.

package util

// stub, since WASM doesn't support syscall.Getrusage
func getMemoryStats(p *phoneHomeStats) error {
	return nil
}
