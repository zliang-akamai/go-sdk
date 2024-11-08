/*
Linode API

Testing LinodeKubernetesEngineLKEAPIService

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

func Test_openapi_LinodeKubernetesEngineLKEAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test LinodeKubernetesEngineLKEAPIService DeleteLkeCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeCluster(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService DeleteLkeClusterAcl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeClusterAcl(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService DeleteLkeClusterKubeconfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeClusterKubeconfig(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService DeleteLkeClusterNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32
		var nodeId string

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeClusterNode(context.Background(), apiVersion, clusterId, nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService DeleteLkeNodePool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32
		var poolId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeNodePool(context.Background(), apiVersion, clusterId, poolId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService DeleteLkeServiceToken", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.DeleteLkeServiceToken(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeCluster(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeClusterAcl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterAcl(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeClusterApiEndpoints", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterApiEndpoints(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeClusterDashboard", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterDashboard(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeClusterKubeconfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterKubeconfig(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeClusterNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32
		var nodeId string

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterNode(context.Background(), apiVersion, clusterId, nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeClusterPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusterPools(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeClusters", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeClusters(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeNodePool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32
		var poolId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeNodePool(context.Background(), apiVersion, clusterId, poolId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeTypes(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeVersion", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var version string

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeVersion(context.Background(), apiVersion, version).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService GetLkeVersions", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.GetLkeVersions(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService PostLkeCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeCluster(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService PostLkeClusterNodeRecycle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32
		var nodeId string

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeClusterNodeRecycle(context.Background(), apiVersion, clusterId, nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService PostLkeClusterPoolRecycle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32
		var poolId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeClusterPoolRecycle(context.Background(), apiVersion, clusterId, poolId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService PostLkeClusterPools", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeClusterPools(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService PostLkeClusterRecycle", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeClusterRecycle(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService PostLkeClusterRegenerate", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.PostLkeClusterRegenerate(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService PutLkeCluster", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.PutLkeCluster(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService PutLkeClusterAcl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.PutLkeClusterAcl(context.Background(), apiVersion, clusterId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test LinodeKubernetesEngineLKEAPIService PutLkeNodePool", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var clusterId int32
		var poolId int32

		resp, httpRes, err := apiClient.LinodeKubernetesEngineLKEAPI.PutLkeNodePool(context.Background(), apiVersion, clusterId, poolId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
