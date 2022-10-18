package plugin

import (
	config "github.com/jsmzr/bootstrap-config"
	"github.com/jsmzr/bootstrap-plugin/plugin"
)

type YamlConfigPlugin struct {
}

func (p *YamlConfigPlugin) Order() int {
	return -1
}

func (p *YamlConfigPlugin) Load() error {
	config.Register("yaml", &YamlConfig{})
	return config.InitInstance("yaml", "application.yaml")
}

func init() {
	plugin.Register("config", &YamlConfigPlugin{})
}
