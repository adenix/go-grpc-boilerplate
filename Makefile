lint:
	buf lint
	buf breaking --against "https://github.com/adenix/go-grpc-boilerplate.git#branch=main"

generate:
	buf generate

clean:
	rm -rf gen/ bin/
