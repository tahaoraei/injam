package config

import (
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/v2"
	"log/slog"
	"strings"
)

func Load(path string) Config {
	var k = koanf.New(".")
	k.Load(env.Provider("INJAM_", ".", func(s string) string {
		return strings.Replace(strings.ToLower(strings.TrimPrefix(s, "INJAM_")), "_", ".", -1)
	}), nil)
	k.Load(file.Provider(path), yaml.Parser())

	var cfg Config
	if err := k.Unmarshal("", &cfg); err != nil {
		slog.Error("reading config: ", err)
	}
	return cfg
}
