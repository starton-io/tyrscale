/*
Tyrscale Manager API

Testing PluginsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package tyrscalesdkgo

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/starton-io/tyrscale/sdk/tyrscale-sdk-go"
)

func Test_tyrscalesdkgo_PluginsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PluginsAPIService ListPlugins", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PluginsAPI.ListPlugins(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
