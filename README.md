# CoG : Code Generator

A minimalist go tool to generate code from templates leveraging on the amazing go template library

## Prerequisites

- [Go 1.17](https://golang.org/dl/)

## Get the tool's binary

### Prebuilt binary

You can refer to the project's [GitHub release page](https://github.com/FrancescoIlario/cog/releases).

### Using Go

```
go get -u -v github.com/FrancescoIlario/cog
```

## Run the example

Clone the repository in a local directory

```
git clone https://github.com/FrancescoIlario/cog.git
cd cog
```

To build the application use

```
make build
```

Only on Linux, you can install the binary by using

```
make install
```

Run the following command to generate the `simplego` example project from templates

```
cog ./examples/simplego
```

You can run the generated code using the following commands

```
cd ./out/simplego
go run main.go
```
