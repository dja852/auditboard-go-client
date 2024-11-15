/*
AuditBoard Developer Portal API Documentation

Testing AttachmentsAPIService

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

func Test_auditboard_AttachmentsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AttachmentsAPIService AttachmentsAttachmentIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var attachmentId int64

		resp, httpRes, err := apiClient.AttachmentsAPI.AttachmentsAttachmentIdDelete(context.Background(), attachmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsAPIService AttachmentsAttachmentIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var attachmentId int64

		resp, httpRes, err := apiClient.AttachmentsAPI.AttachmentsAttachmentIdGet(context.Background(), attachmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsAPIService AttachmentsAttachmentIdPut", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var attachmentId int64

		resp, httpRes, err := apiClient.AttachmentsAPI.AttachmentsAttachmentIdPut(context.Background(), attachmentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsAPIService AttachmentsGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AttachmentsAPI.AttachmentsGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AttachmentsAPIService AttachmentsPost", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.AttachmentsAPI.AttachmentsPost(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
