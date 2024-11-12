/*
AuditBoard Developer Portal API Documentation

Testing ActionPlansAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package auditboard

import (
	"context"
	"testing"

	openapiclient "github.com/dja852/auditboard-go-client"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_auditboard_ActionPlansAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ActionPlansAPIService ActionPlansActionPlanIdDelete", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var actionPlanId int64

		resp, httpRes, err := apiClient.ActionPlansAPI.ActionPlansActionPlanIdDelete(context.Background(), actionPlanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionPlansAPIService ActionPlansActionPlanIdGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var actionPlanId int64

		resp, httpRes, err := apiClient.ActionPlansAPI.ActionPlansActionPlanIdGet(context.Background(), actionPlanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionPlansAPIService ActionPlansActionPlanIdPut", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		var actionPlanId int64

		resp, httpRes, err := apiClient.ActionPlansAPI.ActionPlansActionPlanIdPut(context.Background(), actionPlanId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionPlansAPIService ActionPlansGet", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ActionPlansAPI.ActionPlansGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ActionPlansAPIService ActionPlansPost", func(t *testing.T) {

		t.Skip("skip test") // remove to run test

		resp, httpRes, err := apiClient.ActionPlansAPI.ActionPlansPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
