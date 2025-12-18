# if BUILD_ID is unset, compute metadata that will be used in builds
ifeq ($(strip $(BUILD_ID)),)
	VCS_REF := $(shell git rev-parse --short HEAD)
	BUILD_TIME_EPOCH := $(shell date +"%s")
	BUILD_TIME_RFC_3339 := \
		$(shell date -u -r $(BUILD_TIME_EPOCH) '+%Y-%m-%dT%I:%M:%SZ')
	BUILD_TIME_UTC := \
		$(shell date -u -r $(BUILD_TIME_EPOCH) +'%Y%m%d-%H%M%S')
	BUILD_ID := $(VCS_REF)
endif

ifeq ($(strip $(TAG)),)
	TAG := unknown
endif 

.PHONY: clean
clean:
	@echo "Cleaning"
	rm -rf demo/target

.PHONY: metadata
metadata:
	@echo "Gathering Metadata"
	@echo BUILD_TIME_EPOCH IS $(BUILD_TIME_EPOCH)
	@echo BUILD_TIME_RFC_3339 IS $(BUILD_TIME_RFC_3339)
	@echo BUILD_TIME_UTC IS $(BUILD_TIME_UTC)
	@echo BUILD_ID IS $(BUILD_ID)



target/app.jars:
	@echo "Building App Artifacts"
	docker run -it --rm  -v "$(shell pwd)/demo":/project/ -w /project/ \
    maven:3.9.9-eclipse-temurin-21 \
    mvn clean install

.PHONY: app-artifacts
app-artifacts: target/app.jars

.PHONY: lint-dockerfile
lint-dockerfile:
	@set -e
	@echo "Linting Dockerfile"
	docker run --rm -i hadolint/hadolint < multi-stage-runtime.df
	docker run --rm -i hadolint/hadolint < simple-runtime.df
	docker run --rm -i hadolint/hadolint < all-in-one.df
	@echo "Dockerfile linting passed"
	
.PHONY: app-image-simple-runtime
app-image-simple-runtime: metadata lint-dockerfile app-artifacts
	@echo "Building App Image"
	docker image build -t omerurhan/dockercourse:simple-runtime-$(BUILD_ID) \
	-f simple-runtime.df \
	--build-arg BUILD_ID='$(BUILD_ID)' \
	--build-arg BUILD_DATE='$(BUILD_TIME_RFC_3339)' \
	--build-arg VCS_REF='$(VCS_REF)' \
	demo
	@echo "Built Simple Runtime Image. BUILD_ID: $(BUILD_ID)"

.PHONY: app-image-all-in-one
app-image-all-in-one: metadata lint-dockerfile
	@echo "Building All-in One App Image"
	docker image build -t omerurhan/dockercourse:all-in-one-$(BUILD_ID) \
	-f all-in-one.df \
	--build-arg BUILD_ID='$(BUILD_ID)' \
	--build-arg BUILD_DATE='$(BUILD_TIME_RFC_3339)' \
	--build-arg VCS_REF='$(VCS_REF)' \
	demo
	@echo "Built All-in One App Image. BUILD_ID: $(BUILD_ID)"


.PHONY: image-tests-simple-runtime
image-tests-simple-runtime:
	@echo "Testing image structure"
	docker run --rm -it \
	-v /var/run/docker.sock:/var/run/docker.sock \
	-v $(shell pwd)/structure-tests.yaml:/structure-tests.yaml \
	gcr.io/gcp-runtimes/container-structure-test:v1.6.0 test \
	--image omerurhan/dockercourse:simple-runtime-$(BUILD_ID) \
	--config /structure-tests.yaml

.PHONY: inspect-image-labels-simple-runtime
inspect-image-labels-simple-runtime:
	docker image inspect --format '{{ json .Config.Labels }}' \
		omerurhan/dockercourse:simple-runtime-$(BUILD_ID) | jq
