.DEFAULT_GOAL := help

.PHONY: *
%:
	@go run -C .make . $@
