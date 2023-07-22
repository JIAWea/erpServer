package app

import (
	"fmt"
	"github.com/JIAWea/erpServer/api/erp"
	"github.com/gorilla/mux"
	"github.com/ml444/gkit/transport/httpx"
	log "github.com/ml444/glog"
	"google.golang.org/grpc/xds"
	"net/http"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()

	router.Methods(http.MethodPost).
		Path(fmt.Sprintf("/%s/upload", erp.ClientName)).
		HandlerFunc(File.UploadFile)
	router.Methods(http.MethodGet).
		Path(fmt.Sprintf("/%s/download", erp.ClientName)).
		HandlerFunc(File.DownloadFile)

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
