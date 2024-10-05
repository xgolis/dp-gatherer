package git

import (
	"github.com/go-git/go-git/v5"
	gitHttp "github.com/go-git/go-git/v5/plumbing/transport/http"
	"github.com/spf13/pflag"
)

type GitArgs struct {
	GitRepoPath string `json:"gitRepo,omitempty"`
	GitToken    string `json:"gitToken,omitempty"`
	Username    string `json:"username,omitempty"`
	AppName     string `json:"appname,omitempty"`
}

func Pull(gitStruct GitArgs) error {
	_, err := git.PlainClone("/modules", false, &git.CloneOptions{
		URL:  gitStruct.GitRepoPath,
		Auth: &gitHttp.BasicAuth{Username: gitStruct.Username, Password: gitStruct.GitToken},
	})
	if err != nil {
		return err
	}
	return nil
}

func GetGitPullArgs(flags *pflag.FlagSet) *GitArgs {
	var gitArgs = &GitArgs{}

	flags.StringVarP(
		&gitArgs.GitRepoPath,
		"url", "p",
		"",
		"URL to Git repo",
	)
	flags.StringVarP(
		&gitArgs.Username,
		"username", "u",
		"",
		"Username for the Git repository",
	)

	return gitArgs
}
