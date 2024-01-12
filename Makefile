init:
	@npm install
	@go install github.com/a-h/templ/cmd/templ@latest

tailwind:
	@./node_modules/.bin/tailwindcss -i ./assets/css/input.css -o ./assets/css/tw.css --minify

build:
	@make tailwind
	@~/go/bin/templ generate
	@go build cmd/main.go

run:
	@go run cmd/main.go

fmt:
	@templ fmt .
	@go fmt ./...

docker-build:
	@docker build --name website .

docker-run:
	@docker run -p 3000:3000 website

docker-kill:
	@docker kill website && docker image rm website 
