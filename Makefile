TEST?=$$(go list ./... | grep -v 'vendor')
HOSTNAME=registry.terraform.io
NAMESPACE=dell
NAME=powerstore
BINARY=terraform-provider-${NAME}
VERSION=1.2.1
OS_ARCH=linux_amd64
OPENAPI_CMD?=java -Xmx16G -jar ../terraform-provider-powerscale/openapi-generator-cli-6.6.0.jar
OPENAPI_GEN_DIR=clientgen

default: install

build_client:
	python3 goClientZip/spec.py
	rm -rf ${OPENAPI_GEN_DIR}
	${OPENAPI_CMD} generate -i goClientZip/spec_4_1_filtered.json -g go --type-mappings integer+unsigned64=uint64  -o ${OPENAPI_GEN_DIR} --global-property apis,models,supportingFiles=client.go:README.md:configuration.go:response.go:utils.go,modelTests=false,apiTests=false,modelDocs=false -p enumClassPrefix=true,packageName=clientgen,isGoSubmodule=true -c config.yaml
	#${OPENAPI_CMD} generate -i goClientZip/spec_4_1_filtered.json -g go --type-mappings integer+unsigned64=uint64  -o goclient --global-property supportingFiles=false,apis,apiTests=false -p enumClassPrefix=true,packageName=clientgen,isGoSubmodule=true
	cd ${OPENAPI_GEN_DIR} && goimports -w .

unused:
	# add openapi-generator-cli as java -jar ./openapi-generator-cli-6.6.0.jar
	# --ignore-file-override=/root/terraform-provider-powerstore/.openapi-generator-ignore
	# java -jar openapi-generator-cli.jar generate -i /root/terraform-provider-powerscale/goClientZip/PowerScale_API_9.5.0.json --global-property apis --openapi-normalizer FILTER="tag:Filesystem|Filepool" -g go --type-mappings integer+unsigned64=uint64 -o goclient
	#="/platform/1/filepool/policies" ="/volume_group"
	# --openapi-normalizer FILTER="tag:volume|volume_group"
	# ../terraform-provider-powerscale/openapi-generator-cli-6.6.0.jar
	# ../terraform-provider-powerscale/openapi-generator-cli-6.6.0.jar
	# --global-property apis="VolumeAPI" --global-property models="VolumeInstance|VolumeAttach|ErrorResponse|VolumeClone|VolumeCloneResponse|VolumeConfigureMetro|VolumeConfigureMetroResponse|VolumeDelete|VolumeDetach|VolumeEndMetro|VolumeModify|VolumeRefresh|VolumeRefreshResponse|VolumeRestore|VolumeRestoreResponse|VolumeSnapshot|VolumeSnapshotResponse|VolumeCreate|CreateResponse"

build:
	go build -o ${BINARY}

release:
	GOOS=darwin GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_darwin_amd64
	GOOS=freebsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_freebsd_386
	GOOS=freebsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_freebsd_amd64
	GOOS=freebsd GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_freebsd_arm
	GOOS=linux GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_linux_386
	GOOS=linux GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_linux_amd64
	GOOS=linux GOARCH=arm go build -o ./bin/${BINARY}_${VERSION}_linux_arm
	GOOS=openbsd GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_openbsd_386
	GOOS=openbsd GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_openbsd_amd64
	GOOS=solaris GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_solaris_amd64
	GOOS=windows GOARCH=386 go build -o ./bin/${BINARY}_${VERSION}_windows_386
	GOOS=windows GOARCH=amd64 go build -o ./bin/${BINARY}_${VERSION}_windows_amd64


install: build
	rm -rfv ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	find examples -type d -name ".terraform" -exec rm -rfv "{}" +;
	find examples -type f -name "trace.*" -delete
	find examples -type f -name "*.tfstate" -delete
	find examples -type f -name "*.hcl" -delete
	find examples -type f -name "*.backup" -delete
	rm -rf trace.*
	
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}

uninstall:
	rm -rfv ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${NAME}/${VERSION}/${OS_ARCH}
	find examples -type d -name ".terraform" -exec rm -rfv "{}" +;
	find examples -type f -name "trace.*" -delete
	find examples -type f -name "*.tfstate" -delete
	find examples -type f -name "*.hcl" -delete
	find examples -type f -name "*.backup" -delete
	rm -rf trace.*


test: check
	go test -i $(TEST) || exit 1                                                   
	echo $(TEST) | xargs -t -n4 go test $(TESTARGS) -timeout=30s -parallel=4                    

check:
	terraform fmt -recursive examples/
	gofmt -s -w .
	golangci-lint run --fix --timeout 5m
	go vet

gosec:
	gosec -quiet ./...

testacc:
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 120m   

generate:
	go generate ./...

cover:
	rm -f coverage.*
	go test -coverprofile=coverage.out ./...
	go tool cover -html coverage.out -o coverage.html

all: test gosec testacc generate cover install
