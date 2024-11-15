/*
AuditBoard Developer Portal API Documentation

Testing WorkspacesAPIService

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

func Test_auditboard_WorkspacesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test WorkspacesAPIService WorkspacesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkspacesAPI.WorkspacesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkspacesAPIService WorkspacesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.WorkspacesAPI.WorkspacesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkspacesAPIService WorkspacesWorkspaceIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspaceId int64

		resp, httpRes, err := apiClient.WorkspacesAPI.WorkspacesWorkspaceIdDelete(context.Background(), workspaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkspacesAPIService WorkspacesWorkspaceIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspaceId int64

		resp, httpRes, err := apiClient.WorkspacesAPI.WorkspacesWorkspaceIdGet(context.Background(), workspaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test WorkspacesAPIService WorkspacesWorkspaceIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var workspaceId int64

		resp, httpRes, err := apiClient.WorkspacesAPI.WorkspacesWorkspaceIdPut(context.Background(), workspaceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
