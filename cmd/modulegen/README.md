# Module gen

This package contains code for generating testchain proto and module logic.

## Running it

This should be run in the testchain root directory, as follows:

```bash
make modulegen
```

or

```bash
go run cmd/modulegen/main.go -module_name test_module
make proto-gen
go mod tidy
```
