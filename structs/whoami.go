package structs

type Whoami struct {
	User   string
	Home   string
	Path   string
	UID    string
	GID    string
	Groups []string
}

func NewWhoami() *Whoami {
	return &Whoami{}
}
