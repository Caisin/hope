# D ocument
[kratos文档](https://go-kratos.dev/docs/getting-started/start/?target="blank")

[ent文档](https://entgo.io/zh/docs/getting-started)
# ent 生成代码

```shell
cd ./apps/admin/internal/data && ent generate ./ent/schema
cd ./apps/novel/internal/data && ent generate ./ent/schema
cd ./apps/param/internal/data && ent generate ./ent/schema
```

# Kratos Project Template

## Install Kratos
```
## Install Kratos
go get -u github.com/go-kratos/kratos/cmd/kratos/v2@latest

## Install ent 
go get -u entgo.io/ent/cmd/ent

# install  protobuf 
https://github.com/protocolbuffers/protobuf/releases 

# install  wire 
go get github.com/google/wire/cmd/wire
```

## Create a service

``` shell
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```

## Generate other auxiliary files by Makefile

```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```

## Automated Initialization (wire)

```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker

```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```

