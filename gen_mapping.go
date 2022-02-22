package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
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

	// 校验
	for i, mapping := range mappingList {
		if len(mapping.Content) == 0 {
			return errors.New(fmt.Sprintf("mapping index %d content len is 0", i))
		}
		if len(mapping.ToValue) != 0 {
			if len(mapping.ToValue) != len(mapping.Content) {
				return errors.New(fmt.Sprintf("mapping index %d len to_value not equal to len content", i))
			}
		}
		if len(mapping.FromValue) != 0 {
			if len(mapping.FromValue) != len(mapping.Content) {
				return errors.New(fmt.Sprintf("mapping index %d len from_value not equal to len content", i))
			}
		}
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
