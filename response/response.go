package response

import "net/http"

type RenderableResponse interface {
	Render(w http.ResponseWriter)
}

type Response struct {
	Content string
	Status  int
	Headers map[string]string
}

func (r *Response) Render(w http.ResponseWriter) {
	renderResponse(r, w)
}

type JsonResponse struct {
	Response
	Content     interface{}
	PrettyPrint bool
}

func (r *JsonResponse) Render(w http.ResponseWriter) {
	renderJsonResponse(r, w)
}
