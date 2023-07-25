env:
	- @echo !! make env !!
	- @echo ==================================
	- @echo starting environment for the application, pulling and starting images for psql and redis
	- @echo ==================================
	- docker build --no-cache -t owlery:latest .
	- docker-compose up -d

start:
	- @go run main.go