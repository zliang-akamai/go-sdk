/*
Linode API

Testing ImagesAPIService

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

func Test_openapi_ImagesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ImagesAPIService DeleteImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var imageId string

		resp, httpRes, err := apiClient.ImagesAPI.DeleteImage(context.Background(), apiVersion, imageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService GetImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var imageId string

		resp, httpRes, err := apiClient.ImagesAPI.GetImage(context.Background(), apiVersion, imageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService GetImages", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ImagesAPI.GetImages(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService PostImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ImagesAPI.PostImage(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService PostReplicateImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var imageId string

		resp, httpRes, err := apiClient.ImagesAPI.PostReplicateImage(context.Background(), apiVersion, imageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService PostUploadImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ImagesAPI.PostUploadImage(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ImagesAPIService PutImage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var imageId string

		resp, httpRes, err := apiClient.ImagesAPI.PutImage(context.Background(), apiVersion, imageId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
