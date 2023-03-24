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
	"time"
	
	commonEnvConfig "github.com/crypto-bundle/bc-wallet-common-lib-config/pkg/envconfig"
	commonLogger "github.com/crypto-bundle/bc-wallet-common-lib-logger/pkg/logger"
	commonVault "github.com/crypto-bundle/bc-wallet-common-lib-vault/pkg/vault"
	commonVaultTokenClient "github.com/crypto-bundle/bc-wallet-common-lib-vault/pkg/vault/client/token"

	"go.uber.org/zap"
)

type VaultWrappedConfig struct {
	*commonVault.BaseConfig
	*commonVaultTokenClient.AuthConfig
}

func main() {
	ctx := context.Background()

	// vault prepare
	vaultSrv, err := commonVault.NewService(ctx, vaultCfg, vaultClientSrv)
	if err != nil {
		panic(err)
	}

	_, err = vaultSrv.Login(ctx)
	if err != nil {
		panic(err)
	}

	// logger config prepare
	loggerConfig := commonLogger.LoggerConfig{}
	pgCfgPreparerSrv := commonEnvConfig.NewConfigManager()
	err = pgCfgPreparerSrv.PrepareTo(loggerConfig).With(vaultSrv).Do(ctx)
	if err != nil {
		panic(err)
	}

	// logger instance service creation
	loggerSrv, err := commonLogger.NewService(appCfg)
	if err != nil {
		log.Fatal(err.Error(), err)
	}
	loggerEntry := loggerSrv.NewLoggerEntry("main") // zap.Logger("go.uber.org/zap") instance will be returned here

	// usage 
	loggerEntry.Info("application started successfully", zap.Time("start time", time.Now()))
}

```

## Licence

**bc-wallet-common-lib-postgres** is licensed under the [MIT](./LICENSE) License.