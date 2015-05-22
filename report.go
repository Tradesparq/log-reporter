package main

import (
	"bytes"
	"reflect"
	"text/template"
)

const REPORT_TEMPLATE = `
LineNumber ({{.LineNumber.Cmd}}):
{{.LineNumber.Out}}

ErrorNumber ({{.ErrorNumber.Cmd}}):
{{.ErrorNumber.Out}}

Tail ({{.Tail.Cmd}}):
{{.Tail.Out}}

ErrorSample ({{.ErrorSample.Cmd}}):
{{.ErrorSample.Out}}
`

var commands = map[string]string{
	"LineNumber":  "wc -l {{.Filename}} | cut -d ' ' -f 1",
	"Tail":        "tail -n {{.Ntail}} {{.Filename}}",
	"ErrorNumber": "grep -i error {{.Filename}} | wc -l",
	"ErrorSample": "grep -i error {{.Filename}} | head -n {{.Nerror}}",
}

type Action struct {
	Cmd string
	Out string
}

type Report struct {
	LineNumber  Action
	Tail        Action
	ErrorNumber Action
	ErrorSample Action
}

func NewReport(l Log) (Report, error) {
	var r Report
	for key, cmd := range commands {
		c, out, err := l.Exec(cmd)
		if err != nil {
			return r, err
		}

		a := Action{c, out}

		v := reflect.ValueOf(&r)
		f := v.Elem().FieldByName(key)
		f.Set(reflect.ValueOf(a))
	}

	return r, nil
}

func (r Report) String() string {
	t := template.Must(template.New("report").Parse(REPORT_TEMPLATE))
	var b bytes.Buffer
	t.Execute(&b, r)
	return b.String()
}
