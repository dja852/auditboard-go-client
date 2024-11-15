/*
AuditBoard Developer Portal API Documentation

Testing PoliciesAPIService

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

func Test_auditboard_PoliciesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test PoliciesAPIService PoliciesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PoliciesAPI.PoliciesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService PoliciesPolicyIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.PoliciesAPI.PoliciesPolicyIdDelete(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService PoliciesPolicyIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.PoliciesAPI.PoliciesPolicyIdGet(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService PoliciesPolicyIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var policyId int64

		resp, httpRes, err := apiClient.PoliciesAPI.PoliciesPolicyIdPut(context.Background(), policyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test PoliciesAPIService PoliciesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.PoliciesAPI.PoliciesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
