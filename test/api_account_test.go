/*
Linode API

Testing AccountAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package openapi

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func Test_openapi_AccountAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test AccountAPIService DeleteClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clientId string

		resp, httpRes, err := apiClient.AccountAPI.DeleteClient(context.Background(), apiVersion, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService DeleteEntityTransfer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var token string

		resp, httpRes, err := apiClient.AccountAPI.DeleteEntityTransfer(context.Background(), apiVersion, token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService DeletePaymentMethod", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var paymentMethodId int32

		resp, httpRes, err := apiClient.AccountAPI.DeletePaymentMethod(context.Background(), apiVersion, paymentMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService DeleteServiceTransfer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var token string

		resp, httpRes, err := apiClient.AccountAPI.DeleteServiceTransfer(context.Background(), apiVersion, token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService DeleteUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var username string

		resp, httpRes, err := apiClient.AccountAPI.DeleteUser(context.Background(), apiVersion, username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetAccount(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetAccountAgreements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetAccountAgreements(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetAccountAvailability", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var id string

		resp, httpRes, err := apiClient.AccountAPI.GetAccountAvailability(context.Background(), apiVersion, id).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetAccountLogin", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var loginId int32

		resp, httpRes, err := apiClient.AccountAPI.GetAccountLogin(context.Background(), apiVersion, loginId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetAccountLogins", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetAccountLogins(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetAccountSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetAccountSettings(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetAvailability", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetAvailability(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetChildAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var euuid string

		resp, httpRes, err := apiClient.AccountAPI.GetChildAccount(context.Background(), apiVersion, euuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetChildAccounts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetChildAccounts(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clientId string

		resp, httpRes, err := apiClient.AccountAPI.GetClient(context.Background(), apiVersion, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetClientThumbnail", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clientId string

		resp, httpRes, err := apiClient.AccountAPI.GetClientThumbnail(context.Background(), apiVersion, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetClients", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetClients(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetEnrolledBetaProgram", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var betaId string

		resp, httpRes, err := apiClient.AccountAPI.GetEnrolledBetaProgram(context.Background(), apiVersion, betaId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetEnrolledBetaPrograms", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetEnrolledBetaPrograms(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetEntityTransfer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var token string

		resp, httpRes, err := apiClient.AccountAPI.GetEntityTransfer(context.Background(), apiVersion, token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetEntityTransfers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetEntityTransfers(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetEvent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var eventId int32

		resp, httpRes, err := apiClient.AccountAPI.GetEvent(context.Background(), apiVersion, eventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetEvents", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetEvents(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetInvoice", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var invoiceId int32

		resp, httpRes, err := apiClient.AccountAPI.GetInvoice(context.Background(), apiVersion, invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetInvoiceItems", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var invoiceId int32

		resp, httpRes, err := apiClient.AccountAPI.GetInvoiceItems(context.Background(), apiVersion, invoiceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetInvoices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetInvoices(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetMaintenance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetMaintenance(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetNotifications", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetNotifications(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetPayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var paymentId int32

		resp, httpRes, err := apiClient.AccountAPI.GetPayment(context.Background(), apiVersion, paymentId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetPaymentMethod", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var paymentMethodId int32

		resp, httpRes, err := apiClient.AccountAPI.GetPaymentMethod(context.Background(), apiVersion, paymentMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetPaymentMethods", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetPaymentMethods(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetPayments", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetPayments(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetServiceTransfer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var token string

		resp, httpRes, err := apiClient.AccountAPI.GetServiceTransfer(context.Background(), apiVersion, token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetServiceTransfers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetServiceTransfers(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetTransfer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetTransfer(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var username string

		resp, httpRes, err := apiClient.AccountAPI.GetUser(context.Background(), apiVersion, username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetUserGrants", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var username string

		resp, httpRes, err := apiClient.AccountAPI.GetUserGrants(context.Background(), apiVersion, username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService GetUsers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.GetUsers(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostAcceptEntityTransfer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var token string

		resp, httpRes, err := apiClient.AccountAPI.PostAcceptEntityTransfer(context.Background(), apiVersion, token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostAcceptServiceTransfer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var token string

		resp, httpRes, err := apiClient.AccountAPI.PostAcceptServiceTransfer(context.Background(), apiVersion, token).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostAccountAgreements", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostAccountAgreements(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostBetaProgram", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostBetaProgram(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostCancelAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostCancelAccount(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostChildAccountToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var euuid string

		resp, httpRes, err := apiClient.AccountAPI.PostChildAccountToken(context.Background(), apiVersion, euuid).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostClient(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostCreditCard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostCreditCard(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostEnableAccountManaged", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostEnableAccountManaged(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostEntityTransfer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostEntityTransfer(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostEventRead", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var eventId int32

		resp, httpRes, err := apiClient.AccountAPI.PostEventRead(context.Background(), apiVersion, eventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostEventSeen", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var eventId int32

		resp, httpRes, err := apiClient.AccountAPI.PostEventSeen(context.Background(), apiVersion, eventId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostExecutePayPalPayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostExecutePayPalPayment(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostMakePaymentMethodDefault", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var paymentMethodId int32

		resp, httpRes, err := apiClient.AccountAPI.PostMakePaymentMethodDefault(context.Background(), apiVersion, paymentMethodId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostPayPalPayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostPayPalPayment(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostPayment", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostPayment(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostPaymentMethod", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostPaymentMethod(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostPromoCredit", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostPromoCredit(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostResetClientSecret", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clientId string

		resp, httpRes, err := apiClient.AccountAPI.PostResetClientSecret(context.Background(), apiVersion, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostServiceTransfer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostServiceTransfer(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PostUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PostUser(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PutAccount", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PutAccount(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PutAccountSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.AccountAPI.PutAccountSettings(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PutClient", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clientId string

		resp, httpRes, err := apiClient.AccountAPI.PutClient(context.Background(), apiVersion, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PutClientThumbnail", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clientId string

		resp, httpRes, err := apiClient.AccountAPI.PutClientThumbnail(context.Background(), apiVersion, clientId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PutUser", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var username string

		resp, httpRes, err := apiClient.AccountAPI.PutUser(context.Background(), apiVersion, username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test AccountAPIService PutUserGrants", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var username string

		resp, httpRes, err := apiClient.AccountAPI.PutUserGrants(context.Background(), apiVersion, username).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}