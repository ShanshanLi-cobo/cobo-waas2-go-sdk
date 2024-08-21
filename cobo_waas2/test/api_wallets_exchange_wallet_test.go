/*
Cobo Wallet as a Service 2.0

Testing WalletsExchangeWalletAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package cobo_waas2

import (
	"context"
	"testing"

	coboWaas2 "github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2"
	"github.com/CoboGlobal/cobo-waas2-go-sdk/cobo_waas2/crypto"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/require"
)

func Test_cobo_waas2_WalletsExchangeWalletAPIService(t *testing.T) {

	configuration := coboWaas2.NewConfiguration()
	apiClient := coboWaas2.NewAPIClient(configuration)
	ctx := context.WithValue(context.Background(), coboWaas2.ContextServerHost, "https://api[.xxxx].cobo.com/v2")
	ctx = context.WithValue(ctx, coboWaas2.ContextPortalSigner, crypto.Ed25519Signer{
		Secret: "<YOUR_API_PRIV_KEY_IN_HEX>",
	})

	t.Run("Test WalletsExchangeWalletAPIService ListAssetBalancesForExchangeWallet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var walletId string

		resp, httpRes, err := apiClient.WalletsExchangeWalletAPI.ListAssetBalancesForExchangeWallet(ctx, walletId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsExchangeWalletAPIService ListExchanges", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WalletsExchangeWalletAPI.ListExchanges(ctx).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsExchangeWalletAPIService ListSupportedAssetsForExchange", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var exchangeId ExchangeId

		resp, httpRes, err := apiClient.WalletsExchangeWalletAPI.ListSupportedAssetsForExchange(ctx, exchangeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WalletsExchangeWalletAPIService ListSupportedChainsForExchange", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var exchangeId ExchangeId
		var assetId string

		resp, httpRes, err := apiClient.WalletsExchangeWalletAPI.ListSupportedChainsForExchange(ctx, exchangeId, assetId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}