# ============================= #
# Language Support		#
# =============================	#
LANG_SUPPORT += go
LANG_SUPPORT += cpp
LANG_SUPPORT += java
# LANG_SUPPORT += nodejs
# LANG_SUPPORT += c


ifdef -lang
	PROTOC_LANG := $(lang)
else
	@echo "No set protoc language ,used default golang compile."
	PROTOC_LANG := go
endif

# .PHONY : setProtocLang





