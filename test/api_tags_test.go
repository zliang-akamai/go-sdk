/*
Linode API

Testing TagsAPIService

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

func Test_openapi_TagsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test TagsAPIService DeleteAg", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var tagLabel string

		resp, httpRes, err := apiClient.TagsAPI.DeleteAg(context.Background(), apiVersion, tagLabel).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService GetTaggedObjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var tagLabel string

		resp, httpRes, err := apiClient.TagsAPI.GetTaggedObjects(context.Background(), apiVersion, tagLabel).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService GetTags", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.TagsAPI.GetTags(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test TagsAPIService PostTag", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.TagsAPI.PostTag(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
