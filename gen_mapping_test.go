package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
	"text/template"
)

func TestGenMapping(t *testing.T) {
	str := `[
  {
    "from_to_name": "OsTypeMapping",
	"to_from_name": "ToOsTypeMapping",
    "type": {
      "int": "string"
    },
    "content": {
      "VendorOsTypeIOS": "FancyOsTypeIOS",
      "VendorOsTypeAndroid": "FancyOsTypeAndroid"
    },
    "from_value": {
      "VendorOsTypeIOS": 1,
      "VendorOsTypeAndroid": 2
    },
    "to_value": {
      "FancyOsTypeIOS": "ios",
      "FancyOsTypeAndroid": "android"
    },
    "change": true
  },
  {
    "from_to_name": "SlotTypeMapping",
	"to_from_name": "ToSlotTypeMapping",
    "type": {
      "string": "string"
    },
    "content": {
      "VendorSlotTypeVideo":"FancySlotTypeVideo",
      "VendorSlotTypeImage": "FancySlotTypeImage"
    },
    "from_value": {
      "VendorSlotTypeVideo": "video",
      "VendorSlotTypeImage": "image"
    },
    "to_value": {
      "FancySlotTypeVideo": "video/mp4",
      "FancySlotTypeImage": "image/mp4"
    },
    "change": false
  }
]`
	mappingList := MappingList{}
	err := json.Unmarshal([]byte(str), &mappingList)
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}
	//fmt.Printf("%+v", mappingList)

	tmpl, err := template.ParseFiles("tmpl/go_mapping_file.tmpl")
	if err != nil {
		fmt.Printf("err:%v\n", err)
	}

	packageName := "main"
	outputString := "package " + packageName + `
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
			fmt.Printf("err:%v\n", err)
		}
		fmt.Printf("%v\n", buf.String())
		outputString += buf.String()
	}

	err = ioutil.WriteFile("go_mapping.go", []byte(outputString), 0644)
	if err != nil {
		panic(err)
	}

}
