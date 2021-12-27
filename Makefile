full:
	go test -v ./... -coverprofile cover.out -bench=. -benchmem

docker-build:
	docker build --tag multiply-matrices .

docker-run:
	 docker run multiply-matrices
