/*
AuditBoard Developer Portal API Documentation

Testing AuditRotationSchedulesAPIService

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

func Test_auditboard_AuditRotationSchedulesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuditRotationSchedulesAPIService AuditRotationSchedulesAuditRotationScheduleIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditRotationScheduleId int64

		resp, httpRes, err := apiClient.AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdDelete(context.Background(), auditRotationScheduleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditRotationSchedulesAPIService AuditRotationSchedulesAuditRotationScheduleIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditRotationScheduleId int64

		resp, httpRes, err := apiClient.AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdGet(context.Background(), auditRotationScheduleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditRotationSchedulesAPIService AuditRotationSchedulesAuditRotationScheduleIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditRotationScheduleId int64

		resp, httpRes, err := apiClient.AuditRotationSchedulesAPI.AuditRotationSchedulesAuditRotationScheduleIdPut(context.Background(), auditRotationScheduleId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditRotationSchedulesAPIService AuditRotationSchedulesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuditRotationSchedulesAPI.AuditRotationSchedulesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditRotationSchedulesAPIService AuditRotationSchedulesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuditRotationSchedulesAPI.AuditRotationSchedulesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
