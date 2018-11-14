package appconfig

import (
	"errors"
	"fmt"
	"github.com/robfig/config"
	"strings"
)

type MergedConfig struct {
	config  *config.Config
	section string // Check this section first, then fall back to DEFAULT
}

func NewEmptyConfig() *MergedConfig {
	return &MergedConfig{config.NewDefault(), ""}
}

func LoadConfig(confName string) (*MergedConfig, error) {
	fmt.Println(confName)
	var err error
	conf, err := config.ReadDefault(confName)
	if err == nil {
		return &MergedConfig{conf, ""}, nil
	}
	if err != nil {
		err = errors.New("not found")
	}
	return nil, err
}

func (c *MergedConfig) Int(section string, option string) (result int, found bool) {
	result, err := c.config.Int(section, option)
	if err == nil {
		return result, true
	}
	if _, ok := err.(config.OptionError); ok {
		return 0, false
	}
	return 0, false
}

func (c *MergedConfig) IntDefault(section string, option string, dfault int) int {
	if r, found := c.Int(section, option); found {
		return r
	}
	return dfault
}

func (c *MergedConfig) Bool(section string, option string) (result, found bool) {
	result, err := c.config.Bool(section, option)
	if err == nil {
		return result, true
	}
	if _, ok := err.(config.OptionError); ok {
		return false, false
	}
	return false, false
}

func (c *MergedConfig) BoolDefault(section string, option string, dfault bool) bool {
	if r, found := c.Bool(section, option); found {
		return r
	}
	return dfault
}

func (c *MergedConfig) String(section string, option string) (result string, found bool) {
	if r, err := c.config.String(section, option); err == nil {
		return stripQuotes(r), true
	}
	return "", false
}

func (c *MergedConfig) StringDefault(section string, option, dfault string) string {
	if r, found := c.String(section, option); found {
		return r
	}
	return dfault
}

func (c *MergedConfig) HasSection(section string) bool {
	return c.config.HasSection(section)
}

// Options returns all configuration option keys.
// If a prefix is provided, then that is applied as a filter.
func (c *MergedConfig) Options(prefix string) []string {
	var options []string
	keys, _ := c.config.Options(c.section)
	for _, key := range keys {
		if strings.HasPrefix(key, prefix) {
			options = append(options, key)
		}
	}
	return options
}

// Helpers

func stripQuotes(s string) string {
	if s == "" {
		return s
	}

	if s[0] == '"' && s[len(s)-1] == '"' {
		return s[1 : len(s)-1]
	}

	return s
}
