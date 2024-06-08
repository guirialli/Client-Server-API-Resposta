package routers

import (
	"net/http"

	"github.com/guirialli/dolar-api/controllers"
)

func AddContabilRouters(r *http.ServeMux) {
	r.HandleFunc("GET /cotacao", controllers.ContabilGetBid)
	r.HandleFunc("GET /cotacao/all", controllers.ContabilGetAll)
}
