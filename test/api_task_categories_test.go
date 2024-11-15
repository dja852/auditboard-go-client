/*
AuditBoard Developer Portal API Documentation

Testing TaskCategoriesAPIService

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

func Test_auditboard_TaskCategoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TaskCategoriesAPIService TaskCategoriesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TaskCategoriesAPI.TaskCategoriesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskCategoriesAPIService TaskCategoriesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TaskCategoriesAPI.TaskCategoriesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskCategoriesAPIService TaskCategoriesTaskCategoryIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskCategoryId int64

		resp, httpRes, err := apiClient.TaskCategoriesAPI.TaskCategoriesTaskCategoryIdDelete(context.Background(), taskCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskCategoriesAPIService TaskCategoriesTaskCategoryIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskCategoryId int64

		resp, httpRes, err := apiClient.TaskCategoriesAPI.TaskCategoriesTaskCategoryIdGet(context.Background(), taskCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TaskCategoriesAPIService TaskCategoriesTaskCategoryIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var taskCategoryId int64

		resp, httpRes, err := apiClient.TaskCategoriesAPI.TaskCategoriesTaskCategoryIdPut(context.Background(), taskCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
