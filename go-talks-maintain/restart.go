package main

import (
	"log"
	"strings"
	"time"

	"github.com/codeskyblue/go-sh"
)

const (
	lsofKeyWord1 = "python"
	lsofKeyWord2 = "root"
)

func resloveServicePID() string {
	s := sh.NewSession()
	o, e := s.Command("lsof", "-i", `tcp:8080`).Output()
	if e != nil {
		log.Println("lsof error", e)
		return ""
	}

	resultOflsof := string(o)
	log.Println("lsof result:")
	log.Println(resultOflsof)

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
		log.Println("awk error", e)
		return ""
	}

	pid := string(o)
	pid = strings.TrimSpace(pid)
	log.Printf("find pid = %s\n", pid)
	return pid
}

func main() {
	s := sh.NewSession()

	pid := resloveServicePID()
	if pid != "" {
		e := s.Command("kill", pid).Run()
		if e != nil {
			log.Println("kill error", e)
			return
		}
	}

	time.Sleep(time.Second * 5)

	//Although supervisor will restart the go-talks service automatically, but
	//when memory not enough, supervisor will stop restarting when restart go-talks three times.
	//so each restart stage, we are trying to re-detect whether service being restarted
	//if not, restart the service. we don't care failing.
	pid = resloveServicePID()
	if pid == "" {
		e := s.Command("supervisorctl", "restart", "go-talks").Run()
		if e != nil {
			log.Println("supervisorctl error", e)
			return
		}
		log.Println("restart go-talks ok")
	}
}
