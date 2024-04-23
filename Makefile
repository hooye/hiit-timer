
main=main.go

## api: 启动
api:
	@go run ${main}
	
## diff: 比对代码
diff:
	@git config diff.nodiff.command trues
	@git diff

.PHONY: help
## help: prints this help message
help:
	@echo "Usage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"