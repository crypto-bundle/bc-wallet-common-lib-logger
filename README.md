# bc-wallet-common-lib-logger

## Description

Library for manage logger config and zap.Logger cores

Library contains:
* common logger config struct
* management of zap.Core instances via small wrapper service-component
* management of std logger cores

## Usage example

### Environment variables

* `LOGGER_LEVEL` - MinimalLogsLevel is a level for setup minimal logger information level.
These values of logger level will be passed to zap.Logger setup.
Allowed values: debug, info, warn, error, dpanic, panic, fatal. 
* `LOGGER_STACKTRACE_ENABLE` - Enable logger stacktrace
* `LOGGER_SKIP_BUILD_INFO` - Build zap.Logger cores without application build-info fields

Examples of preparing config and create instance of logger service-component

### Config and connection

```go
package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	commonConfig "github.com/crypto-bundle/bc-wallet-common-lib-config/pkg/config"
	commonLogger "github.com/crypto-bundle/bc-wallet-common-lib-logger/pkg/logger"

	"go.uber.org/zap"
)

type applicationConfig struct {
	*commonConfig.BaseConfig
	*commonLogger.LoggerConfig
}

func main() {
	var err error
	ctx, _ := context.WithCancel(context.Background())

	appCfg := &applicationConfig{}

	baseCfgPreparerSvc := commonConfig.NewConfigManager()

	baseCfg := commonConfig.NewBaseConfig("application_name")
	err = baseCfgPreparerSvc.PrepareTo(baseCfg).Do(ctx)
	if err != nil {
		log.Fatal(err.Error(), err)
	}

	loggerSvc, err := commonLogger.NewService(appCfg)
	if err != nil {
		log.Fatal(err.Error(), err)
	}

	stdLoggerEntry := loggerSvc.NewStdLoggerEntry("stdLoggerName", []any{
		"test_field":  "test_value",
		"test_field1": "test_value2",
	})
	slogLoggerEntry := loggerSvc.NewSlogLoggerEntry("slogLoggerName", []any{
		"test_field":  "test_value",
		"test_field1": "test_value2",
	})
	zapLoggerEntry := loggerSvc.NewZapLoggerEntry("main", []any{
		"test_field":  "test_value",
		"test_field1": "test_value2",
	})

	// usage
	stdLoggerEntry.Print("std logger info")
	slogLoggerEntry.Info("msg", slog.String("key", "value"))
	zapLoggerEntry.Info("zap logger msg", zap.String("key", "value"))

	os.Exit(0)
}
```

## Contributors
* Author and maintainer - [@gudron (Alex V Kotelnikov)](https://github.com/gudron)

## Licence
**bc-wallet-common-lib-logger** is licensed under the [MIT NON-AI](./LICENSE) License.