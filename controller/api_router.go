package controller

import(
	"net/http"
)

type Router interface {
	HandleApiRequest(w http.ResponseWriter, r *http.Request) 
}

type router struct {
	tc ApiContoller
}

func NewRouter(tc ApiContoller) *router {
	return &router{tc}
}

func (ro *router) HandleApiRequest(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
    case "GET":
		ro.tc.GetApi(w, r)
	case "POST":
	    ro.tc.PostApi(w, r)
	case "PUT":
		ro.tc.UpdateApi(w, r)
	case "DELETE":
		ro.tc.DeleteApi(w, r)
	default :
	    w.WriteHeader(405)
	}
}
