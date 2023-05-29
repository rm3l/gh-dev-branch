package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/rm3l/gh-dev-branch/pkg/branch"
	"github.com/rm3l/gh-dev-branch/pkg/issue"
)

func main() {

	var repo string

	flag.StringVar(&repo, "repo", "", "repository to use for finding the issue")

	flag.Usage = func() {
		//goland:noinspection GoUnhandledErrorResult
		fmt.Fprintln(os.Stderr, "Usage: gh dev-branch <issue> [options]")
		fmt.Println("Options: ")
		flag.PrintDefaults()
	}
	if len(os.Args) < 2 {
		//goland:noinspection GoUnhandledErrorResult
		fmt.Fprintln(os.Stderr, "missing issue")
		flag.Usage()
		os.Exit(1)
	}

	iss := os.Args[1]
	if iss == "-h" || iss == "-help" || iss == "--help" {
		flag.Usage()
		os.Exit(1)
	} else {
		// Ignore errors since flag.CommandLine is set for ExitOnError.
		_ = flag.CommandLine.Parse(os.Args[2:])
	}

	ghIssueFinder := issue.NewGHIssueFinder()
	issueInfo, err := ghIssueFinder.FindById(os.Stderr, repo, iss)
	if err != nil {
		log.Fatalln(err)
	}

	branchName, err := branch.GenerateName(issueInfo.Id, issueInfo.Title)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(branchName)
}
