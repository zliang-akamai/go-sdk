/*
Linode API

Testing DatabasesAPIService

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

func Test_openapi_DatabasesAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test DatabasesAPIService DeleteDatabaseMysqlInstanceBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32
		var backupId int32

		resp, httpRes, err := apiClient.DatabasesAPI.DeleteDatabaseMysqlInstanceBackup(context.Background(), apiVersion, instanceId, backupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService DeleteDatabasePostgreSqlInstanceBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32
		var backupId int32

		resp, httpRes, err := apiClient.DatabasesAPI.DeleteDatabasePostgreSqlInstanceBackup(context.Background(), apiVersion, instanceId, backupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService DeleteDatabasesMysqlInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.DeleteDatabasesMysqlInstance(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService DeleteDatabasesPostgreSqlInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.DeleteDatabasesPostgreSqlInstance(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesEngine", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var engineId string

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesEngine(context.Background(), apiVersion, engineId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesEngines", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesEngines(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesInstances(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesMysqlInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstance(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesMysqlInstanceBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32
		var backupId int32

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstanceBackup(context.Background(), apiVersion, instanceId, backupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesMysqlInstanceBackups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstanceBackups(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesMysqlInstanceCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstanceCredentials(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesMysqlInstanceSsl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstanceSsl(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesMysqlInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesMysqlInstances(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesPostgreSqlInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesPostgreSqlInstance(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesPostgreSqlInstanceBackups", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesPostgreSqlInstanceBackups(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesPostgreSqlInstanceCredentials", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesPostgreSqlInstanceCredentials(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesPostgreSqlInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesPostgreSqlInstances(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesPostgresqlInstanceBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32
		var backupId int32

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesPostgresqlInstanceBackup(context.Background(), apiVersion, instanceId, backupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesPostgresqlInstanceSsl", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesPostgresqlInstanceSsl(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesType", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var typeId string

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesType(context.Background(), apiVersion, typeId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService GetDatabasesTypes", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.DatabasesAPI.GetDatabasesTypes(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PostDatabasesMysqlInstanceBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.PostDatabasesMysqlInstanceBackup(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PostDatabasesMysqlInstanceBackupRestore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32
		var backupId int32

		resp, httpRes, err := apiClient.DatabasesAPI.PostDatabasesMysqlInstanceBackupRestore(context.Background(), apiVersion, instanceId, backupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PostDatabasesMysqlInstanceCredentialsReset", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.PostDatabasesMysqlInstanceCredentialsReset(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PostDatabasesMysqlInstancePatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.PostDatabasesMysqlInstancePatch(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PostDatabasesMysqlInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.DatabasesAPI.PostDatabasesMysqlInstances(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PostDatabasesPostgreSqlInstanceBackup", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.PostDatabasesPostgreSqlInstanceBackup(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PostDatabasesPostgreSqlInstanceBackupRestore", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32
		var backupId int32

		resp, httpRes, err := apiClient.DatabasesAPI.PostDatabasesPostgreSqlInstanceBackupRestore(context.Background(), apiVersion, instanceId, backupId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PostDatabasesPostgreSqlInstanceCredentialsReset", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.PostDatabasesPostgreSqlInstanceCredentialsReset(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PostDatabasesPostgreSqlInstancePatch", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.PostDatabasesPostgreSqlInstancePatch(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PostDatabasesPostgreSqlInstances", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string

		resp, httpRes, err := apiClient.DatabasesAPI.PostDatabasesPostgreSqlInstances(context.Background(), apiVersion).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PutDatabasesMysqlInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.PutDatabasesMysqlInstance(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test DatabasesAPIService PutDatabasesPostgreSqlInstance", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var apiVersion string
		var instanceId int32

		resp, httpRes, err := apiClient.DatabasesAPI.PutDatabasesPostgreSqlInstance(context.Background(), apiVersion, instanceId).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}