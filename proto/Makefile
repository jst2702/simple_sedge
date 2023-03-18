.PHONY: gen-go,gen-py

default:

install-tools:
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc \
	github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc
	yarn install --frozen-lockfile

clean:
	rm -r gen/ || true
	rm -r dist/ || true
	rm -r ../intel-web/proto || true

# (optionally, specify a folder and the end) 
# buf generate --template buf.gen.go.yaml feed/webapis
gen-go:
	buf generate --template buf.gen.go.yaml

gen-py:
	buf generate --template buf.gen.py.yaml

gen-ts:
	buf generate --template buf.gen.ts.yaml

gen: gen-go gen-py gen-ts

# build:
# 	poetry build
