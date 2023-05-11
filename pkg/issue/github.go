package issue

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/cli/go-gh/v2"
)

// GHIssueFinder is a struct type that implements the issue.Finder interface
type GHIssueFinder struct{}

// This provides a compile-safe way of making sure GHIssueFinder implements the issue.Finder interface
var _ Finder = GHIssueFinder{}

// _GhIssueInfo is a struct used to unmarshal GitHub API responses.
// It needs to be exported but is not intended to be used directly.
type _GhIssueInfo struct {
	// Number is the GitHub issue number
	Number int `json:"number"`
	// Title is the GitHub issue title
	Title string `json:"title"`
}

// NewGHIssueFinder returns a new object of type GHIssueFinder
func NewGHIssueFinder() *GHIssueFinder {
	return &GHIssueFinder{}
}

// FindById allows to retrieve an issue given its identifier.
// It optionally writes any relevant warnings or messages to eventually output into the io.Writer object specified.
func (g GHIssueFinder) FindById(w io.Writer, repo string, id string) (Info, error) {
	args := []string{"issue", "view", id, "--json", "number,title"}
	if repo != "" {
		args = append(args, "-R", repo)
	}
	stdOut, stdErr, err := gh.Exec(args...)
	if err != nil {
		return Info{}, err
	}
	if stdErrStr := stdErr.String(); stdErrStr != "" {
		//goland:noinspection GoUnhandledErrorResult
		fmt.Fprintln(w, stdErrStr)
	}
	var issueInfo _GhIssueInfo
	err = json.Unmarshal(stdOut.Bytes(), &issueInfo)
	return Info{
		Id:    strconv.Itoa(issueInfo.Number),
		Title: issueInfo.Title,
	}, err
}
