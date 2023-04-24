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
	
	commonEnvConfig "gitlab.heronodes.io/bc-platform/bc-wallet-common-lib-config/pkg/config"
	commonLogger "gitlab.heronodes.io/bc-platform/bc-wallet-common-lib-logger/pkg/logger"
	commonVault "gitlab.heronodes.io/bc-platform/bc-wallet-common-lib-vault/pkg/vault"
	commonVaultTokenClient "gitlab.heronodes.io/bc-platform/bc-wallet-common-lib-vault/pkg/vault/client/token"

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

**bc-wallet-common-lib-logger** has a proprietary license.

Switched to proprietary license from MIT - [CHANGELOG.MD - v0.0.12](./CHANGELOG.md)