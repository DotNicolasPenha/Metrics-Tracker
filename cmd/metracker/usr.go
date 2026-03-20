package metracker

import (
	"os"
	"path/filepath"

	"github.com/DotNicolasPenha/Metrics-Tracker/core"
)

type Usr struct {
	metrackerpath   string
	dbHandler       *core.DataBaseHandler
	driverDBHandler *core.DBDriverHandler
}

func NewUser() (*Usr, error) {
	configDir, err := os.UserConfigDir()
	if err != nil {
		return nil, err
	}

	metrackerdirpath := filepath.Join(configDir, "metracker")

	err = os.MkdirAll(metrackerdirpath, 0755)
	if err != nil {
		return nil, err
	}

	dbHandler, err := core.NewDataBaseHandler(metrackerdirpath)
	if err != nil {
		return nil, err
	}
	driverDBHandler, err := core.NewDBDriverHandler(dbHandler)
	if err != nil {
		return nil, err
	}

	return &Usr{
		driverDBHandler: driverDBHandler,
		dbHandler:       dbHandler,
		metrackerpath:   metrackerdirpath,
	}, nil
}
