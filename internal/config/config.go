package config

import (
	"log"

	"github.com/alexedwards/scs/v2"
)

// Appconfid holds the application config
type AppConfig struct {
	UseCache     bool
	InfoLog      *log.Logger
	ErrorLog     *log.Logger
	Session      *scs.SessionManager
	InProduction bool
}
