package handler

type Router interface {
	Path() (path string)
	Method() (method string)
}
