/*
AuditBoard Developer Portal API Documentation

Testing SubprocessesDataAPIService

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

func Test_auditboard_SubprocessesDataAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test SubprocessesDataAPIService SubprocessesDataGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SubprocessesDataAPI.SubprocessesDataGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubprocessesDataAPIService SubprocessesDataPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.SubprocessesDataAPI.SubprocessesDataPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubprocessesDataAPIService SubprocessesDataSubprocessesDatumIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subprocessesDatumId int64

		resp, httpRes, err := apiClient.SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdDelete(context.Background(), subprocessesDatumId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubprocessesDataAPIService SubprocessesDataSubprocessesDatumIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subprocessesDatumId int64

		resp, httpRes, err := apiClient.SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdGet(context.Background(), subprocessesDatumId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test SubprocessesDataAPIService SubprocessesDataSubprocessesDatumIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var subprocessesDatumId int64

		resp, httpRes, err := apiClient.SubprocessesDataAPI.SubprocessesDataSubprocessesDatumIdPut(context.Background(), subprocessesDatumId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}