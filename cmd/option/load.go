package option

import (
	"fmt"

	"github.com/spf13/viper"
)

// LoadCmdConfig is config for sum command
type LoadCmdConfig struct {
	Dir   string
	Depth uint
}

// NewLoadCmdConfigFromViper generate config for sum command from viper
func NewLoadCmdConfigFromViper(args []string) (*LoadCmdConfig, error) {
	var conf LoadCmdConfig
	if err := viper.Unmarshal(&conf); err != nil {
		return nil, fmt.Errorf("failed to unmarshal config from viper: %w", err)
	}

	if err := conf.validate(); err != nil {
		return nil, fmt.Errorf("failed to create sum cmd config: %w", err)
	}

	return &conf, nil
}

func (c *LoadCmdConfig) validate() error {
	return nil
}
