# Genarate Mapping

> 根据指定结构的映射关系json文件来生成mapping映射关系go文件。



## 定义映射关系文件

##### 文件内容：

1、mapping名称（其中包括from_to_value 和 to_from_value）

2、type类型映射关系

3、content映射表内容,别名（FromName、ToName ）

4、Form映射值和To映射值（from_value、to_value）

5、change是否交换生成映射关系

##### 示例：

```json
[
	{
    "from_to_name": "OsTypeMapping",
    "to_from_name": "ToOsTypeMapping",
		"type":{
			"int":"string",
		},
		"content":{
				"[fromName]":"[toName]",
				...
		},
    "form_value":{
      "fromName":value,
      ...
    },
    "to_value":{
      "toName":value,
    },  
		"change":true
	},
	...
]
```

##### struct：

```go
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
```



##### 校验：

1、如果from_value或to_value长度不为0，长度要等于content长度

2、如果from_value或to_value长度为0，说明已经定义过值

3、type中map的key的值为from_value的类型，value为to_value的值的类型





## 使用

在根目录编译运行，或者执行`run.sh`脚本.

参数说明

```
  -i string
        input json file  path
  -o string
        output go file path
  -p string
        package name
```





## 生成映射文件示例：

##### 定义文件：

```json
[
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
]
```

##### 生成文件：

```go
package main

const (
	VendorOsTypeAndroid = 2
	VendorOsTypeIOS     = 1
)
const (
	FancyOsTypeAndroid = "android"
	FancyOsTypeIOS     = "ios"
)

var (
	OsTypeMapping = map[int]string{
		VendorOsTypeAndroid: FancyOsTypeAndroid,
		VendorOsTypeIOS:     FancyOsTypeIOS,
	}
)

var (
	ToOsTypeMapping = map[string]int{
		FancyOsTypeAndroid: VendorOsTypeAndroid,
		FancyOsTypeIOS:     VendorOsTypeIOS,
	}
)

const (
	VendorSlotTypeImage = "image"
	VendorSlotTypeVideo = "video"
)
const (
	FancySlotTypeImage = "image/mp4"
	FancySlotTypeVideo = "video/mp4"
)

var (
	SlotTypeMapping = map[string]string{
		VendorSlotTypeImage: FancySlotTypeImage,
		VendorSlotTypeVideo: FancySlotTypeVideo,
	}
)

```





