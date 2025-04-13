// A generated module for Pctf functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/pctf/internal/dagger"
)

type Pctf struct {
	ctr *dagger.Container
	src *dagger.Directory
}

func (m *Pctf) SetContainer() {
	if m.ctr == nil {
		repo := dag.Git("https://github.com/antonbabenko/pre-commit-terraform.git").
			Tag("v1.98.1").
			Tree().Filter(dagger.DirectoryFilterOpts{
			Include: []string{"tools", "Dockerfile"},
		})

		buildArgs := []dagger.BuildArg{
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

		builder := dag.Container().
			Build(repo, dagger.ContainerBuildOpts{
				Dockerfile: "Dockerfile",
				BuildArgs:  buildArgs,
				Target:     "builder",
			})

		// Reproduce the final image setup by copying manually from builder stage
		m.ctr = dag.Container().
			From("python:3.12.0-alpine3.17").
			WithExec([]string{"apk", "add", "git"}).
			WithDirectory("/usr/bin", builder.Directory("/bin_dir")).
			WithDirectory("/usr/bin/", builder.Directory("/usr/local/bin")).
			WithDirectory("/usr/local/lib/python3.12/site-packages", builder.Directory("/usr/local/lib/python3.12/site-packages")).
			WithDirectory("/root", builder.Directory("/root")).
			WithExec([]string{"chmod", "-R", "+x", "/usr/bin"})
	}
}

// +ignore=["*", "!.git","!terraform"]

// Run pre-commit in source directory
func (m *Pctf) Run(
	ctx context.Context,
	// +ignore=["*", "!.git","!terraform", "!.pre-commit-config.yaml"]
	source *dagger.Directory,
) *dagger.Container {
	m.SetContainer()
	return m.ctr.
		WithDirectory("/src", source).
		WithWorkdir("/src").
		// WithExec([]string{"pre-commit", "run", "-a"}).
		Terminal()
	// Stdout(ctx)
}

// Open interactive debug shell
// example usage: dagger shell debug
func (m *Pctf) Debug() *dagger.Container {
	m.SetContainer()
	return m.ctr.Terminal() // Return the base container for debugging
}
