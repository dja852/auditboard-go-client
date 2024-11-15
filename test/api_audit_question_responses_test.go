/*
AuditBoard Developer Portal API Documentation

Testing AuditQuestionResponsesAPIService

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

func Test_auditboard_AuditQuestionResponsesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AuditQuestionResponsesAPIService AuditQuestionResponsesAuditQuestionResponseIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditQuestionResponseId int64

		resp, httpRes, err := apiClient.AuditQuestionResponsesAPI.AuditQuestionResponsesAuditQuestionResponseIdGet(context.Background(), auditQuestionResponseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditQuestionResponsesAPIService AuditQuestionResponsesAuditQuestionResponseIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var auditQuestionResponseId int64

		resp, httpRes, err := apiClient.AuditQuestionResponsesAPI.AuditQuestionResponsesAuditQuestionResponseIdPut(context.Background(), auditQuestionResponseId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AuditQuestionResponsesAPIService AuditQuestionResponsesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AuditQuestionResponsesAPI.AuditQuestionResponsesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
