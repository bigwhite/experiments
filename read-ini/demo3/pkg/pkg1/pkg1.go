package pkg1

import (
	"fmt"

	"github.com/bigwhite/readini/pkg/config"
)

func Foo() {
	id, ok := config.GetSectionKey("server.id")
	fmt.Printf("id = [%v], ok = [%t]\n", id, ok)
	tlsPort, ok := config.GetSectionKey("server.tls_port")
	fmt.Printf("tls_port = [%v], ok = [%t]\n", tlsPort, ok)
	logPath, ok := config.GetSectionKey("log.path")
	fmt.Printf("path = [%v], ok = [%t]\n", logPath, ok)
	logPath1, ok := config.GetSectionKey("log.path1")
	fmt.Printf("path1 = [%v], ok = [%t]\n", logPath1, ok)

	b, ok := config.GetBool("debug.profile_on")
	fmt.Printf("profile_on = [%t], ok = [%t]\n", b, ok)
}
