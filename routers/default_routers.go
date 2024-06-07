package routers

import "net/http"

func AddDefaultRouters(r *http.ServeMux) {
	AddContabilRouters(r)
}
