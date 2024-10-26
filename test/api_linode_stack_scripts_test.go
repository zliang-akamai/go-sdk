/*
Linode API

Testing LinodeStackScriptsAPIService

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

func Test_openapi_LinodeStackScriptsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LinodeStackScriptsAPIService DeleteStackScript", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var stackscriptId string

		resp, httpRes, err := apiClient.LinodeStackScriptsAPI.DeleteStackScript(context.Background(), apiVersion, stackscriptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeStackScriptsAPIService GetStackScript", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var stackscriptId string

		resp, httpRes, err := apiClient.LinodeStackScriptsAPI.GetStackScript(context.Background(), apiVersion, stackscriptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeStackScriptsAPIService GetStackScripts", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.LinodeStackScriptsAPI.GetStackScripts(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeStackScriptsAPIService PostAddStackScript", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.LinodeStackScriptsAPI.PostAddStackScript(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeStackScriptsAPIService PutStackScript", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var stackscriptId string

		resp, httpRes, err := apiClient.LinodeStackScriptsAPI.PutStackScript(context.Background(), apiVersion, stackscriptId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
