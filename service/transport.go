package app

import (
	"net/http"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/gorilla/mux"
	"github.com/ml444/gkit/transport/httpx"
	log "github.com/ml444/glog"
	"google.golang.org/grpc/xds"
)



func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()

	err := httpx.ParseService2HTTP(
        NewErpService(), 
        router, 
        httpx.SetTimeoutMap(nil),
		httpx.SetJwtHook(nil),
    )
    if err != nil {
		log.Errorf("err: %v", err)
	}

	return router
}

func RegisterServerToGRPC(s *xds.GRPCServer) {

	erp.RegisterErpServer(s, NewErpService())

}
