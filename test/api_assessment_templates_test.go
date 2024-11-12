/*
AuditBoard Developer Portal API Documentation

Testing AssessmentTemplatesAPIService

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

func Test_auditboard_AssessmentTemplatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AssessmentTemplatesAPIService AssessmentTemplatesAssessmentTemplateIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assessmentTemplateId int64

		resp, httpRes, err := apiClient.AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdDelete(context.Background(), assessmentTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssessmentTemplatesAPIService AssessmentTemplatesAssessmentTemplateIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assessmentTemplateId int64

		resp, httpRes, err := apiClient.AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdGet(context.Background(), assessmentTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssessmentTemplatesAPIService AssessmentTemplatesAssessmentTemplateIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var assessmentTemplateId int64

		resp, httpRes, err := apiClient.AssessmentTemplatesAPI.AssessmentTemplatesAssessmentTemplateIdPut(context.Background(), assessmentTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssessmentTemplatesAPIService AssessmentTemplatesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AssessmentTemplatesAPI.AssessmentTemplatesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AssessmentTemplatesAPIService AssessmentTemplatesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AssessmentTemplatesAPI.AssessmentTemplatesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
