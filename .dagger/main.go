// This is the main orchestration dagger function.

package main

import (
	"context"
	"dagger/dagger/internal/dagger"
	"fmt"
)

type Dagger struct{}

// Test runs terraform plan and apply against localstack
// usage:  dagger call test --directory terraform
func (d *Dagger) Test(
	// +defaultPath="../terraform"
	directory *dagger.Directory,
) (string, error) {

	ctx := context.Background()

	localstack := dag.Localstack().Serve().WithHostname("localstack")

	localstack_endpoint, _ := localstack.Endpoint(ctx)

	terraform, err := dag.Container().
		From("hashicorp/terraform:latest").
		WithServiceBinding("localstack", localstack).
		WithEnvVariable("AWS_ENDPOINT_URL", fmt.Sprintf("http://%s", localstack_endpoint)).
		WithDirectory("/src", directory).
		WithWorkdir("/src").
		WithExec([]string{"terraform", "init"}).
		WithExec([]string{"terraform", "test"}).Stdout(ctx)

	if err != nil {
		return "", err
	}

	return terraform, nil
}

// Lint runs the pre-commit and commitlint linters
func (d *Dagger) Lint(
	// +defaultPath="../"
	directory *dagger.Directory,
) (string, error) {

	dir := directory.Filter(dagger.DirectoryFilterOpts{
		Include: []string{".git", "terraform", ".pre-commit-config.yaml", "commitlint.config.mjs"},
	})

	ctx := context.Background()

	type result struct {
		name string
		out  string
		err  error
	}

	results := make(chan result, 2)

	// Run Commitlint
	// Example using a public module in this module
	go func() {
		out, err := dag.Commitlint().Lint(dir, dagger.CommitlintLintOpts{
			Args: []string{"--last"},
		}).Stderr(ctx)
		results <- result{name: "commitlint", out: out, err: err}
	}()

	// Run PreCommit
	// Example using a local module in this module
	go func() {
		out, err := dag.PreCommit().Run(ctx, dir)
		results <- result{name: "precommit", out: out, err: err}
	}()

	// Collect both results
	var combinedOutput string
	var returnErr error
	for i := 0; i < 2; i++ {
		res := <-results
		if res.err != nil {
			returnErr = res.err
		}
		combinedOutput += fmt.Sprintf("== %s output ==\n%s\n", res.name, res.out)
	}

	return combinedOutput, returnErr
}
