package issue

import (
	"encoding/json"
	"fmt"
	"github.com/cli/go-gh"
	"io"
	"strconv"
)

type Info struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}

func Lookup(w io.Writer, repo string, issue int) (Info, error) {
	args := []string{"issue", "view", strconv.Itoa(issue), "--json", "number,title"}
	if repo != "" {
		args = append(args, "-R", repo)
	}
	var issueInfo Info
	stdOut, stdErr, err := gh.Exec(args...)
	if err != nil {
		return issueInfo, err
	}
	//goland:noinspection GoUnhandledErrorResult
	fmt.Fprintln(w, stdErr.String())
	err = json.Unmarshal(stdOut.Bytes(), &issueInfo)
	return issueInfo, err
}
