.PHONY: init
init:
	@npm install
	@go install github.com/a-h/templ/cmd/templ@latest

.PHONY: tailwind
tailwind:
	@./node_modules/.bin/tailwindcss -i ./assets/css/input.css -o ./assets/css/tw.css --minify

.PHONY: css-watch
css-watch:
	@./node_modules/.bin/tailwindcss -i ./assets/css/input.css -o ./assets/css/tw.css --watch

# Build everything
.PHONY: build
build:
	@make tailwind
	@~/go/bin/templ generate
	@go build cmd/main.go

# Run the already built code
.PHONY: run
run:
	@go run cmd/main.go

# Format the templates and code
.PHONY: fmt
fmt:
	@templ fmt .
	@go fmt ./...

.PHONY: docker-build
docker-build:
	@docker build --name website .

.PHONY: docker-run
docker-run:
	@docker run -p 3000:3000 website

.PHONY: docker-kill
docker-kill:
	@docker kill website && docker image rm website 
