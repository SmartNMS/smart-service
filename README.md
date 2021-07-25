# smart-service
Smart API and RPC service


## Go Module 设置

- 查看GO111MODULE开启情况
    ```bash
    $ go env GO111MODULE
    on
    ```

- 开启GO111MODULE，如果已开启（即执行go env GO111MODULE结果为on）请跳过。
    ```bash
    $ go env -w GO111MODULE="on"
    ```

- 设置GOPROXY
    ```bash
    $ go env -w GOPROXY=https://goproxy.cn
    ```

- 查看设置GOMODCACHE
    ```bash
    $ go env GOMODCACHE
    ```

    如果目录不为空或者/dev/null，请跳过。

    ```bash
    go env -w GOMODCACHE=$GOPATH/pkg/mod
    ```

## Goctl 安装
- 环境变量检测

    下载编译后的二进制文件位于$GOPATH/bin目录下，要确保$GOPATH/bin已经添加到环境变量。
- download&install
    ```bash
    go get -u github.com/tal-tech/go-zero/tools/goctl
    ```

## kratos 安装
- 环境变量检测

    下载编译后的二进制文件位于$GOPATH/bin目录下，要确保$GOPATH/bin已经添加到环境变量。
- download&install
    ```bash
    go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest
    ```

## protoc 安装

进入 [protobuf release](https://github.com/protocolbuffers/protobuf/releases) 页面，选择适合自己操作系统的压缩包文件下载，并将解压后的 protoc 二进制文件移动 $GOPATH/bin 下面。(也可以放到环境变量的任意path下面)

## protoc-gen-go 安装

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

确保 protoc-gen-go 在 $GOPATH/bin 下。(也可以放到环境变量的任意path下面)

## Quick Start with kratos

- Project Creation

```
# create project's layout
kratos new helloworld

cd helloworld
# pull dependencies
go mod download
# generate proto template
kratos proto add api/helloworld/demo.proto
# generate proto code
kratos proto client api/helloworld/demo.proto
# generate server template
kratos proto server api/helloworld/demo.proto -t internal/service
```

- Compilation and Running
```bash
# generate all codes of proto and wire etc.
go generate ./...

# run the application
kratos run
```

- Try it out
```bash
curl 'http://127.0.0.1:8000/helloworld/kratos'

The response should be
{
  "message": "Hello kratos"
}
```
