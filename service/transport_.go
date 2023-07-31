package app

import (
	"fmt"
	"github.com/JIAWea/erpServer/api/erp"
	"github.com/gorilla/mux"
	"github.com/ml444/gkit/auth/jwt"
	"github.com/ml444/gkit/core"
	"github.com/ml444/gkit/pkg/env"
	"github.com/ml444/gkit/transport/httpx"
	log "github.com/ml444/glog"
	"google.golang.org/grpc/xds"
	"net/http"
	"strings"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()

	router.Methods(http.MethodPost).
		Path(fmt.Sprintf("/%s/UploadFile", erp.ClientName)).
		HandlerFunc(File.UploadFile)
	router.Methods(http.MethodPost).
		Path(fmt.Sprintf("/%s/ImportExpense", erp.ClientName)).
		HandlerFunc(File.ImportExpense)
	router.Methods(http.MethodGet).
		Path(fmt.Sprintf("/%s/DownloadFile", erp.ClientName)).
		HandlerFunc(File.DownloadFile)
	router.Methods(http.MethodPost).
		Path(fmt.Sprintf("/%s/ImportIncome", erp.ClientName)).
		HandlerFunc(File.ImportIncome)

	err := httpx.ParseService2HTTP(
		NewErpService(),
		router,
		httpx.SetTimeoutMap(nil),
		httpx.AddBeforeHandler(CheckPerm),
	)
	if err != nil {
		log.Errorf("err: %v", err)
	}

	return router
}

func RegisterServerToGRPC(s *xds.GRPCServer) {

	erp.RegisterErpServer(s, NewErpService())

}

func GetUserId(r *http.Request) uint64 {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		log.Warn("not found Authorization")
		if env.IsLocalEnv() {
			return core.GetUserId4Headers(r.Header)
		}
		return 0
	}

	authHeaderParts := strings.Split(authHeader, " ")
	if len(authHeaderParts) != 2 || !strings.EqualFold(authHeaderParts[0], jwt.Bearer) {
		log.Error("token invalid")
		return 0
	}
	tokenString := authHeaderParts[1]
	partList := strings.Split(tokenString, ".")
	if len(partList) != 3 {
		log.Error("token invalid")
		return 0
	}

	var (
		claims *jwt.CustomClaims
		err    error
	)
	claims, err = jwt.ParsePayload(partList[1])
	if err != nil {
		log.Error("err:", err)
		return 0
	}

	return claims.UserId
}
