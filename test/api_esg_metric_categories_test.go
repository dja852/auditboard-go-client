/*
AuditBoard Developer Portal API Documentation

Testing EsgMetricCategoriesAPIService

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

func Test_auditboard_EsgMetricCategoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EsgMetricCategoriesAPIService EsgMetricCategoriesEsgMetricCategoryIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgMetricCategoryId int64

		resp, httpRes, err := apiClient.EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdDelete(context.Background(), esgMetricCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricCategoriesAPIService EsgMetricCategoriesEsgMetricCategoryIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgMetricCategoryId int64

		resp, httpRes, err := apiClient.EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdGet(context.Background(), esgMetricCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricCategoriesAPIService EsgMetricCategoriesEsgMetricCategoryIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgMetricCategoryId int64

		resp, httpRes, err := apiClient.EsgMetricCategoriesAPI.EsgMetricCategoriesEsgMetricCategoryIdPut(context.Background(), esgMetricCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricCategoriesAPIService EsgMetricCategoriesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EsgMetricCategoriesAPI.EsgMetricCategoriesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricCategoriesAPIService EsgMetricCategoriesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EsgMetricCategoriesAPI.EsgMetricCategoriesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
