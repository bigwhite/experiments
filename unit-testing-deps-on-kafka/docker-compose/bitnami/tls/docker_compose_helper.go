package main

import (
	"fmt"
	"os/exec"
	"strings"
	"time"
)

// helpler function for operating docker container through docker-compose command

const (
	defaultCmd     = "docker-compose"
	defaultCfgFile = "docker-compose.yml"
)

func execCliCommand(cmd string, opts ...string) ([]byte, error) {
	cmds := cmd + " " + strings.Join(opts, " ")
	fmt.Println("exec command:", cmds)
	return exec.Command(cmd, opts...).CombinedOutput()
}

func execDockerComposeCommand(cmd string, cfgFile string, opts ...string) ([]byte, error) {
	var allOpts = []string{"-f", cfgFile}
	allOpts = append(allOpts, opts...)
	return execCliCommand(cmd, allOpts...)
}

func UpKakfa(composeCfgFile string) ([]byte, error) {
	b, err := execDockerComposeCommand(defaultCmd, composeCfgFile, "up", "-d")
	if err != nil {
		return nil, err
	}
	time.Sleep(10 * time.Second)
	return b, nil
}

func UpDefaultKakfa() ([]byte, error) {
	return UpKakfa(defaultCfgFile)
}

func DownKakfa(composeCfgFile string) ([]byte, error) {
	b, err := execDockerComposeCommand(defaultCmd, composeCfgFile, "down", "-v")
	if err != nil {
		return nil, err
	}
	time.Sleep(10 * time.Second)
	return b, nil
}

func DownDefaultKakfa() ([]byte, error) {
	return DownKakfa(defaultCfgFile)
}
