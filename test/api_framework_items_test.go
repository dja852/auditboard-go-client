/*
AuditBoard Developer Portal API Documentation

Testing FrameworkItemsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package auditboard

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/dja852/auditboard-go-client"
)

func Test_auditboard_FrameworkItemsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FrameworkItemsAPIService FrameworkItemsFrameworkItemIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var frameworkItemId int64

		resp, httpRes, err := apiClient.FrameworkItemsAPI.FrameworkItemsFrameworkItemIdDelete(context.Background(), frameworkItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FrameworkItemsAPIService FrameworkItemsFrameworkItemIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var frameworkItemId int64

		resp, httpRes, err := apiClient.FrameworkItemsAPI.FrameworkItemsFrameworkItemIdGet(context.Background(), frameworkItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FrameworkItemsAPIService FrameworkItemsFrameworkItemIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var frameworkItemId int64

		resp, httpRes, err := apiClient.FrameworkItemsAPI.FrameworkItemsFrameworkItemIdPut(context.Background(), frameworkItemId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FrameworkItemsAPIService FrameworkItemsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FrameworkItemsAPI.FrameworkItemsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FrameworkItemsAPIService FrameworkItemsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FrameworkItemsAPI.FrameworkItemsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
