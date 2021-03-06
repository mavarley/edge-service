/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hyperledger/aries-framework-go/pkg/doc/did"
	vdrapi "github.com/hyperledger/aries-framework-go/pkg/framework/aries/api/vdr"
	mockkms "github.com/hyperledger/aries-framework-go/pkg/mock/kms"
	mockstorage "github.com/hyperledger/aries-framework-go/pkg/mock/storage"
	"github.com/hyperledger/aries-framework-go/pkg/mock/vdr"
	"github.com/stretchr/testify/require"

	"github.com/trustbloc/edge-service/pkg/restapi/comparator/operation"
)

func Test_New(t *testing.T) {
	t.Run("test success", func(t *testing.T) {
		s := &mockstorage.MockStore{Store: make(map[string][]byte)}
		op, err := operation.New(&operation.Config{StoreProvider: &mockstorage.MockStoreProvider{
			Store: s}, KeyManager: &mockkms.KeyManager{}, VDR: &vdr.MockVDRegistry{
			CreateFunc: func(s string, doc *did.Doc, option ...vdrapi.DIDMethodOption) (*did.DocResolution, error) {
				return &did.DocResolution{DIDDocument: &did.Doc{ID: "did:ex:123"}}, nil
			}}})
		require.NoError(t, err)
		require.NotNil(t, op)

		require.Equal(t, 4, len(op.GetRESTHandlers()))
	})

	t.Run("test failed to create store", func(t *testing.T) {
		_, err := operation.New(&operation.Config{StoreProvider: &mockstorage.MockStoreProvider{
			ErrOpenStoreHandle: fmt.Errorf("failed to open store")}})
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to open store")
	})

	t.Run("test failed to export public key", func(t *testing.T) {
		s := &mockstorage.MockStore{Store: make(map[string][]byte)}
		_, err := operation.New(&operation.Config{StoreProvider: &mockstorage.MockStoreProvider{
			Store: s}, KeyManager: &mockkms.KeyManager{CrAndExportPubKeyErr: fmt.Errorf("failed to export")}})
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to export")
	})

	t.Run("test failed to get config", func(t *testing.T) {
		s := &mockstorage.MockStore{Store: make(map[string][]byte)}
		s.ErrGet = fmt.Errorf("failed to get config")
		_, err := operation.New(&operation.Config{StoreProvider: &mockstorage.MockStoreProvider{
			Store: s}})
		require.Error(t, err)
		require.Contains(t, err.Error(), "failed to get config")
	})
}

func TestOperation_CreateAuthorization(t *testing.T) {
	t.Run("TODO - creates an authorization", func(t *testing.T) {
		s := &mockstorage.MockStore{Store: make(map[string][]byte)}
		s.Store["config"] = []byte(`{}`)
		op, err := operation.New(&operation.Config{StoreProvider: &mockstorage.MockStoreProvider{
			Store: s}})
		require.NoError(t, err)
		require.NotNil(t, op)
		result := httptest.NewRecorder()
		op.CreateAuthorization(result, nil)
		require.Equal(t, http.StatusCreated, result.Code)
	})
}

func TestOperation_Compare(t *testing.T) {
	t.Run("TODO - runs a comparison", func(t *testing.T) {
		s := &mockstorage.MockStore{Store: make(map[string][]byte)}
		s.Store["config"] = []byte(`{}`)
		op, err := operation.New(&operation.Config{StoreProvider: &mockstorage.MockStoreProvider{
			Store: s}})
		require.NoError(t, err)
		require.NotNil(t, op)
		result := httptest.NewRecorder()
		op.Compare(result, nil)
		require.Equal(t, http.StatusOK, result.Code)
	})
}

func TestOperation_Extract(t *testing.T) {
	t.Run("TODO - performs an extraction", func(t *testing.T) {
		s := &mockstorage.MockStore{Store: make(map[string][]byte)}
		s.Store["config"] = []byte(`{}`)
		op, err := operation.New(&operation.Config{StoreProvider: &mockstorage.MockStoreProvider{
			Store: s}})
		require.NoError(t, err)
		require.NotNil(t, op)
		result := httptest.NewRecorder()
		op.Extract(result, nil)
		require.Equal(t, http.StatusOK, result.Code)
	})
}

func TestOperation_GetConfig(t *testing.T) {
	t.Run("get config success", func(t *testing.T) {
		s := make(map[string][]byte)
		s["config"] = []byte(`{}`)
		op, err := operation.New(&operation.Config{StoreProvider: &mockstorage.MockStoreProvider{
			Store: &mockstorage.MockStore{Store: s}}})
		require.NoError(t, err)
		require.NotNil(t, op)
		result := httptest.NewRecorder()
		op.GetConfig(result, nil)
		require.Equal(t, http.StatusOK, result.Code)
		require.Contains(t, result.Body.String(), "did")
	})

	t.Run("get config not found", func(t *testing.T) {
		s := &mockstorage.MockStore{Store: make(map[string][]byte)}
		s.Store["config"] = []byte(`{}`)
		op, err := operation.New(&operation.Config{StoreProvider: &mockstorage.MockStoreProvider{
			Store: s}})
		delete(s.Store, "config")
		require.NoError(t, err)
		require.NotNil(t, op)
		result := httptest.NewRecorder()
		op.GetConfig(result, nil)
		require.Equal(t, http.StatusNotFound, result.Code)
	})

	t.Run("get config error", func(t *testing.T) {
		s := &mockstorage.MockStore{Store: make(map[string][]byte)}
		s.Store["config"] = []byte(`{}`)
		op, err := operation.New(&operation.Config{StoreProvider: &mockstorage.MockStoreProvider{
			Store: s}})
		s.ErrGet = fmt.Errorf("failed to get config")
		require.NoError(t, err)
		require.NotNil(t, op)
		result := httptest.NewRecorder()
		op.GetConfig(result, nil)
		require.Equal(t, http.StatusInternalServerError, result.Code)
		require.Contains(t, result.Body.String(), "failed to get config")
	})
}
