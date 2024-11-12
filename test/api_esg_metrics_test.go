/*
AuditBoard Developer Portal API Documentation

Testing EsgMetricsAPIService

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

func Test_auditboard_EsgMetricsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EsgMetricsAPIService EsgMetricsEsgMetricIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgMetricId int64

		resp, httpRes, err := apiClient.EsgMetricsAPI.EsgMetricsEsgMetricIdDelete(context.Background(), esgMetricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricsAPIService EsgMetricsEsgMetricIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgMetricId int64

		resp, httpRes, err := apiClient.EsgMetricsAPI.EsgMetricsEsgMetricIdGet(context.Background(), esgMetricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricsAPIService EsgMetricsEsgMetricIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgMetricId int64

		resp, httpRes, err := apiClient.EsgMetricsAPI.EsgMetricsEsgMetricIdPut(context.Background(), esgMetricId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricsAPIService EsgMetricsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EsgMetricsAPI.EsgMetricsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricsAPIService EsgMetricsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EsgMetricsAPI.EsgMetricsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
