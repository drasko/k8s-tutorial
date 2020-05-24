# Copyright (c) Mainflux
# SPDX-License-Identifier: Apache-2.0

PROGRAM = app
SOURCES = $(wildcard *.go) cmd/main.go
IMAGE = k8s-tutorial/lesson-1

all: $(PROGRAM)

.PHONY: all clean

$(PROGRAM): $(SOURCES)
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o $@ cmd/main.go

docker:
	docker build -t $(IMAGE) .

docker-run:
	docker run -p 3333:3333 $(IMAGE)

clean:
	rm -rf $(PROGRAM)
