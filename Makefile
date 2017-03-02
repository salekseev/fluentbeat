BEAT_NAME=fluentbeat
BEAT_PATH=github.com/salekseev/fluentbeat
BEAT_GOPATH=$(firstword $(subst :, ,${GOPATH}))
BEAT_URL=https://${BEAT_PATH}
SYSTEM_TESTS=false
TEST_ENVIRONMENT=false
ES_BEATS?=./vendor/github.com/elastic/beats
GOPACKAGES=$(shell glide novendor)
PREFIX?=.
NOTICE_FILE=NOTICE

# Path to the libbeat Makefile
-include $(ES_BEATS)/libbeat/scripts/Makefile

# dependencies that are used by the build&test process, these need to be installed in the
# global Go env and not in the vendor sub-tree
DEPEND=golang.org/x/tools/cmd/cover github.com/onsi/ginkgo/ginkgo \
       github.com/onsi/gomega github.com/rlmcpherson/s3gof3r/gof3r \
       github.com/Masterminds/glide github.com/golang/lint/golint

# Initial beat setup
.PHONY: setup
setup: copy-vendor depend
	make update

# Copy beats into vendor directory
.PHONY: copy-vendor
copy-vendor:
	mkdir -p vendor/github.com/elastic/
	cp -R ${BEAT_GOPATH}/src/github.com/elastic/beats vendor/github.com/elastic/
	rm -rf vendor/github.com/elastic/beats/.git

# Installing build dependencies. You will need to run this once manually when you clone the repo
.PHONY: depend
depend:
	go get -v $(DEPEND)
	glide install

# run gofmt and complain if a file is out of compliance
# run go vet and similarly complain if there are issues
# run go lint and complain if there are issues
.PHONY: lint
lint:
	@if gofmt -l . | egrep -v ^vendor/ | grep .go; then \
	  echo "^- Repo contains improperly formatted go files; run gofmt -w *.go" && exit 1; \
	  else echo "All .go files formatted correctly"; fi
	#go tool vet -v -composites=false *.go
	#go tool vet -v -composites=false **/*.go
	for pkg in $$(go list ./... |grep -v /vendor/); do golint $$pkg; done

.PHONY: git-init
git-init:
	git init
	git add README.md CONTRIBUTING.md
	git commit -m "Initial commit"
	git add LICENSE
	git commit -m "Add the LICENSE"
	git add .gitignore
	git commit -m "Add git settings"
	git add .
	git reset -- .travis.yml
	git commit -m "Add fluentbeat"
	git add .travis.yml
	git commit -m "Add Travis CI"

# This is called by the beats packer before building starts
.PHONY: before-build
before-build:

# Collects all dependencies and then calls update
.PHONY: collect
collect:
