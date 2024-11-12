/*
AuditBoard Developer Portal API Documentation

Testing LocationsAPIService

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

func Test_auditboard_LocationsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LocationsAPIService LocationsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocationsAPI.LocationsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationsAPIService LocationsLocationIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var locationId int64

		resp, httpRes, err := apiClient.LocationsAPI.LocationsLocationIdDelete(context.Background(), locationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationsAPIService LocationsLocationIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var locationId int64

		resp, httpRes, err := apiClient.LocationsAPI.LocationsLocationIdGet(context.Background(), locationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationsAPIService LocationsLocationIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var locationId int64

		resp, httpRes, err := apiClient.LocationsAPI.LocationsLocationIdPut(context.Background(), locationId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LocationsAPIService LocationsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.LocationsAPI.LocationsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
