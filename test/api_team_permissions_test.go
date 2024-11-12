/*
AuditBoard Developer Portal API Documentation

Testing TeamPermissionsAPIService

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

func Test_auditboard_TeamPermissionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TeamPermissionsAPIService TeamPermissionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TeamPermissionsAPI.TeamPermissionsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamPermissionsAPIService TeamPermissionsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TeamPermissionsAPI.TeamPermissionsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamPermissionsAPIService TeamPermissionsTeamPermissionIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var teamPermissionId int64

		resp, httpRes, err := apiClient.TeamPermissionsAPI.TeamPermissionsTeamPermissionIdDelete(context.Background(), teamPermissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TeamPermissionsAPIService TeamPermissionsTeamPermissionIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var teamPermissionId int64

		resp, httpRes, err := apiClient.TeamPermissionsAPI.TeamPermissionsTeamPermissionIdGet(context.Background(), teamPermissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
