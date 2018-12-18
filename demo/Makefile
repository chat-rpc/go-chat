# Makefile 
# bin dest ./bin
# make all : protoc_go/protoc_java/protoc_cpp/server/client
# make protoc_go
# make sever/client

# 
# extension define default window
EXT_SUFFIX := exe
include mk/util.mk

pro_name=nbs-chat
MDP = mkdir -p

#MKFILE_PATH := $(abspath $(lastword $(MAKEFILE_LIST)))
BASEDIR += .
bin_out := bin
PATH := $(bin_out):$(PATH)
SHELL := /bin/bash

.PHONY: all mkout clean protoc_go chatsvr

all: mkout protoc_go chatsvr

# @echo $(EXT_SUFFIX)

mkout:
	@[ -d "./$(bin_out)" ] || mkdir $(bin_out)

# compile protobuf to go
proto_package = helloworld
protoc_exc = $GOPATH/bin/protoc-gen-go
protoc_go:
	protoc -I ./$(proto_package)/ --go_out=./$(proto_package) ./$(proto_package)/*.proto

# server
svr_dir = server
svr_exec := $(pro_name)-svr$(EXT_SUFFIX)
chatsvr: protoc_go
	go build -a -x -o $(bin_out)/$(svr_exec) ./$(svr_dir)/; \
	echo "chat_svr exec complete."

# client 
cli_dir=client
cli_exec=$(pro_name)-cli$(EXT_SUFFIX)
chatcli: protoc_go
	go build -a -x o $(bin_out)/$(cli_exec) ./$(cli_dir)/; \
	echo "chat_client exec complete."
	
help:
	@echo -e "all		: build all"
	@echo -e "clean		: clean "
	@echo -e "protoc_go	: protoc $(proto_package)/$(proto_package).proto "
	@echo -e "chatsvr\t\t: build server exec by golang \n " \
		"\t\t and the exec file in $(bin_out)/$(svr_exec)"
	@echo -e "chatcli\t\t: build go default client \n " \
		"\t\t and the exec file in $(bin_out)/$(svr_exec)"

#clean
clean:
	rm -rf bin/*
	rm -f $(proto_package)/$(proto_package).pb.go







