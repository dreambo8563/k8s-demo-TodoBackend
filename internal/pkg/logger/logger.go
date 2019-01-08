package logger

import (
	"encoding/json"

	"go.uber.org/zap"
)

// Logger - exported instantce
// var _logger *zap.Logger

//Client -
type Client struct {
	logger *zap.Logger
}

var _logger *Client

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
	initialLogger()
}

func initialLogger() {
	if err = json.Unmarshal(rawJSON, &cfg); err != nil {
		panic(err)
	}
	logger, err := cfg.Build()
	if err != nil {
		panic(err)
	}
	_logger = &Client{
		logger: logger,
	}
}

//Logger - get
func Logger() *Client {
	if _logger != nil {
		initialLogger()
	}
	return _logger
}

//Sugar -
func (c *Client) Sugar() *zap.SugaredLogger {
	return c.logger.Sugar()
}

//Info -
func (c *Client) Info(msg string, fields ...zap.Field) {
	c.logger.Info(msg, fields...)
}

//Error -
func (c *Client) Error(msg string, fields ...zap.Field) {
	c.logger.Error(msg, fields...)
}

//Panic -
func (c *Client) Panic(msg string, fields ...zap.Field) {
	c.logger.Panic(msg, fields...)
}

//Fatal -
func (c *Client) Fatal(msg string, fields ...zap.Field) {
	c.logger.Fatal(msg, fields...)
}

//StringField -
func (c *Client) String(key string, val string) zap.Field {
	return zap.String(key, val)
}

//Sync -
func (c *Client) Sync() error {
	return c.logger.Sync()
}
