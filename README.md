# smart-service
Smart API and RPC service based on Kratos T


## Go Module setting

- Check GO111MODULE
    ```bash
    $ go env GO111MODULE
    on
    ```

- Turn on GO111MODULE
    ```bash
    $ go env -w GO111MODULE="on"
    ```

- Setting GOPROXY
    ```bash
    $ go env -w GOPROXY=https://goproxy.cn
    ```

- Check GOMODCACHE
    ```bash
    $ go env GOMODCACHE
    ```

    Setting go mod cache directory

    ```bash
    go env -w GOMODCACHE=$GOPATH/pkg/mod
    ```

## Install kratos
- Check enviroment

    Download kratos to $GOPATH/bin.

- Download and install
    ```bash
    go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
    ```

## Install protoc

Download from [protobuf release](https://github.com/protocolbuffers/protobuf/releases), unzip protoc to $GOPATH/bin.

## Install protoc-gen-go

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

Move protoc-gen-go to $GOPATH/bin.

## Quick Start with kratos

```bash
# mkdir demo
# create project's layout
kratos new demo
cd demo
# modify go.mod module if needed, then pull dependencies
go mod download
go mod tidy
# init enviroment
make init
# generate all codes
make all
# build
make build
# run
make run
# or run
# ./bin/helloworld -conf ./configs
# try it out now
curl 'http://127.0.0.1:8000/helloworld/kratos'
# The response should be
{
  "message": "Hello kratos"
}
```


## Create new service with kratos
```bash
# generate proto template
kratos proto add api/blog/v1/blog.proto
# modify the proto template if needed
# generate proto code
# kratos proto client api/blog/v1/blog.proto
# if no http created, then try following cmd
# kratos proto client api/blog/v1/blog.proto --go-http_opt=omitempty=false
# generate all new codes
make all
# generate server template
#kratos proto server api/blog/v1/blog.proto -t internal/service
```

## Use ent as ORM

- install ent
    ```bash
    go get entgo.io/ent/cmd/ent
    ```

- create schemas
    ```bash
    cd internal/data
    # Create User Schema file under ./ent/schema/ directory, or create schema file manually
    # Notice: should create file generate.go under ./ent, and include 'package ent'
    # go run entgo.io/ent/cmd/ent init User
    # Generate all schemas
    go generate ./ent
    ```

## Customize Layout

```tree
|----cmd
|     |
|     |----project
|           |
|           |----main.go
|           |
|           |----wire.go
|
|----api
|     |
|     |----project
|           |
|           |----*.proto
|
|----internal
|     |
|     |----biz
|     |
|     |----conf
|     |     |
|     |     |----*.proto
|     |
|     |----data
|     |     |
|     |     |----ent
|     |           | 
|     |           |----schema
|     |
|     |----server
|     |
|     |----service
|
|-----Makefile

```