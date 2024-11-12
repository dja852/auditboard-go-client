/*
AuditBoard Developer Portal API Documentation

Testing TestsAPIService

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

func Test_auditboard_TestsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TestsAPIService TestsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TestsAPI.TestsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.TestsAPI.TestsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdAmendPrimaryReviewPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdAmendPrimaryReviewPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdAmendSecondaryReviewPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdAmendSecondaryReviewPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdApprovePrimaryReviewPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdApprovePrimaryReviewPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdApproveSecondaryReviewPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdApproveSecondaryReviewPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdBackToTesterPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdBackToTesterPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdCompleteTestPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdCompleteTestPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdDelete(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdGet(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdMarkEffectivePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdMarkEffectivePost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdMarkIneffectivePost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdMarkIneffectivePost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdMarkNoPopulationPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdMarkNoPopulationPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdPut(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdReopenPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdReopenPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdSendToSecondaryReviewerPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdSendToSecondaryReviewerPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdStartReviewPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdStartReviewPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdStartTestPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdStartTestPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TestsAPIService TestsTestIdStopTestPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var testId int64

		resp, httpRes, err := apiClient.TestsAPI.TestsTestIdStopTestPost(context.Background(), testId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}