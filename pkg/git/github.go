package git

import (
	"context"
	"log"

	"github.com/google/go-github/github"
)

// verify github credentials
func Verify() bool {
	client := github.NewClient(nil)

	// list public repositories for org "github"
	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)
	log.Print(repos)
	if err != nil {
		return false
	} else {
		return true
	}
}
