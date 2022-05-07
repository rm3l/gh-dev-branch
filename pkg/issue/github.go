package issue

import (
	"encoding/json"
	"fmt"
	"github.com/cli/go-gh"
	"io"
	"strconv"
)

type GHIssueFinder struct{}

type _GhIssueInfo struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

func NewGHIssueFinder() *GHIssueFinder {
	return &GHIssueFinder{}
}

func (g GHIssueFinder) FindById(w io.Writer, repo string, id string) (Info, error) {
	args := []string{"issue", "view", id, "--json", "number,title"}
	if repo != "" {
		args = append(args, "-R", repo)
	}
	var issueInfo _GhIssueInfo
	stdOut, stdErr, err := gh.Exec(args...)
	if err != nil {
		return Info{}, err
	}
	if stdErrStr := stdErr.String(); stdErrStr != "" {
		//goland:noinspection GoUnhandledErrorResult
		fmt.Fprintln(w, stdErrStr)
	}
	err = json.Unmarshal(stdOut.Bytes(), &issueInfo)
	return Info{
		Id:    strconv.Itoa(issueInfo.Number),
		Title: issueInfo.Title,
	}, err
}
