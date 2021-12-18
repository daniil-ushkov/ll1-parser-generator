package generator

import (
	"bytes"
	"os"
	"os/exec"
	"text/template"
)

func execute(text string, data interface{}) (string, error) {
	tmpl, err := template.New("RuleParser").Parse(text)
	if err != nil {
		return "", err
	}

	buf := bytes.Buffer{}

	err = tmpl.Execute(&buf, data)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}

func closeFmt(file *os.File) {
	_ = file.Close()
	_ = exec.Command("gofmt", "-w", file.Name()).Run()
	_ = exec.Command("go", "generate", file.Name()).Run()
}
