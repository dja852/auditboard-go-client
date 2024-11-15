/*
AuditBoard Developer Portal API Documentation

Testing RcwRequestsAPIService

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

func Test_auditboard_RcwRequestsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RcwRequestsAPIService RcwRequestsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RcwRequestsAPI.RcwRequestsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RcwRequestsAPIService RcwRequestsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RcwRequestsAPI.RcwRequestsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RcwRequestsAPIService RcwRequestsRcwRequestIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rcwRequestId int64

		resp, httpRes, err := apiClient.RcwRequestsAPI.RcwRequestsRcwRequestIdDelete(context.Background(), rcwRequestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RcwRequestsAPIService RcwRequestsRcwRequestIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rcwRequestId int64

		resp, httpRes, err := apiClient.RcwRequestsAPI.RcwRequestsRcwRequestIdGet(context.Background(), rcwRequestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RcwRequestsAPIService RcwRequestsRcwRequestIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rcwRequestId int64

		resp, httpRes, err := apiClient.RcwRequestsAPI.RcwRequestsRcwRequestIdPut(context.Background(), rcwRequestId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
