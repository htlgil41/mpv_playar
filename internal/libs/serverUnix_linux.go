package libs

import (
	"fmt"
	"net"
)

type ConnectionUnix struct {
	Connect net.Conn
}

func ServerSocketForUnix(path string) (*ConnectionUnix, error) {
	con, err_con := net.Dial("unix", path)
	if err_con != nil {
		return nil, err_con
	}

	_, err_writestatus := con.Write([]byte(`{ "command": ["get_property", "mpv-version"] }` + "\n"))
	if err_writestatus != nil {
		fmt.Println("Error al escribir el comando de estado en mpv")
	}
	return &ConnectionUnix{Connect: con}, nil
}
