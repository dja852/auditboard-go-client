/*
AuditBoard Developer Portal API Documentation

Testing RcwSchedulesAPIService

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

func Test_auditboard_RcwSchedulesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RcwSchedulesAPIService RcwSchedulesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RcwSchedulesAPI.RcwSchedulesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RcwSchedulesAPIService RcwSchedulesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RcwSchedulesAPI.RcwSchedulesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RcwSchedulesAPIService RcwSchedulesRcwScheduleIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rcwScheduleId int64

		resp, httpRes, err := apiClient.RcwSchedulesAPI.RcwSchedulesRcwScheduleIdDelete(context.Background(), rcwScheduleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RcwSchedulesAPIService RcwSchedulesRcwScheduleIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rcwScheduleId int64

		resp, httpRes, err := apiClient.RcwSchedulesAPI.RcwSchedulesRcwScheduleIdGet(context.Background(), rcwScheduleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RcwSchedulesAPIService RcwSchedulesRcwScheduleIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rcwScheduleId int64

		resp, httpRes, err := apiClient.RcwSchedulesAPI.RcwSchedulesRcwScheduleIdPut(context.Background(), rcwScheduleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}