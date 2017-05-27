// Package config owns common functions to get the application's configuration.
package config

import (
	"os"
	"strings"
)

// Auth0Config represents the Auth0 configuration.
type Auth0Config struct {
	Issuer   string
	Secret   string
	Audience []string
}

// Port returns the PORT value from environment or it defaults to `3000`.
func Port() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	return port
}

// DSN returns the HOST value from environment.
func DSN() string {
	return os.Getenv("CRICK_DSN")
}

// Auth0 returns the Auth0 configuration.
func Auth0() Auth0Config {
	return Auth0Config{
		Issuer:   os.Getenv("AUTH0_ISSUER"),
		Secret:   os.Getenv("AUTH0_SECRET"),
		Audience: strings.Split(os.Getenv("AUTH0_AUDIENCE"), ","),
	}
}
