# Change Log

## [v0.0.11] - 27.09.2024
### Added
* Support of new lib-errors method - ErrNoWrap
* New settings for wrapcheck linter

## [v0.0.8, v0.0.9, v0.0.10] - 26.09.2024
### Added
* Added golangci linter settings
* Added support error value type to slog.Handler implementation for zap.Logger
* Added new root service-component for create logger entry with tuple-base fields(zap.Field, slog.Attr)
### Changed
* Changed linter rules
* Cleaned up code by linter rules
* Removed support of options flow from zap.Logger implementation of slog.Handler

## [v0.0.7] - 25.09.2024
### Added
* Added slog Logger support
    * Added slog Logger fabric method for making new instances of slog.Logger
    * Added slog Logger sub-package
### Changed
* Refactored std log package
* Refactored and moved zap.Logger to separated package
* Refactored root bc-wallet-common-lib-logger package
* Changed example section in README.md file 

## [v0.0.5 - v0.0.6] - 27.07.2024
### Added 
* Added std Logger fabric method for making new instances of log.Logger 
### Changed
* Changed MIT License to NON-AI MIT
* Added NON-AI MIT copyright banner to all *.go files
* Changed example section in README.md file

## [v0.0.4] - 16.04.2024
### Changed
* Bump golang version 1.19 -> 1.22
* Bump zap.Logger version
* Added support of new config manager dependencies(support of ldflasg https://github.com/crypto-bundle/bc-wallet-common-lib-config)

## [v0.0.3] - 09.02.2024 14:22 MSK
### Changed
* Replaced content of LICENSE file - to MIT license

## [v0.0.2] - 09.02.2024 12:56 MSK
### Changed
* Removed graylog, sentry and bugsnag config fields from default config

## [v0.0.1] - 18.02.2023 20:38 MSK
### Changed
* Separated logger in another repository - https://github.com/crypto-bundle/bc-wallet-common-lib-logger. 
* Extended main logger default fields
* Removed ld-flags from default logger fields
### Added
* Licensed under MIT-license

[@gudron (Kotelnikov Aleksei)](https://github.com/gudron) - author and maintainer of [crypto-bundle project](https://github.com/crypto-bundle)

The commit is signed with the key -
gudron2s@gmail.com
E456BB23A18A9347E952DBC6655133DD561BF3EC