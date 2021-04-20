package services

import (
	"fmt"
	"net"
	"time"

	"github.com/urfave/cli"
)

func CheckServer(c *cli.Context) {
	var host string = c.String("host")
	var port string = c.String("port")

	timeout := time.Duration(time.Second)

	_, erro := net.DialTimeout("tcp", host+":"+port, timeout)

	if erro != nil {
		fmt.Printf("Server %s does not connected on port %s\n", host, port)
	} else {
		fmt.Printf("Server %s connected on port %s\n", host, port)
	}
}
