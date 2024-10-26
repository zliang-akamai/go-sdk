/*
Linode API

Testing ManagedAPIService

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

func Test_openapi_ManagedAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ManagedAPIService DeleteManagedContact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var contactId int32

		resp, httpRes, err := apiClient.ManagedAPI.DeleteManagedContact(context.Background(), apiVersion, contactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService DeleteManagedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var serviceId int32

		resp, httpRes, err := apiClient.ManagedAPI.DeleteManagedService(context.Background(), apiVersion, serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedContact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var contactId int32

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedContact(context.Background(), apiVersion, contactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedContacts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedContacts(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedCredential", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var credentialId int32

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedCredential(context.Background(), apiVersion, credentialId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedCredentials(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedIssue", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var issueId int32

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedIssue(context.Background(), apiVersion, issueId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedIssues", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedIssues(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedLinodeSetting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var linodeId int32

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedLinodeSetting(context.Background(), apiVersion, linodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedLinodeSettings", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedLinodeSettings(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var serviceId int32

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedService(context.Background(), apiVersion, serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedServices", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedServices(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedSshKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedSshKey(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService GetManagedStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ManagedAPI.GetManagedStats(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PostDisableManagedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var serviceId int32

		resp, httpRes, err := apiClient.ManagedAPI.PostDisableManagedService(context.Background(), apiVersion, serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PostEnableManagedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var serviceId int32

		resp, httpRes, err := apiClient.ManagedAPI.PostEnableManagedService(context.Background(), apiVersion, serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PostManagedContact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ManagedAPI.PostManagedContact(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PostManagedCredential", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ManagedAPI.PostManagedCredential(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PostManagedCredentialRevoke", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var credentialId int32

		resp, httpRes, err := apiClient.ManagedAPI.PostManagedCredentialRevoke(context.Background(), apiVersion, credentialId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PostManagedCredentialUsernamePassword", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var credentialId int32

		resp, httpRes, err := apiClient.ManagedAPI.PostManagedCredentialUsernamePassword(context.Background(), apiVersion, credentialId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PostManagedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ManagedAPI.PostManagedService(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PutManagedContact", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var contactId int32

		resp, httpRes, err := apiClient.ManagedAPI.PutManagedContact(context.Background(), apiVersion, contactId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PutManagedCredential", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var credentialId int32

		resp, httpRes, err := apiClient.ManagedAPI.PutManagedCredential(context.Background(), apiVersion, credentialId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PutManagedLinodeSetting", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var linodeId int32

		resp, httpRes, err := apiClient.ManagedAPI.PutManagedLinodeSetting(context.Background(), apiVersion, linodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ManagedAPIService PutManagedService", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var serviceId int32

		resp, httpRes, err := apiClient.ManagedAPI.PutManagedService(context.Background(), apiVersion, serviceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}