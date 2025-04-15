// This module runs pre-commit for terraform in specified root directory

package main

import (
	"context"
	"dagger/pre-commit/internal/dagger"
	"fmt"
	"log/slog"
)

// PreCommit is a struct that holds the Dagger container for pre-commit-terraform
type PreCommit struct {
	container *dagger.Container
}

var preCommitBuildArgs = []dagger.BuildArg{
	{Name: "PRE_COMMIT_VERSION", Value: "latest"},
	{Name: "OPENTOFU_VERSION", Value: "false"},
	{Name: "TERRAFORM_VERSION", Value: "latest"},
	{Name: "CHECKOV_VERSION", Value: "latest"},
	{Name: "HCLEDIT_VERSION", Value: "false"},
	{Name: "INFRACOST_VERSION", Value: "false"},
	{Name: "TERRAFORM_DOCS_VERSION", Value: "latest"},
	{Name: "TERRAGRUNT_VERSION", Value: "false"},
	{Name: "TERRASCAN_VERSION", Value: "false"},
	{Name: "TFLINT_VERSION", Value: "latest"},
	{Name: "TFSEC_VERSION", Value: "false"},
	{Name: "TFUPDATE_VERSION", Value: "false"},
	{Name: "TRIVY_VERSION", Value: "false"},
}

// SetContainer sets up the Dagger container for pre-commit-terraform
func SetContainer() *dagger.Container {
	repo := dag.Git("https://github.com/antonbabenko/pre-commit-terraform.git").
		Tag("v1.98.1").
		Tree().Filter(dagger.DirectoryFilterOpts{
		Include: []string{"tools", "Dockerfile"},
	})

	builder := dag.Container().
		Build(repo, dagger.ContainerBuildOpts{
			Dockerfile: "Dockerfile",
			BuildArgs:  preCommitBuildArgs,
			Target:     "builder",
		})

	// Reproduce the final image setup by copying manually from builder stage
	return dag.Container().
		From("python:3.12.0-alpine3.17").
		WithExec([]string{"apk", "add", "git", "bash"}).
		WithDirectory("/usr/bin", builder.Directory("/bin_dir")).
		WithDirectory("/usr/bin/", builder.Directory("/usr/local/bin")).
		WithDirectory("/usr/local/lib/python3.12/site-packages", builder.Directory("/usr/local/lib/python3.12/site-packages")).
		WithDirectory("/root", builder.Directory("/root")).
		WithExec([]string{"chmod", "-R", "+x", "/usr/bin"})
}

func (m *PreCommit) initContainer() {
	slog.Debug("Building pre-commit container")
	m.container = SetContainer()
}

// Run Pre-Commit-Terraform
// example usage: dagger call run
func (m *PreCommit) Run(ctx context.Context, directory *dagger.Directory) (string, error) {

	if directory == nil {
		return "", fmt.Errorf("Run: input directory is nil")
	}

	m.initContainer()
	slog.Debug("Running pre-commit-terraform")

	return m.container.
		WithDirectory("/src", directory).
		WithWorkdir("/src").
		WithExec([]string{"pre-commit", "run", "-a"}).
		Stdout(ctx)
}

// Open interactive debug shell
// example usage: dagger shell debug
func (m *PreCommit) Debug() *dagger.Container {
	return m.container.Terminal() // Return the base container for debugging
}
