package response

import (
	"encoding/json"
	"net/http"
)

func renderResponse(r *Response, w http.ResponseWriter) {
	writeStatus(r.Status, w)
	writeHeaders(r.Headers, w)
	writeContent([]byte(r.Content), w)
}

func renderJsonResponse(r *JsonResponse, w http.ResponseWriter) {
	writeStatus(r.Status, w)
	writeHeaders(r.Headers, w)
	writeContent(renderJson(r), w)
}

func writeStatus(status int, w http.ResponseWriter) {
	if status != 0 {
		w.WriteHeader(http.StatusOK)
	}
}

func writeHeaders(headers map[string]string, w http.ResponseWriter) {
	for k, v := range headers {
		w.Header().Add(k, v)
	}
}

func writeContent(content []byte, w http.ResponseWriter) {
	w.Write(content)
}

func renderJson(r *JsonResponse) []byte {
	var data interface{} = r.Content

	if e, ok := r.Content.(error); ok {
		data = map[string]string{"error": e.Error()}
	}

	if r.PrettyPrint {
		js, _ := json.MarshalIndent(data, " ", "\t")
		return js
	}

	js, _ := json.Marshal(data)
	return js
}
