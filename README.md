# gocg : go code generation

A minimalist go tool to generate code from templates leveraging on the amazing go template library

## Prerequisites

- [Go 1.15](https://golang.org/dl/)

## Get the tool's binary

### Prebuilt binary

You can refer to the project's [GitHub release page](https://github.com/FrancescoIlario/gocg/releases).

### Using Go

```
go get -u -v github.com/FrancescoIlario/gocg
```

## Run the example

Clone the repository in a local directory

```
git clone https://github.com/FrancescoIlario/gocg.git
cd gocg
```

Run the following command to generate the `simplego` example project from templates

```
gocg ./examples/simplego
```

You can run the generated code using the following commands

```
cd ./out/simplego
go run main.go
```
