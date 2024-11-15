/*
AuditBoard Developer Portal API Documentation

Testing GlobalAttributesAPIService

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

func Test_auditboard_GlobalAttributesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test GlobalAttributesAPIService GlobalAttributesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GlobalAttributesAPI.GlobalAttributesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAttributesAPIService GlobalAttributesGlobalAttributeIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalAttributeId int64

		resp, httpRes, err := apiClient.GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdDelete(context.Background(), globalAttributeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAttributesAPIService GlobalAttributesGlobalAttributeIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalAttributeId int64

		resp, httpRes, err := apiClient.GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdGet(context.Background(), globalAttributeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAttributesAPIService GlobalAttributesGlobalAttributeIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var globalAttributeId int64

		resp, httpRes, err := apiClient.GlobalAttributesAPI.GlobalAttributesGlobalAttributeIdPut(context.Background(), globalAttributeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test GlobalAttributesAPIService GlobalAttributesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.GlobalAttributesAPI.GlobalAttributesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
