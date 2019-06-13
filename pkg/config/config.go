package config

import (
	"flag"
	"os"
)

// Config contains all configuration of App
type Config struct {
	DSN      string
	HTTPAddr string
	NATSAddr string
	NATSSub  string
}

// GetConfig returns a fulfilled Config
func GetConfig() Config {
	fs := flag.NewFlagSet("simpleNews", flag.ExitOnError)
	var (
		dsn      = fs.String("dsn", "root:root@tcp(mysql:3306)/tasks?charset=utf8&parseTime=True&loc=Local", "Database Source Name")
		httpAddr = fs.String("http-addr", ":8080", "HTTP listen address")
		natsAddr = fs.String("nats-addr", "nats://nats:4222", "The NATS server URL")
		natsSub  = fs.String("nats-subject", "simpleNews", "The NATS subject")
	)

	fs.Parse(os.Args[1:])

	cfg := Config{
		DSN:      *dsn,
		HTTPAddr: *httpAddr,
		NATSAddr: *natsAddr,
		NATSSub:  *natsSub,
	}

	return cfg
}
