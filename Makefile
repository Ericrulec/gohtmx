init:
	@npm install
	@go install github.com/a-h/templ/cmd/templ@latest

tailwind:
	@./node_modules/.bin/tailwindcss -i ./assets/css/input.css -o ./assets/css/tw.css

run:
	@make tailwind
	@templ generate
	@go run cmd/main.go

lint:
	@templ fmt ./view/**/*.templ
	@go fmt ./...

docker-build:
	@docker build --name website .

docker-run:
	@docker run -p 3000:3000 website

docker-kill:
	@docker kill website && docker image rm website 
