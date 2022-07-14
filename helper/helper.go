package helper

import (
	"bufio"
	"bytes"
	"go/format"
	"io/ioutil"
	"os"
	"text/template"
)

func ProcessTemplate(templatePath string, outPutPath string, data map[string]interface{}) error {
	// парсим наш шаблон
	fbytes, _ := ioutil.ReadFile(templatePath)
	tmpl := template.Must(template.New("Text").Parse(string(fbytes)))
	// создаём пустой буфер для записи
	buf := new(bytes.Buffer)
	// рендерим шаблон с данными data в наш буфер buf
	err := tmpl.Execute(buf, data)
	if err != nil {
		return err
	}
	// необходимо отформатировать данные
	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}
	f, _ := os.Create(outPutPath)
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(formatted))
	if err != nil {
		return err
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}

func ProcessSQLTemplate(templatePath string, outPutPath string, data map[string]interface{}) error {
	// парсим наш шаблон
	tmpl := template.Must(template.ParseFiles(templatePath))
	// создаём пустой буфер для записи
	buf := new(bytes.Buffer)
	// рендерим шаблон с данными data в наш буфер buf
	err := tmpl.Execute(buf, data)
	if err != nil {
		return err
	}
	f, _ := os.Create(outPutPath)
	w := bufio.NewWriter(f)
	_, err = w.WriteString(string(buf.Bytes()))
	if err != nil {
		return err
	}
	err = w.Flush()
	if err != nil {
		return err
	}
	return nil
}
