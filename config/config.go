// SPDX-FileCopyrightText: 2021-2022 Adam Tauber, <asciimoo@gmail.com> et al.
//
// SPDX-License-Identifier: AGPL-3.0-only

package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	App     App     `yaml:"app"`
	Server  Server  `yaml:"server"`
	DB      DB      `yaml:"db"`
	Storage Storage `yaml:"storage"`
	SMTP    SMTP    `yaml:"smtp"`
}

type App struct {
	Debug            bool  `yaml:"debug"`
	BookmarksPerPage int64 `yaml:"bookmarks_per_page"`
	DisableSignup    bool  `yaml:"disable_signup"`
}

type Server struct {
	Address string `yaml:"address"`
	BaseURL string `yaml:"base_url"`
}

type DB struct {
	Connection string `yaml:"connection"`
	Type       string `yaml:"type"`
}

type Storage struct {
	Type string `yaml:"type"`
	Root string `yaml:"root"`
}

type SMTP struct {
	Host              string `yaml:"host"`
	Port              int    `yaml:"port"`
	Username          string `yaml:"username"`
	Password          string `yaml:"password"`
	Sender            string `yaml:"sender"`
	TLS               bool   `yaml:"tls"`
	TLSAllowInsecure  bool   `yaml:"tls_allow_insecure"`
	SendTimeout       int    `yaml:"send_timeout"`
	ConnectionTimeout int    `yaml:"connection_timeout"`
}

func Load(filename string) (*Config, error) {
	var c *Config
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return c, err
	}
	err = yaml.Unmarshal(b, &c)
	// TODO validate config
	return c, err
}
