package main

import (
	"fmt"
	"strings"
	//"time"

	"github.com/codeskyblue/go-sh"
)

const (
	lsofKeyWord1 = "python"
	lsofKeyWord2 = "root"
)

func main() {
	s := sh.NewSession()
	o, e := s.Command("lsof", "-i", `tcp:8080`).Output()
	if e != nil {
		fmt.Println("lsof error", e)
		return
	}

	resultOflsof := string(o)
	fmt.Println("lsof result:")
	fmt.Println(resultOflsof)

	a := strings.Split(resultOflsof, "\n")
	var final string
	for _, line := range a {
		if strings.Contains(line, lsofKeyWord1) && strings.Contains(line, lsofKeyWord2) {
			final = line
			break
		}
	}

	o, e = s.Command("echo", final).Command("awk", []string{"{print $2}"}).Output()
	if e != nil {
		fmt.Println("awk error", e)
		return
	}

	pid := string(o)
	pid = strings.TrimSpace(pid)
	fmt.Printf("find pid = %s\n", pid)

	e = s.Command("kill", pid).Run()
	if e != nil {
		fmt.Println("kill error", e)
		return
	}


//supervisor will restart the go-talks service automatically.

/*
	time.Sleep(time.Second * 3)

	e = s.Command("supervisorctl", "start", "go-talks").Run()
	if e != nil {
		fmt.Println("supervisorctl error", e)
		return
	}
*/
}
