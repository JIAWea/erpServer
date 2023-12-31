package erp

import (
    "errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	xdscreds "google.golang.org/grpc/credentials/xds"
)

// NewXDSConn new a connection of xDs
// Note: call `conn.Close()` when the server exits
func NewXDSConn() (*grpc.ClientConn, error) {
	var conn *grpc.ClientConn
	var err error
	creds, err := xdscreds.NewClientCredentials(
		xdscreds.ClientOptions{FallbackCreds: insecure.NewCredentials()},
	)
	if err != nil {
		return nil, err
	}
	conn, err = grpc.Dial(
    "xds:///erp:5040",
		grpc.WithTransportCredentials(creds),
	)
	if err != nil {
		return nil, err
	}
	//NOTE: defer conn.Close()
	return conn, nil
}


type CliManager struct {
	conn    *grpc.ClientConn
    cli     ErpClient
	initErr error
}

func InjectConn(conn *grpc.ClientConn) {
	if conn == nil {
		return
	}
	cliMgr.conn = conn
    cliMgr.cli = NewErpClient(conn)
}

func CloseConn() error {
    if cliMgr.conn == nil {
        return nil
    }
    return cliMgr.conn.Close()
}

var cliMgr = CliManager{
	initErr: errors.New(
		"not yet initialized grpc client, " +
		"please call 'NewXDSConn()' and 'InjectConn()' on server, " +
		"and conn.Close() when the server exits"),
}



