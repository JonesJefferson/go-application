
pb:
	rm -f proto/_generated/*.go
	buf generate

docker:
	go build cmd/user/main.go
	docker build -t user cmd/user
	docker run user

	

