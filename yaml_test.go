package plugin

import (
	"testing"
)

func TestLoad(t *testing.T) {
	yamlConfig := YamlConfig{}
	if _, err := yamlConfig.Load("not_exist.yaml"); err == nil {
		t.Fatal("file not exist")
	}

}

func TestConvertDict(t *testing.T) {
	dict := make(map[interface{}]interface{})
	dict[0] = 0
	dict["str_1"] = "1"
	dict1 := make(map[string]interface{})
	dict1["num_3"] = 1
	dict["dict_1"] = dict1

	res := convertDict(dict)
	if v0, ok := res["0"]; !ok || v0 != 0 {
		t.Fatal("convert type failed")
	}
	if v1, ok := res["str_1"]; !ok || v1 != "1" {
		t.Fatal("convert type failed")
	}

}
