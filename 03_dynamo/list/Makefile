build:
	GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap main.go

zip:
	zip lambda-handler.zip bootstrap

clean:
	rm bootstrap lambda-handler.zip
	
do:
	make build
	make zip