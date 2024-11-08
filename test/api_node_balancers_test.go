/*
Linode API

Testing NodeBalancersAPIService

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

func Test_openapi_NodeBalancersAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test NodeBalancersAPIService DeleteNodeBalancer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.DeleteNodeBalancer(context.Background(), apiVersion, nodeBalancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService DeleteNodeBalancerConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32
		var configId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.DeleteNodeBalancerConfig(context.Background(), apiVersion, nodeBalancerId, configId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService DeleteNodeBalancerConfigNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32
		var configId int32
		var nodeId string

		resp, httpRes, err := apiClient.NodeBalancersAPI.DeleteNodeBalancerConfigNode(context.Background(), apiVersion, nodeBalancerId, configId, nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService GetNodeBalancer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.GetNodeBalancer(context.Background(), apiVersion, nodeBalancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService GetNodeBalancerConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32
		var configId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.GetNodeBalancerConfig(context.Background(), apiVersion, nodeBalancerId, configId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService GetNodeBalancerConfigNodes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32
		var configId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.GetNodeBalancerConfigNodes(context.Background(), apiVersion, nodeBalancerId, configId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService GetNodeBalancerConfigs", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.GetNodeBalancerConfigs(context.Background(), apiVersion, nodeBalancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService GetNodeBalancerFirewalls", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.GetNodeBalancerFirewalls(context.Background(), apiVersion, nodeBalancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService GetNodeBalancerNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32
		var configId int32
		var nodeId string

		resp, httpRes, err := apiClient.NodeBalancersAPI.GetNodeBalancerNode(context.Background(), apiVersion, nodeBalancerId, configId, nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService GetNodeBalancerStats", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.GetNodeBalancerStats(context.Background(), apiVersion, nodeBalancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService GetNodeBalancerTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.NodeBalancersAPI.GetNodeBalancerTypes(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService GetNodeBalancers", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.NodeBalancersAPI.GetNodeBalancers(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService PostNodeBalancer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.NodeBalancersAPI.PostNodeBalancer(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService PostNodeBalancerConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.PostNodeBalancerConfig(context.Background(), apiVersion, nodeBalancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService PostNodeBalancerNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32
		var configId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.PostNodeBalancerNode(context.Background(), apiVersion, nodeBalancerId, configId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService PostRebuildNodeBalancerConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32
		var configId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.PostRebuildNodeBalancerConfig(context.Background(), apiVersion, nodeBalancerId, configId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService PutNodeBalancer", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.PutNodeBalancer(context.Background(), apiVersion, nodeBalancerId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService PutNodeBalancerConfig", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32
		var configId int32

		resp, httpRes, err := apiClient.NodeBalancersAPI.PutNodeBalancerConfig(context.Background(), apiVersion, nodeBalancerId, configId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test NodeBalancersAPIService PutNodeBalancerNode", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var nodeBalancerId int32
		var configId int32
		var nodeId string

		resp, httpRes, err := apiClient.NodeBalancersAPI.PutNodeBalancerNode(context.Background(), apiVersion, nodeBalancerId, configId, nodeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}
