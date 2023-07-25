env:
	- @echo !! make env !!
	- @echo ==================================
	- @echo starting environment for the application, pulling and starting images for psql and redis
	- @echo ==================================
	- docker-compose up -d

start:
	- @go run main.go