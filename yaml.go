package plugin

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/jsmzr/bootstrap-config/config"
	"gopkg.in/yaml.v2"
)

type YamlConfig struct{}

type YamlContainer struct {
	content string
}

func (c *YamlConfig) Load(filename string) (config.Configer, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var instance map[interface{}]interface{}
	if err := yaml.Unmarshal(data, &instance); err != nil {
		return nil, err
	}
	dict := convertDict(instance)
	if contentBytes, err := json.Marshal(dict); err != nil {
		return nil, err
	} else {
		return &YamlContainer{content: string(contentBytes)}, nil
	}
}

func convertDict(data map[interface{}]interface{}) map[string]interface{} {
	dict := make(map[string]interface{})
	for k, v := range data {
		keyStr := fmt.Sprintf("%v", k)
		convertVaule, ok := v.(map[interface{}]interface{})
		if ok {
			dict[keyStr] = convertDict(convertVaule)
		} else {
			dict[keyStr] = v
		}
	}
	return dict
}

func (c *YamlContainer) GetJson() string {
	return c.content
}
