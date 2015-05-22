package main

import (
	"bytes"
	"reflect"
	"text/template"
)

const REPORT_TEMPLATE = `
LineNumber: {{.LineNumber}}
ErrorNumber: {{.ErrorNumber}}

Tail:
{{.Tail}}

ErrorSample:
{{.ErrorSample}}
`

var commands = map[string]string{
	"LineNumber":  "wc -l {{.Filename}} | cut -d ' ' -f 1",
	"Tail":        "tail -n {{.Ntail}} {{.Filename}}",
	"ErrorNumber": "grep -i error {{.Filename}} | wc -l",
	"ErrorSample": "grep -i error {{.Filename}} | head -n {{.Nerror}}",
}

type Report struct {
	LineNumber  string
	Tail        string
	ErrorNumber string
	ErrorSample string
}

func NewReport(l Log) (Report, error) {
	var r Report
	for key, cmd := range commands {
		out, err := l.Exec(cmd)
		if err != nil {
			return r, err
		}

		v := reflect.ValueOf(&r)
		f := v.Elem().FieldByName(key)
		f.SetString(out)
	}

	return r, nil
}

func (r Report) String() string {
	t := template.Must(template.New("report").Parse(REPORT_TEMPLATE))
	var b bytes.Buffer
	t.Execute(&b, r)
	return b.String()
}
