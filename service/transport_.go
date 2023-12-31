package app

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/JIAWea/erpServer/api/erp"
	_ "github.com/JIAWea/erpServer/cmd/docs"
	"github.com/gorilla/mux"
	"github.com/ml444/gkit/auth/jwt"
	"github.com/ml444/gkit/core"
	"github.com/ml444/gkit/pkg/env"
	"github.com/ml444/gkit/transport/httpx"
	log "github.com/ml444/glog"
	"github.com/ml444/gutil/netx"
	httpSwagger "github.com/swaggo/http-swagger"
	"google.golang.org/grpc/xds"
)

func MakeHTTPHandler() http.Handler {
	router := mux.NewRouter()

	router.Methods(http.MethodGet).
		PathPrefix("/swagger/").
		Handler(httpSwagger.Handler(
			httpSwagger.DeepLinking(true),
			httpSwagger.DocExpansion("none"),
			httpSwagger.DomID("swagger-ui"),
		))

	router.Methods(http.MethodGet).
		Path(fmt.Sprintf("/%s/hello", erp.ClientName)).
		HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		})
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
	router.Methods(http.MethodPost).
		Path(fmt.Sprintf("/%s/ExportExpense", erp.ClientName)).
		HandlerFunc(File.ExportExpense)
	router.Methods(http.MethodPost).
		Path(fmt.Sprintf("/%s/ExportIncome", erp.ClientName)).
		HandlerFunc(File.ExportIncome)
	router.Methods(http.MethodPost).
		Path(fmt.Sprintf("/%s/ImportPlan", erp.ClientName)).
		HandlerFunc(File.ImportPlan)
	router.Methods(http.MethodPost).
		Path(fmt.Sprintf("/%s/ExportPlan", erp.ClientName)).
		HandlerFunc(File.ExportPlan)

	// 第三方
	router.Methods(http.MethodPost).
		Path("/third/Import").
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			body, err := io.ReadAll(request.Body)
			if err != nil {
				log.Error("err:", err)
				return
			}
			log.Infof("====> data: %+v", string(body))
		})
	router.Methods(http.MethodGet).
		Path("/third/Import").
		HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			log.Infof("====> url: %+v", request.URL)
		})

	err := httpx.ParseService2HTTP(
		NewErpService(),
		router,
		httpx.SetTimeoutMap(nil),
		// httpx.AddBeforeHandler(CheckPerm),
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

func ParseCtx(r *http.Request) context.Context {
	userId := GetUserId(r)

	header := core.Header{
		core.ClientTypeKey: core.GetHeader4HTTP(r.Header, core.HttpHeaderClientType),
		core.RemoteIp:      netx.GetRemoteIp(r),
		core.UserIdKey:     userId,
	}

	return context.WithValue(r.Context(), core.HeadersKey, header)
}
