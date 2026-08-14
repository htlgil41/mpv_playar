package libs

import "net"

func ServerSocketForUnix(path string) (net.Conn, error) {
	con, err_con := net.Dial("unix", path)
	if err_con != nil {
		return nil, err_con
	}
	return con, nil
}
