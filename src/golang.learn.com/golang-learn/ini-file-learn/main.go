package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strconv"
	"strings"
)

// ini配置文件解析器

// MysqlConfig MySql配置结构体
type MysqlConfig struct {
	Address  string `ini:"address"`
	Port     int    `ini:"port"`
	UserName string `ini:"username"`
	Password string `ini:"password"`
}

// RedisConfig Redis配置结构体
type RedisConfig struct {
	Host     string `ini:"host"`
	Port     int    `ini:"port"`
	Password string `ini:"password"`
	Database string `ini:"database"`
}

// Config config配置内容结构体
type Config struct {
	MysqlConfig `ini:"mysql"`
	RedisConfig `ini:"redis"`
}

func AnalysisIni(fileName string, distribution interface{}) (err error) {
	t := reflect.TypeOf(distribution)
	if t.Kind() != reflect.Ptr {
		err = errors.New("distribution param should be a pointer")
		return
	}
	// 使用专有的Elem()方法来获取指针对应的值
	if t.Elem().Kind() != reflect.Struct {
		err = errors.New("distribution param should be a struct pointer")
		return
	}
	// 获取文件内容
	byteContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		return
	}
	// 切割字符串
	linSlice := strings.Split(string(byteContent), "\r\n")
	// fmt.Println(linSlice)
	var structName string
	for idx, line := range linSlice {
		line = strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}
		if strings.HasPrefix(line, ";") || strings.HasPrefix(line, "#") {
			continue
		}
		if strings.HasPrefix(line, "[") {
			distributeType := strings.TrimSpace(line[1 : len(line)-1])
			if strings.HasSuffix(line, "]") && len(distributeType) > 0 {
				for i := 0; i < t.Elem().NumField(); i++ {
					filed := t.Elem().Field(i)
					if distributeType == filed.Tag.Get("ini") {
						structName = filed.Name
						fmt.Println("get field", structName)
						break
					}
				}
			} else {
				err = fmt.Errorf("line:%d syntax error", idx+1)
				return
			}
		} else {
			if strings.Index(line, "=") == -1 || strings.HasPrefix(line, "=") {
				err = fmt.Errorf("line: %d syntax error", idx+1)
				return
			}
			v := reflect.ValueOf(distribution)
			structVal := v.Elem().FieldByName(structName)
			structType := structVal.Type()
			if structVal.Kind() != reflect.Struct {
				err = fmt.Errorf("字段中的%s应该是一个结构体", structName)
				return
			}
			indexOfEqual := strings.Index(line, "=")
			key := strings.TrimSpace(line[:indexOfEqual])
			value := strings.TrimSpace(line[indexOfEqual+1:])
			var fieldName string
			var fieldType reflect.StructField
			for i := 0; i < structVal.NumField(); i++ {
				filed := structType.Field(i)
				if key == filed.Tag.Get("ini") {
					fieldType = filed
					fieldName = filed.Name
					break
				}
			}
			fieldObj := structVal.FieldByName(fieldName)
			fmt.Println(fieldName, fieldType.Type.Kind(), value)
			switch fieldType.Type.Kind() {
			case reflect.String:
				fieldObj.SetString(value)
			case reflect.Int, reflect.Int16, reflect.Int32, reflect.Int64:
				var valueInt int64
				valueInt, err = strconv.ParseInt(value, 10, 64)
				if err != nil {
					err = fmt.Errorf("line:%d value type error", idx+1)
					return
				}
				fieldObj.SetInt(valueInt)
			}
		}

	}
	return
}

func main() {
	var cfg Config
	err := paseIni("./conf.ini", &cfg)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("配置完成后的结构体:%#v\n", cfg)
}
