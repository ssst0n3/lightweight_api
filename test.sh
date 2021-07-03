#!/bin/bash
#go test -count=1 -parallel 1 -v ./...
# https://github.com/golang/go/issues/33038
for s in $(go list ./...); do if ! go test -failfast -v -p 1 $s; then break; fi; done