/*
AuditBoard Developer Portal API Documentation

Testing RolePermissionsAPIService

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

func Test_auditboard_RolePermissionsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test RolePermissionsAPIService RolePermissionsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.RolePermissionsAPI.RolePermissionsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolePermissionsAPIService RolePermissionsRolePermissionIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rolePermissionId int64

		resp, httpRes, err := apiClient.RolePermissionsAPI.RolePermissionsRolePermissionIdGet(context.Background(), rolePermissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test RolePermissionsAPIService RolePermissionsRolePermissionIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var rolePermissionId int64

		resp, httpRes, err := apiClient.RolePermissionsAPI.RolePermissionsRolePermissionIdPut(context.Background(), rolePermissionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
