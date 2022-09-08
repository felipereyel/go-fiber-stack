package auth

type withId struct {
	Id string
}

type workspace struct {
	Id      string
	Forms   []withId
	Hooks   []withId
	Jobs    []withId
	Scripts []withId
}

type AuthInfo struct {
	AuthorId   string
	Workspaces []workspace
}
