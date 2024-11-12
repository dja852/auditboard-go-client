/*
AuditBoard Developer Portal API Documentation

Testing RiskTypesAPIService

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

func Test_auditboard_RiskTypesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RiskTypesAPIService RiskTypesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RiskTypesAPI.RiskTypesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskTypesAPIService RiskTypesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RiskTypesAPI.RiskTypesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskTypesAPIService RiskTypesRiskTypeIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var riskTypeId int64

		resp, httpRes, err := apiClient.RiskTypesAPI.RiskTypesRiskTypeIdDelete(context.Background(), riskTypeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskTypesAPIService RiskTypesRiskTypeIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var riskTypeId int64

		resp, httpRes, err := apiClient.RiskTypesAPI.RiskTypesRiskTypeIdGet(context.Background(), riskTypeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RiskTypesAPIService RiskTypesRiskTypeIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var riskTypeId int64

		resp, httpRes, err := apiClient.RiskTypesAPI.RiskTypesRiskTypeIdPut(context.Background(), riskTypeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
