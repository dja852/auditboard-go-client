/*
AuditBoard Developer Portal API Documentation

Testing EsgMetricAuditableEntitiesAPIService

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

func Test_auditboard_EsgMetricAuditableEntitiesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EsgMetricAuditableEntitiesAPIService EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgMetricAuditableEntityId int64

		resp, httpRes, err := apiClient.EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdDelete(context.Background(), esgMetricAuditableEntityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricAuditableEntitiesAPIService EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgMetricAuditableEntityId int64

		resp, httpRes, err := apiClient.EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdGet(context.Background(), esgMetricAuditableEntityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricAuditableEntitiesAPIService EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgMetricAuditableEntityId int64

		resp, httpRes, err := apiClient.EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesEsgMetricAuditableEntityIdPut(context.Background(), esgMetricAuditableEntityId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricAuditableEntitiesAPIService EsgMetricAuditableEntitiesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricAuditableEntitiesAPIService EsgMetricAuditableEntitiesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EsgMetricAuditableEntitiesAPI.EsgMetricAuditableEntitiesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
