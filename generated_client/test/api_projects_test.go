/*
NuoDB Control Plane REST API

Testing ProjectsAPIService

*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech);

package nuodbaas

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	openapiclient "github.com/nuodb/nuodbaas-tf-plugin/generated_client"
)

func Test_nuodbaas_ProjectsAPIService(t *testing.T) {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)

	t.Run("Test ProjectsAPIService CreateProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string

		httpRes, err := apiClient.ProjectsAPI.CreateProject(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService DeleteProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string

		httpRes, err := apiClient.ProjectsAPI.DeleteProject(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetAllProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		resp, httpRes, err := apiClient.ProjectsAPI.GetAllProjects(context.Background()).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string

		resp, httpRes, err := apiClient.ProjectsAPI.GetProject(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService GetProjects", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string

		resp, httpRes, err := apiClient.ProjectsAPI.GetProjects(context.Background(), organization).Execute()

		require.Nil(t, err)
		require.NotNil(t, resp)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

	t.Run("Test ProjectsAPIService PatchProject", func(t *testing.T) {

		t.Skip("skip test")  // remove to run test

		var organization string
		var project string

		httpRes, err := apiClient.ProjectsAPI.PatchProject(context.Background(), organization, project).Execute()

		require.Nil(t, err)
		assert.Equal(t, 200, httpRes.StatusCode)

	})

}