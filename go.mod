module hope

go 1.17

//直接依赖
require (
	github.com/go-kratos/kratos/v2 v2.1.4
	github.com/google/wire v0.5.0
	google.golang.org/genproto v0.0.0-20211223182754-3ac035c7e7cb
	google.golang.org/grpc v1.42.0
	google.golang.org/protobuf v1.27.1
)

//间接依赖
require (
	github.com/fsnotify/fsnotify v1.4.9 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0 // indirect
	github.com/imdario/mergo v0.3.12 // indirect
	go.opentelemetry.io/otel v1.0.0 // indirect
	go.opentelemetry.io/otel/trace v1.0.0 // indirect
	golang.org/x/net v0.0.0-20210405180319-a5a99cb37ef4 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20210816074244-15123e1e1f71 // indirect
	golang.org/x/text v0.3.5 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

require entgo.io/ent v0.9.1
