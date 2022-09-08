package auth

type withId struct {
	Id string
}

type workspace struct {
	Id      string
	Scripts []withId
}

type AuthInfo struct {
	AuthorId   string
	Workspaces []workspace
}
