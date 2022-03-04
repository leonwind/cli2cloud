# Create Typescript gRPC Definitions:

1. Copy service.proto in this directory: 
```bash
cp ../../../service/api/proto/service.proto ./
```
2. Generate the gRPC files:
```bash
protoc -I=./ service.proto \
    --js_out=import_style=commonjs,binary:./ \
    --grpc-web_out=import_style=typescript,mode=grpcwebtext:./
```
