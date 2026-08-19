package libs

import (
	"net"

	"github.com/Microsoft/go-winio"
)

type ConnectionUnix struct {
	Connect net.Conn
}

func ServerSocketForUnix(path string) (*ConnectionUnix, error) {
	con, err_con := winio.DialPipe(path, nil)
	if err_con != nil {
		return nil, err_con
	}
	return &ConnectionUnix{Connect: con}, nil
}
