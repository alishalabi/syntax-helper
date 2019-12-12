#
# Makefile
#
VERSION = 1.0.0
GHRFLAGS =
.PHONY: build release .scannerwork
default: build
build:
	goxc -d=pkg -pv=$(VERSION) -bc="darwin"
release:
	ghr  -u alishalabi  $(GHRFLAGS) v$(VERSION) pkg/$(VERSION)