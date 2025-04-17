// This module creates a github pull-request comment
package main

import (
	"context"
	"dagger/comment/internal/dagger"
	"fmt"
	"strings"
)

type Comment struct {
	Commit string // +private
	Repo   string // +private
	// Token  *Secret // +private
	Token *dagger.Secret // +private
}

func normalizeRepoURL(repo string) string {
	// Remove known prefixes/suffixes
	repo = strings.TrimPrefix(repo, "https://github.com/")
	repo = strings.TrimPrefix(repo, "github.com/")
	repo = strings.TrimSuffix(repo, ".git")
	return repo
}

func New(
	// Comment on the given commit
	// +optional
	commit string,
	// The github repository
	// Supported formats:
	// - github.com/dagger/dagger
	// - dagger/dagger
	// - https://github.com/dagger/dagger
	// - https://github.com/dagger/dagger.git
	// +optional
	repo string,
	// The github token
	// +optional
	token *dagger.Secret,
) *Comment {
	return &Comment{
		Commit: commit,
		Repo:   normalizeRepoURL(repo),
	}
}

// Post creates a github comment on the given commit
func (c *Comment) Post(ctx context.Context, body string) (string, error) {
	message := fmt.Sprintf("Commenting on commit %s in repo %s: payload: %s", c.Commit, c.Repo, body)
	return message, nil
}
