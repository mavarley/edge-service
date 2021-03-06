// Copyright SecureKey Technologies Inc. All Rights Reserved.
//
// SPDX-License-Identifier: Apache-2.0

module github.com/trustbloc/edge-service/cmd/vault-server

go 1.15

require (
	github.com/gorilla/mux v1.8.0
	github.com/hyperledger/aries-framework-go v0.1.6-0.20210203215323-783f4d353384
	github.com/rs/cors v1.7.0
	github.com/spf13/cobra v1.1.1
	github.com/stretchr/testify v1.6.1
	github.com/trustbloc/edge-core v0.1.6-0.20210127161542-9e174750f523
	github.com/trustbloc/edge-service v0.0.0-00010101000000-000000000000
)

replace github.com/trustbloc/edge-service => ../..
