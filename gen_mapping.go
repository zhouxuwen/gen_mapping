package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"os"
	"text/template"
)

type Mapping struct {
	FromToName string                 `json:"from_to_Name"`
	ToFromName string                 `json:"to_from_name"`
	Type       map[string]string      `json:"type"`
	Content    map[string]string      `json:"content"`
	FromValue  map[string]interface{} `json:"from_value"`
	ToValue    map[string]interface{} `json:"to_value"`
	Change     bool                   `json:"change"`
}
type MappingList []Mapping

func GenMappingFile() error {
	file, err := os.Open(InputFilePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	mappingList := MappingList{}
	err = json.Unmarshal(content, &mappingList)
	if err != nil {
		return err
	}

	tmpl, err := template.ParseFiles("tmpl/go_mapping_file.tmpl")
	if err != nil {
		return err
	}

	outputString := "package " + PackageName + `
`
	for _, m := range mappingList {
		for i, v := range m.FromValue {
			switch m.FromValue[i].(type) {
			case string:
				m.FromValue[i] = `"` + v.(string) + `"`
			}
		}
		for i, v := range m.ToValue {
			switch m.ToValue[i].(type) {
			case string:
				m.ToValue[i] = `"` + v.(string) + `"`
			}
		}
		buf := bytes.Buffer{}
		err = tmpl.Execute(&buf, m)
		if err != nil {
			return err
		}
		outputString += buf.String()
	}

	err = ioutil.WriteFile(OutputFilePath, []byte(outputString), 0644)
	if err != nil {
		return err
	}

	return nil
}
