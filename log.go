package main

import (
	"bytes"
	"os/exec"
	"text/template"
	//"strconv"
)

type Log struct {
	Filename string
	Ntail    uint
	Nerror   uint
}

func (l *Log) Exec(command string) (string, error) {
	var b bytes.Buffer
	t := template.Must(template.New("cmd").Parse(command))

	err := t.Execute(&b, l)
	if err != nil {
		return "", err
	}

	cmd := exec.Command("sh", "-c", b.String())
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()

	return out.String(), err
}
