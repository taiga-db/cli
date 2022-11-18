package config_tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestJobAndPipelineDevelopment(t *testing.T) {
	root := loadEnvironment(t, "./job_and_pipeline", "development")
	assert.Len(t, root.Resources.Jobs, 0)
	assert.Len(t, root.Resources.Pipelines, 1)

	p := root.Resources.Pipelines["nyc_taxi_pipeline"]
	assert.True(t, p.Development)
	require.Len(t, p.Libraries, 1)
	assert.Equal(t, "./dlt/nyc_taxi_loader", p.Libraries[0].Notebook.Path)
	assert.Equal(t, "nyc_taxi_development", p.Target)
}

func TestJobAndPipelineStaging(t *testing.T) {
	root := loadEnvironment(t, "./job_and_pipeline", "staging")
	assert.Len(t, root.Resources.Jobs, 0)
	assert.Len(t, root.Resources.Pipelines, 1)

	p := root.Resources.Pipelines["nyc_taxi_pipeline"]
	assert.False(t, p.Development)
	require.Len(t, p.Libraries, 1)
	assert.Equal(t, "./dlt/nyc_taxi_loader", p.Libraries[0].Notebook.Path)
	assert.Equal(t, "nyc_taxi_staging", p.Target)
}

func TestJobAndPipelineProduction(t *testing.T) {
	root := loadEnvironment(t, "./job_and_pipeline", "production")
	assert.Len(t, root.Resources.Jobs, 1)
	assert.Len(t, root.Resources.Pipelines, 1)

	p := root.Resources.Pipelines["nyc_taxi_pipeline"]
	assert.False(t, p.Development)
	require.Len(t, p.Libraries, 1)
	assert.Equal(t, "./dlt/nyc_taxi_loader", p.Libraries[0].Notebook.Path)
	assert.Equal(t, "nyc_taxi_production", p.Target)

	j := root.Resources.Jobs["pipeline_schedule"]
	assert.Equal(t, "Daily refresh of production pipeline", j.Name)
	require.Len(t, j.Tasks, 1)
	assert.NotEmpty(t, j.Tasks[0].PipelineTask.PipelineId)
}
