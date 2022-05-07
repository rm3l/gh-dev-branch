package issue

import "io"

type Info struct {
	Id    string
	Title string
}

type Finder interface {
	FindById(w io.Writer, repo string, issue string) (Info, error)
}
