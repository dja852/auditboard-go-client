/*
AuditBoard Developer Portal API Documentation

Testing EsgCategoriesAPIService

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

func Test_auditboard_EsgCategoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EsgCategoriesAPIService EsgCategoriesEsgCategoryIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgCategoryId int64

		resp, httpRes, err := apiClient.EsgCategoriesAPI.EsgCategoriesEsgCategoryIdDelete(context.Background(), esgCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgCategoriesAPIService EsgCategoriesEsgCategoryIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgCategoryId int64

		resp, httpRes, err := apiClient.EsgCategoriesAPI.EsgCategoriesEsgCategoryIdGet(context.Background(), esgCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgCategoriesAPIService EsgCategoriesEsgCategoryIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgCategoryId int64

		resp, httpRes, err := apiClient.EsgCategoriesAPI.EsgCategoriesEsgCategoryIdPut(context.Background(), esgCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgCategoriesAPIService EsgCategoriesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EsgCategoriesAPI.EsgCategoriesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgCategoriesAPIService EsgCategoriesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EsgCategoriesAPI.EsgCategoriesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
