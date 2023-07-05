package app

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/ml444/gkit/transport/httpx"
	log "github.com/ml444/glog"
	"google.golang.org/grpc/xds"
	"github.com/JIAWea/erpServer/api/erp"
)



func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()

	err := httpx.ParseService2HTTP(
        NewErpService(),
        router, 
        httpx.SetTimeoutMap(nil), 
    )
    if err != nil {
		log.Errorf("err: %v", err)
	}

	return router
}

func RegisterServerToGRPC(s *xds.GRPCServer) {

	erp.RegisterErpServer(s, NewErpService())

}
