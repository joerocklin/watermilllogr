# watermilllogr

[![Go Report Card](https://goreportcard.com/badge/github.com/joerocklin/watermilllogr)](https://goreportcard.com/report/github.com/joerocklin/watermilllogr)

This package provides a `watermill.LoggerAdapter` that wraps a [`go-logr/logr` Logger](https://github.com/go-logr/logr).

## Notes

- `Debug` calls `logr.Logger.V(1).Info(...)`
- `Trace` called `logr.Logger.V(2).Info(...)`
