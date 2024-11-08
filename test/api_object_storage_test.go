/*
Linode API

Testing ObjectStorageAPIService

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

func Test_openapi_ObjectStorageAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ObjectStorageAPIService DeleteObjectStorageBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.DeleteObjectStorageBucket(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService DeleteObjectStorageKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var keyId int32

		resp, httpRes, err := apiClient.ObjectStorageAPI.DeleteObjectStorageKey(context.Background(), apiVersion, keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService DeleteObjectStorageSsl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.DeleteObjectStorageSsl(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageBucket(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageBucketAcl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageBucketAcl(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageBucketContent", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageBucketContent(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageBucketinCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageBucketinCluster(context.Background(), apiVersion, regionId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageBuckets", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageBuckets(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageCluster(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageClusters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageClusters(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var keyId int32

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageKey(context.Background(), apiVersion, keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageKeys(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageSsl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageSsl(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageTransfer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageTransfer(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService GetObjectStorageTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ObjectStorageAPI.GetObjectStorageTypes(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService PostCancelObjectStorage", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ObjectStorageAPI.PostCancelObjectStorage(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService PostObjectStorageBucket", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ObjectStorageAPI.PostObjectStorageBucket(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService PostObjectStorageBucketAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.PostObjectStorageBucketAccess(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService PostObjectStorageKeys", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.ObjectStorageAPI.PostObjectStorageKeys(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService PostObjectStorageObjectUrl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.PostObjectStorageObjectUrl(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService PostObjectStorageSsl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.PostObjectStorageSsl(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService PutObjectStorageBucketAcl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.PutObjectStorageBucketAcl(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService PutObjectStorageKey", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var keyId int32

		resp, httpRes, err := apiClient.ObjectStorageAPI.PutObjectStorageKey(context.Background(), apiVersion, keyId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ObjectStorageAPIService PutStorageBucketAccess", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var regionId string
		var bucket string

		resp, httpRes, err := apiClient.ObjectStorageAPI.PutStorageBucketAccess(context.Background(), apiVersion, regionId, bucket).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
