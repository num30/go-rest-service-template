package server

type FooGetRequest struct {
	Name string `path:"name" validate:"required"`
}

type FooGetResponse struct {
	Result string
}

type FooPutRequest struct {
	Name  string `path:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type FooPutResponse struct {
}
