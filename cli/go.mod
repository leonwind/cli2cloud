module github.com/leonwind/cli2cloud/cli

go 1.16

replace github.com/leonwind/cli2cloud/service => ../service

require (
	github.com/leonwind/cli2cloud/service v0.0.0-00010101000000-000000000000
	golang.org/x/crypto v0.0.0-20210711020723-a769d52b0f97
	google.golang.org/grpc v1.44.0
)
