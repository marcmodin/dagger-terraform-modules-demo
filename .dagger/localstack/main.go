// This module runs localstack as-a-service on port 4566

package main

import "dagger/localstack/internal/dagger"

type Localstack struct{}

// localStack returns a new localstack service
// exposed on port 4566
// usage:  dagger call serve up --ports 4566:4566
func (m *Localstack) Serve() *dagger.Service {
	return dag.Container().
		From("localstack/localstack:latest").
		WithExposedPort(4566).
		AsService()
}
