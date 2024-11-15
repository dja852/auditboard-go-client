/*
AuditBoard Developer Portal API Documentation

Testing QuestionResponsesAPIService

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

func Test_auditboard_QuestionResponsesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test QuestionResponsesAPIService QuestionResponsesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.QuestionResponsesAPI.QuestionResponsesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuestionResponsesAPIService QuestionResponsesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.QuestionResponsesAPI.QuestionResponsesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuestionResponsesAPIService QuestionResponsesQuestionResponseIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var questionResponseId int64

		resp, httpRes, err := apiClient.QuestionResponsesAPI.QuestionResponsesQuestionResponseIdDelete(context.Background(), questionResponseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuestionResponsesAPIService QuestionResponsesQuestionResponseIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var questionResponseId int64

		resp, httpRes, err := apiClient.QuestionResponsesAPI.QuestionResponsesQuestionResponseIdGet(context.Background(), questionResponseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test QuestionResponsesAPIService QuestionResponsesQuestionResponseIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var questionResponseId int64

		resp, httpRes, err := apiClient.QuestionResponsesAPI.QuestionResponsesQuestionResponseIdPut(context.Background(), questionResponseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
