/*
AuditBoard Developer Portal API Documentation

Testing AuditableEntitiesAPIService

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

func Test_auditboard_AuditableEntitiesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuditableEntitiesAPIService AuditableEntitiesAuditableEntityIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditableEntityId int64

		resp, httpRes, err := apiClient.AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdDelete(context.Background(), auditableEntityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditableEntitiesAPIService AuditableEntitiesAuditableEntityIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditableEntityId int64

		resp, httpRes, err := apiClient.AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdGet(context.Background(), auditableEntityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditableEntitiesAPIService AuditableEntitiesAuditableEntityIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditableEntityId int64

		resp, httpRes, err := apiClient.AuditableEntitiesAPI.AuditableEntitiesAuditableEntityIdPut(context.Background(), auditableEntityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditableEntitiesAPIService AuditableEntitiesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuditableEntitiesAPI.AuditableEntitiesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditableEntitiesAPIService AuditableEntitiesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuditableEntitiesAPI.AuditableEntitiesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
