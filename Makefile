# go-chat Makefile
# lanbery language 
SHELL=PATH='$(PATH)' /bin/bash
PRO_NAME := nbs-chat

# =============================================	#
# Variables					#
# =============================================	#
TARGET_BIN = bin

# set language set
#include proto/lang-support.mk	


# =============================================	#
# Make cmd help					#
# =============================================	#
help :
	@echo -e "all		: build all contians protoc,server and client."
	@echo -e "help		: help."
	@echo -e "protoc		: use default golang compile.\n " "\t\t and you can use -lang=go/cpp/java."


#include Rules.mk

.PHONY : all clean


#protoc buildsvr buildcli

# =============================================	#
# all						#
# =============================================	#
all : help init
	@echo $(PROTOC_LANG)

init : 
	@[ -d "./$(TARGET_BIN)" ] || mkdir $(TARGET_BIN)
