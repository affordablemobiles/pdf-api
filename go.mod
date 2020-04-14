module github.com/a1comms/pdf-api

go 1.13

require (
	cloud.google.com/go/storage v1.6.0
	github.com/a1comms/go-gaelog/v2 v2.0.0
	github.com/gorilla/mux v1.7.4
	github.com/mattetti/filebuffer v1.0.0
	github.com/unidoc/unipdf/v3 v3.0.0-00010101000000-000000000000
	github.com/urfave/negroni v1.0.0
)

replace github.com/unidoc/unipdf/v3 => github.com/a1comms/unipdf/v3 v3.6.3
