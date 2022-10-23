package plugin

import (
	config "github.com/jsmzr/boot-config"
	plugin "github.com/jsmzr/boot-plugin"
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

func (p *YamlConfigPlugin) Enabled() bool {
	return true
}

func init() {
	plugin.Register("config", &YamlConfigPlugin{})
}
