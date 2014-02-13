#Copyright (c) 2012 Megam Systems.
#
#   Licensed under the Apache License, Version 2.0 (the "License");
#   you may not use this file except in compliance with the License.
#   You may obtain a copy of the License at
#
#       http://www.apache.org/licenses/LICENSE-2.0
#
#   Unless required by applicable law or agreed to in writing, software
#   distributed under the License is distributed on an "AS IS" BASIS,
#   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#   See the License for the specific language governing permissions and
#   limitations under the License.
###############################################################################
# Makefile to compile seru.
# lists all the dependencies for test, prod and we can run a go build aftermath.
###############################################################################
                            

SERUCODE_HOME = $(HOME)/code/megam/workspace/seru

export GOPATH=$(SERUCODE_HOME)

define HG_ERROR

FATAL: you need mercurial (hg) to download seru dependencies.
       Check README.md for details
endef

define GIT_ERROR

FATAL: you need git to download seru dependencies.
       Check README.md for details
endef

define BZR_ERROR

FATAL: you need bazaar (bzr) to download gulp dependencies.
       Check README.md for details
endef

.PHONY: all check-path get hg git bzr get-test get-prod test client

all: check-path get test

# It does not support GOPATH with multiple paths.
check-path:
ifndef GOPATH
	@echo "FATAL: you must declare GOPATH environment variable, for more"
	@echo "       details, please check README.md file and/or"
	@echo "       http://golang.org/cmd/go/#GOPATH_environment_variable"
	@exit 1
endif


clean:
	@/bin/rm -f -r $(SERUCODE_HOME)/pkg	
	@go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | tr ' ' '\n' |\
		grep '^.*\..*/.*$$' | grep -v 'github.com/indykish/seru' |\
		sort | uniq | xargs -I{} rm -f -r $(SERUCODE_HOME)/src/{}	
	@go list -f '{{range .Imports}}{{.}} {{end}}' ./... | tr ' ' '\n' |\
		grep '^.*\..*/.*$$' | grep -v 'github.com/indykish/seru' |\
		sort | uniq | xargs -I{} rm -f -r $(SERUCODE_HOME)/src/{} 
	@/bin/echo "Clean ...ok"

get: hg git bzr get-test get-prod

hg:
	$(if $(shell hg), , $(error $(HG_ERROR)))

git:
	$(if $(shell git), , $(error $(GIT_ERROR)))

bzr:
	$(if $(shell bzr), , $(error $(BZR_ERROR)))

get-test:
	@/bin/echo -n "Installing test dependencies... "
	@go list -f '{{range .TestImports}}{{.}} {{end}}' ./... | tr ' ' '\n' |\
		grep '^.*\..*/.*$$' | grep -v 'github.com/indykish/seru' |\
		sort | uniq | xargs go get -u >/tmp/.get-test 2>&1 || (cat /tmp/.get-test && exit 1)	
	@/bin/echo "ok"
	@rm -f /tmp/.get-test

get-prod:
	@/bin/echo -n "Installing production dependencies... "
	@go list -f '{{range .Imports}}{{.}} {{end}}' ./... | tr ' ' '\n' |\
		grep '^.*\..*/.*$$' | grep -v 'github.com/indykish/seru' |\
		sort | uniq | xargs go get -u >/tmp/.get-prod 2>&1 || (cat /tmp/.get-prod && exit 1)
	@/bin/echo "ok"
	@rm -f /tmp/.get-prod

_go_test:
	@go test -i ./...
	@go test ./...

test: _go_test 

client:
	@go build -o seru ./cmd/seru
	@echo "Done."