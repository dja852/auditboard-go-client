/*
AuditBoard Developer Portal API Documentation

Testing ControlsSegmentsAPIService

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

func Test_auditboard_ControlsSegmentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ControlsSegmentsAPIService ControlsSegmentsControlsSegmentIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var controlsSegmentId int64

		resp, httpRes, err := apiClient.ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdDelete(context.Background(), controlsSegmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ControlsSegmentsAPIService ControlsSegmentsControlsSegmentIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var controlsSegmentId int64

		resp, httpRes, err := apiClient.ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdGet(context.Background(), controlsSegmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ControlsSegmentsAPIService ControlsSegmentsControlsSegmentIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var controlsSegmentId int64

		resp, httpRes, err := apiClient.ControlsSegmentsAPI.ControlsSegmentsControlsSegmentIdPut(context.Background(), controlsSegmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ControlsSegmentsAPIService ControlsSegmentsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ControlsSegmentsAPI.ControlsSegmentsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ControlsSegmentsAPIService ControlsSegmentsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ControlsSegmentsAPI.ControlsSegmentsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}