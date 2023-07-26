env:
	- @echo !! make env !!
	- @echo ==================================
	- @echo starting environment for the application, pulling and starting images for psql and redis
	- @echo ==================================
	- go mod tidy
	- go get ./...
	- docker build --no-cache -t owlery:latest .
	- docker-compose up -d

start:
	- @go run main.go

unit-test:
	- @go test ./...

build:
	- go mod tidy
	- go get ./...
	- go build
	- export MAILCHIMP_KEY=<mailchimp-apikey>
	- export OMETRIA_KEY=<ometria-apikey>
	- ./owlery