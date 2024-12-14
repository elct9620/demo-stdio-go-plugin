plugin-%:
	@echo "Building plugin-$*"
	@go build -C ./plugins/$* -o $(PWD)/plugin-bin/$* .
