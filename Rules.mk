# target bin
# clean 
# dist clean
#
TARGET_BIN := bin
CLEAN :=
DISTCLEAN :=

# file suffix
EXT_SUFFIX := .exe

#include mk/util.mk

# =====================================	#
# import proto mk			#
# =====================================	#
#include proto/protoc.mk

#.PHONY : all clean protoc buildsvr buildcli

#all : help
#	TARGET_BIN=.
#	@echo $(PROTOC_LANG)
