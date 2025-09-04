package main

import (
	"html/template"
	"os"
)

type TemplateData struct {
	HorseNames []string
	RaceNumber int64
	VenueName  string
}

func First(slice []string) string {
	if len(slice) == 0 {
		var zero string
		return zero
	}
	return slice[0]
}

func main() {
	tmpl, err := template.New("message").Parse(
		`{{- if eq (len .HorseNames) 1 -}}
				{{ index .HorseNames 0 }}が出走した{{ .VenueName }}競馬{{ .RaceNumber }}Rの結果が確定しました
			  {{- else if gt (len .HorseNames) 1 -}}
				{{ index .HorseNames 0 }}など{{ len .HorseNames }}頭が出走した{{ .VenueName }}競馬{{ .RaceNumber }}Rの結果が確定しました
			  {{- end -}}`,
	)
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, TemplateData{
		HorseNames: []string{
			"イクイノックス",
			// "ドウデュース",
		},
		RaceNumber: 10,
		VenueName:  "東京",
	})
	if err != nil {
		panic(err)
	}
}
