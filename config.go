package geolocation

import (
	"fmt"

	"github.com/joaosoft/logger"

	"github.com/joaosoft/manager"
)

// AppConfig ...
type AppConfig struct {
	GeoLocation *GeoLocationConfig `json:"geo-location"`
}

// GeoLocationConfig ...
type GeoLocationConfig struct {
	Log struct {
		Level string `json:"level"`
	} `json:"log"`
	Api string `json:"api"`
}

// NewConfig ...
func NewConfig(config ...interface{}) (*AppConfig, manager.IConfig, error) {
	appConfig := &AppConfig{}
	simpleConfig, err := manager.NewSimpleConfig(fmt.Sprintf("/config/app.%s.json", GetEnv()), appConfig)

	if appConfig.GeoLocation == nil {
		appConfig.GeoLocation = &GeoLocationConfig{}
		appConfig.GeoLocation.Log.Level = logger.ErrorLevel.String()
	}

	if len(config) > 0 {
		appConfig.GeoLocation.Api = config[0].(string)
	}

	return appConfig, simpleConfig, err
}
