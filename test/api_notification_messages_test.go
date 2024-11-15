/*
AuditBoard Developer Portal API Documentation

Testing NotificationMessagesAPIService

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

func Test_auditboard_NotificationMessagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NotificationMessagesAPIService NotificationMessagesGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.NotificationMessagesAPI.NotificationMessagesGet(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationMessagesAPIService NotificationMessagesNotificationMessageIdDelete", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var notificationMessageId int64

		resp, httpRes, err := apiClient.NotificationMessagesAPI.NotificationMessagesNotificationMessageIdDelete(context.Background(), notificationMessageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NotificationMessagesAPIService NotificationMessagesNotificationMessageIdGet", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var notificationMessageId int64

		resp, httpRes, err := apiClient.NotificationMessagesAPI.NotificationMessagesNotificationMessageIdGet(context.Background(), notificationMessageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
