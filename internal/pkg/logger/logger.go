package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

// Logger - exported instantce
var Logger *zap.Logger

var (
	cfg     zap.Config
	err     error
	rawJSON = []byte(`{
		"level": "debug",
		"encoding": "json",
		"outputPaths": ["stdout", "./todobackend-log"],
		"errorOutputPaths": ["stderr"],
		"initialFields": {"service": "todo-backend"},
		"encoderConfig": {
			"messageKey": "message",
			"levelKey": "level",
			"timeKey": "timestamp",
	    	"timeEncoder": "ISO8601",
			"levelEncoder": "lowercase"
		}
	}`)
)

func init() {
	if err = json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	Logger, err = cfg.Build()
	if err != nil {
		panic(err)
	}

}
