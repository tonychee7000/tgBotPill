package config

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"path/filepath"
)

// Config for this bot
type Config struct {
	// APIToken for bot
	APIToken string `json:"apiToken"`
	// Debug mode?
	Debug bool `json:"debug"`
	// UpdateTimeout for bot
	UpdateTimeout int `json:"updateTimeout"`
}

// DefaultConfig is global.
var DefaultConfig Config

// ReadConfig from file
func (c *Config) ReadConfig(cfgFile string) error {
	absCfgFile, err := filepath.Abs(cfgFile)
	if err != nil {
		return err
	}

	f, err := os.Open(absCfgFile)
	if err != nil {
		return err
	}
	defer f.Close()

	j := json.NewDecoder(f)

	if err := j.Decode(c); err != nil {
		return err
	}
	return nil
}

func init() {
	appPath, _ := filepath.Abs(path.Dir(os.Args[0]))
	if err := DefaultConfig.ReadConfig(path.Join(appPath, "conf.json")); err != nil {
		log.Panicln(err)
	}
}
