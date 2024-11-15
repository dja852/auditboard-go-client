/*
AuditBoard Developer Portal API Documentation

Testing EsgMetricValueAggregatesAPIService

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

func Test_auditboard_EsgMetricValueAggregatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test EsgMetricValueAggregatesAPIService EsgMetricValueAggregatesEsgMetricValueAggregateIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var esgMetricValueAggregateId int64

		resp, httpRes, err := apiClient.EsgMetricValueAggregatesAPI.EsgMetricValueAggregatesEsgMetricValueAggregateIdGet(context.Background(), esgMetricValueAggregateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test EsgMetricValueAggregatesAPIService EsgMetricValueAggregatesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.EsgMetricValueAggregatesAPI.EsgMetricValueAggregatesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
