package issue

import "io"

// Info is a generic struct representing a given issue
type Info struct {
	Id    string
	Title string
}

// Finder is the interface used to manipulate issues
type Finder interface {

	// FindById allows to retrieve an issue given its identifier.
	// It optionally writes any relevant warnings or messages to eventually output into the io.Writer object specified.
	FindById(w io.Writer, repo string, id string) (Info, error)
}
