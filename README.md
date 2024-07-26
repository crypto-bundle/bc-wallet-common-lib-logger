# bc-wallet-common-lib-logger

## Description

Library for manage logger config and zap.Logger cores

Library contains:
* common logger config struct
* management of zap.Core instances via small wrapper service-component

## Usage example

Examples of preparing config and create instance of logger service-component

### Config and connection

```go
package main

import (
	"context"
	"log"

	commonConfig "github.com/crypto-bundle/bc-wallet-common-lib-config/pkg/config"
	commonLogger "github.com/crypto-bundle/bc-wallet-common-lib-logger/pkg/logger"
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
	stdLoggerFabric := loggerSvc.NewStdLogMaker()
	zapLoggerEntry := loggerSvc.NewLoggerEntry("main")
	stdLoggerEntry := stdLoggerFabric.WithFields("stdLoggerName", map[string]interface{}{
		"test_field":  "test_value",
		"test_field1": "test_value2",
	})

	stdLoggerEntry.Print("std logger info")
	// usage
	zapLoggerEntry.Info("zap logger info")
}

}

```


## Licence

**bc-wallet-common-lib-logger** is licensed under the [MIT](./LICENSE) License.