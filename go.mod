module github.com/a1comms/pdf-api

go 1.13

require (
	cloud.google.com/go v0.102.0 // indirect
	cloud.google.com/go/storage v1.22.1
	github.com/a1comms/go-gaelog/v2 v2.0.1
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/gorilla/mux v1.8.0
	github.com/mattetti/filebuffer v1.0.1
	github.com/sirupsen/logrus v1.8.1 // indirect
	github.com/stretchr/testify v1.7.1 // indirect
	github.com/unidoc/unipdf/v3 v3.0.0-00010101000000-000000000000
	github.com/urfave/negroni v1.0.0
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/image v0.0.0-20220413100746-70e8d0d3baa9 // indirect
	golang.org/x/net v0.0.0-20220526153639-5463443f8c37 // indirect
	golang.org/x/oauth2 v0.0.0-20220524215830-622c5d57e401 // indirect
	google.golang.org/api v0.81.0 // indirect
	google.golang.org/genproto v0.0.0-20220527130721-00d5c0f3be58 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/unidoc/unipdf/v3 => github.com/a1comms/unipdf/v3 v3.6.3
