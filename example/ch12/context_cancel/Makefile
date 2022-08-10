.DEFAULT_GOAL := build

.PHONY: build run-ok run-cancel

build:
	go build

run-ok: build
	./context_cancel false

run-cancel: build
	./context_cancel true
