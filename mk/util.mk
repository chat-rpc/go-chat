# util functions makefile
OS ?= $(shell sh -C 'uname -s 2>/dev/null || echo not')

ifeq ($(OS),Windows_NT)
	WINDOWS := 1
	EXT_SUFFIX := .exe
else
	EXT_SUFFIX :=
endif

SUPPORT_PLATFORMS += windows-386
SUPPORT_PLATFORMS += windows-amd64

SUPPORT_PLATFORMS += linux-386
SUPPORT_PLATFORMS += linux-amd64

SUPPORT_PLATFORMS += darwin-386
SUPPORT_PLATFORMS += darwin-amd64

#debug target ,prints variable.
print-% : 
	@echo $*=$($*)

