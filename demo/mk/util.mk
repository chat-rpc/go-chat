#util functions
OS ?= $(shell sh -C 'uname -s 2>/dev/null || echo not')
ifeq ($(OS),Windows_NT)
	WINDOWS := 1
	EXT_SUFFIX := .exe
else
	EXT_SUFFIX :=
endif

#debug target,prints varaible. Use like `make print-GOFLAGS`
print-%:
	@echo $*=$($*)
