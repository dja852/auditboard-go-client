/*
AuditBoard Developer Portal API Documentation

Testing RiskCategoriesAPIService

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

func Test_auditboard_RiskCategoriesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RiskCategoriesAPIService RiskCategoriesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RiskCategoriesAPI.RiskCategoriesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskCategoriesAPIService RiskCategoriesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RiskCategoriesAPI.RiskCategoriesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskCategoriesAPIService RiskCategoriesRiskCategoryIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var riskCategoryId int64

		resp, httpRes, err := apiClient.RiskCategoriesAPI.RiskCategoriesRiskCategoryIdDelete(context.Background(), riskCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskCategoriesAPIService RiskCategoriesRiskCategoryIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var riskCategoryId int64

		resp, httpRes, err := apiClient.RiskCategoriesAPI.RiskCategoriesRiskCategoryIdGet(context.Background(), riskCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskCategoriesAPIService RiskCategoriesRiskCategoryIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var riskCategoryId int64

		resp, httpRes, err := apiClient.RiskCategoriesAPI.RiskCategoriesRiskCategoryIdPut(context.Background(), riskCategoryId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
