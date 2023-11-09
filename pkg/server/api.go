package server

type ThingGetRequest struct {
	Name string `path:"name" validate:"required"`
}

type ThingGetResponse struct {
	Result string
}

type ThingPutRequest struct {
	Name  string `path:"name" validate:"required"`
	Value string `json:"value" validate:"required"`
}

type ThingPutResponse struct {
}
