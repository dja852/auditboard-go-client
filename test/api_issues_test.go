/*
AuditBoard Developer Portal API Documentation

Testing IssuesAPIService

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

func Test_auditboard_IssuesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test IssuesAPIService IssuesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IssuesAPI.IssuesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssuesAPIService IssuesIssueIdClosePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issueId int64

		resp, httpRes, err := apiClient.IssuesAPI.IssuesIssueIdClosePost(context.Background(), issueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssuesAPIService IssuesIssueIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issueId int64

		resp, httpRes, err := apiClient.IssuesAPI.IssuesIssueIdDelete(context.Background(), issueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssuesAPIService IssuesIssueIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issueId int64

		resp, httpRes, err := apiClient.IssuesAPI.IssuesIssueIdGet(context.Background(), issueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssuesAPIService IssuesIssueIdMovePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issueId int64

		resp, httpRes, err := apiClient.IssuesAPI.IssuesIssueIdMovePost(context.Background(), issueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssuesAPIService IssuesIssueIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issueId int64

		resp, httpRes, err := apiClient.IssuesAPI.IssuesIssueIdPut(context.Background(), issueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssuesAPIService IssuesIssueIdRemediatePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var issueId int64

		resp, httpRes, err := apiClient.IssuesAPI.IssuesIssueIdRemediatePost(context.Background(), issueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test IssuesAPIService IssuesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.IssuesAPI.IssuesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}