
.PHONY: build
build:
	@echo "  >  \033[32mBuilding binary...\033[0m "
	cd cmd && env GOARCH=arm64 go build -o ../build/HI