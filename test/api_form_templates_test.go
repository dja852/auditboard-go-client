/*
AuditBoard Developer Portal API Documentation

Testing FormTemplatesAPIService

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

func Test_auditboard_FormTemplatesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test FormTemplatesAPIService FormTemplatesFormTemplateIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formTemplateId int64

		resp, httpRes, err := apiClient.FormTemplatesAPI.FormTemplatesFormTemplateIdDelete(context.Background(), formTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FormTemplatesAPIService FormTemplatesFormTemplateIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formTemplateId int64

		resp, httpRes, err := apiClient.FormTemplatesAPI.FormTemplatesFormTemplateIdGet(context.Background(), formTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FormTemplatesAPIService FormTemplatesFormTemplateIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var formTemplateId int64

		resp, httpRes, err := apiClient.FormTemplatesAPI.FormTemplatesFormTemplateIdPut(context.Background(), formTemplateId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FormTemplatesAPIService FormTemplatesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FormTemplatesAPI.FormTemplatesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test FormTemplatesAPIService FormTemplatesPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.FormTemplatesAPI.FormTemplatesPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
