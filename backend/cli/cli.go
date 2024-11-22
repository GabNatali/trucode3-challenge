package cli

import (
	"github.com/GabNatali/trucode3-challenge-final/config"
	"github.com/alecthomas/kong"
)

// Scheme

type scheme struct {
	EnvPath string `help:"Path to env config file" type:"path" optional:""`
}

// Parser

type Parser struct {
	scheme scheme
}

func NewParser() *Parser {
	return &Parser{scheme: scheme{}}
}

func (p *Parser) ParseConfig() (*config.Config, error) {
	kong.Parse(&p.scheme)
	return config.ParseEnv(p.scheme.EnvPath)
}
