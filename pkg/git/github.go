package git

import (
	"context"
	"log"
	"os/exec"

	"github.com/google/go-github/github"
	"github.com/lysimon/terrabender/pkg/config"
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

func Clone(gihubRepository string, branch string, finalPath string) (string, error) {

	cmd := exec.Command("git", "clone", "--single-branch", "-b", branch, "https://"+config.GitApiKey+"@"+gihubRepository, finalPath)
	_, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return "some string", nil
}
