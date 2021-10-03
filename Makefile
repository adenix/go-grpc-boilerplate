clean:
	rm -rf gen/go gen/openapiv2 bin/

lint:
	buf lint
	buf breaking --against "https://github.com/adenix/go-grpc-boilerplate.git#branch=main"

generate: clean lint
	buf generate
