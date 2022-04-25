
.PHONY: build
build:
	@echo "  >  \033[32mBuilding...\033[0m "
	cd cmd && go build -o ../build/HI