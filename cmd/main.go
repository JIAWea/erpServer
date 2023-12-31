package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/http/pprof"

	"github.com/JIAWea/erpServer/api/erp"
	"github.com/JIAWea/erpServer/config"
	"github.com/JIAWea/erpServer/internal"
	"github.com/JIAWea/erpServer/internal/middleware"
	"github.com/JIAWea/erpServer/service"
	"github.com/ml444/gkit/errorx"
	log "github.com/ml444/glog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	xdscreds "google.golang.org/grpc/credentials/xds"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/xds"
)

// Run starts a new http server, gRPC server, and a debug server with the
// passed config and logger
func Run(cfg *config.Config) {
	var err error

	// Mechanical domain.
	errCh := make(chan error)

	// Interrupt handler.
	go internal.InterruptHandler(errCh)

	// Debug pprof listener.
	if cfg.Debug {
		go func() {
			log.Info("transport pprof addr", cfg.PprofAddr)

			m := http.NewServeMux()
			m.Handle("/debug/pprof/", http.HandlerFunc(pprof.Index))
			m.Handle("/debug/pprof/cmdline", http.HandlerFunc(pprof.Cmdline))
			m.Handle("/debug/pprof/profile", http.HandlerFunc(pprof.Profile))
			m.Handle("/debug/pprof/symbol", http.HandlerFunc(pprof.Symbol))
			m.Handle("/debug/pprof/trace", http.HandlerFunc(pprof.Trace))

			errCh <- http.ListenAndServe(cfg.PprofAddr, m)
		}()
	}

	// HTTP transport.
	if cfg.EnableHTTP {
		go func() {
			log.Info("transport HTTP addr", cfg.HTTPAddr)
			h := app.MakeHTTPHandler()
			errCh <- http.ListenAndServe(cfg.HTTPAddr, h)
		}()
	}

	// gRPC transport.
	if cfg.EnableGRPC {
		go func() {
			log.Info("transport gRPC addr", cfg.GRPCAddr)
			ln, err := net.Listen("tcp", cfg.GRPCAddr)
			if err != nil {
				errCh <- err
				return
			}
			creds, err := xdscreds.NewServerCredentials(xdscreds.ServerOptions{FallbackCreds: insecure.NewCredentials()})
			if err != nil {
				println(err.Error())
				errCh <- err
				return
			}

			s := xds.NewGRPCServer(
				grpc.Creds(creds),
				grpc.ChainUnaryInterceptor(middleware.InterceptorList...),
			)
			app.RegisterServerToGRPC(s)

			healthPort := fmt.Sprintf(":%d", 5041)
			healthLis, err := net.Listen("tcp", healthPort)
			if err != nil {
				log.Errorf("net.Listen(tcp4, %q) failed: %v", healthPort, err)
				errCh <- err
				return
			}
			grpcServer := grpc.NewServer()
			healthServer := health.NewServer()
			healthServer.SetServingStatus("", healthpb.HealthCheckResponse_SERVING)
			healthpb.RegisterHealthServer(grpcServer, healthServer)

			log.Infof("Serving GreeterService on %s and HealthService on %s", ln.Addr().String(), healthLis.Addr().String())
			go func() {
				err = grpcServer.Serve(healthLis)
				if err != nil {
					println(err.Error())
				}
			}()
			// errCh <- s.Serve(ln)
			err = s.Serve(ln)
			if err != nil {
				// TODO: xds: failed to create xds-client: xds: bootstrap env vars are unspecified and provided fallback config is nil
				println(err.Error())
			}
		}()
	}

	// Run!
	err = <-errCh
	println("exit", err.Error())
}

//go:generate swag init --parseDependency --parseInternal
// @title Swagger Example API
// @version 1.0
// @description ErpServer API

// @host      localhost:5050
// @BasePath  /erp
func main() {
	flag.Parse()
	config.Init()
	cfg := config.GetConfig()
	errorx.RegisterError(erp.ErrCodeMap, erp.ErrCode4StatusCodeMap)

	// init
	err := Init(cfg)
	if err != nil {
		println(err)
		return
	}
	// defer something. example: rpc connection\ db connection\ file...
	defer Defer()

	app.Init()

	Run(cfg)
}
