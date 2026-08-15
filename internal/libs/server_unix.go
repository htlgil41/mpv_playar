package libs

import "net"

type ConnectionUnix struct {
	Connect net.Conn
}

func ServerSocketForUnix(path string) (*ConnectionUnix, error) {
	con, err_con := net.Dial("unix", path)
	if err_con != nil {
		return nil, err_con
	}
	return &ConnectionUnix{Connect: con}, nil
}
