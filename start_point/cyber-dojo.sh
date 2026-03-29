#!/bin/bash
GOPROXY=off GONOSUMDB='*' GOCACHE=/go/build-cache go test
