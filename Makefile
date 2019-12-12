#
# Makefile
#
VERSION = snapshot
GHRFLAGS =
.PHONY: build release .scannerwork
default: build
build:
	goxc -d=pkg -pv=$(VERSION) -bc="darwin"
release:
	ghr  -u alishalabi  $(GHRFLAGS) v$(VERSION) pkg/$(VERSION)