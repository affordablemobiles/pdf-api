module github.com/a1comms/pdf-api

go 1.13

require (
	cloud.google.com/go v0.103.0 // indirect
	cloud.google.com/go/errorreporting v0.2.0
	cloud.google.com/go/monitoring v1.5.0 // indirect
	cloud.google.com/go/storage v1.23.0
	cloud.google.com/go/trace v1.2.0 // indirect
	contrib.go.opencensus.io/exporter/stackdriver v0.13.13
	github.com/a1comms/go-gaelog/v2 v2.0.1
	github.com/aws/aws-sdk-go v1.44.56 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/mattetti/filebuffer v1.0.1
	github.com/prometheus/prometheus v0.37.0 // indirect
	github.com/unidoc/unipdf/v3 v3.0.0-00010101000000-000000000000
	github.com/urfave/negroni v1.0.1-0.20200608235619-7de0dfc1ff79
	go.opencensus.io v0.23.0
	golang.org/x/image v0.0.0-20220617043117-41969df76e82 // indirect
	golang.org/x/net v0.0.0-20220708220712-1185a9018129 // indirect
	golang.org/x/oauth2 v0.0.0-20220630143837-2104d58473e0 // indirect
	golang.org/x/sys v0.0.0-20220715151400-c0bba94af5f8 // indirect
	google.golang.org/api v0.87.0 // indirect
	google.golang.org/genproto v0.0.0-20220718134204-073382fd740c // indirect
	google.golang.org/grpc v1.48.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/unidoc/unipdf/v3 => github.com/a1comms/unipdf/v3 v3.6.3
