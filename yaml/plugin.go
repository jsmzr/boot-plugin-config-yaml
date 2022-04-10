package yaml

import (
	_ "github.com/jsmzr/bootstrap-config-yaml/yaml"
	"github.com/jsmzr/bootstrap-config/config"
	"github.com/jsmzr/bootstrap-plugin/plugin"
)

type YamlConfigPlugin struct {
}

func (p *YamlConfigPlugin) Order() int {
	return -1
}

func (p *YamlConfigPlugin) Load() error {
	return config.InitInstance("yaml", "application.yaml")
}

func init() {
	plugin.Register("config", &YamlConfigPlugin{})
}
