.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test ./...

.PHONY: coverage
coverage:
	go test -cover ./...

.PHONY: html
html:
	go test ./... -coverprofile=coverage.out
	go tool cover -html=coverage.out -o coverage.html

.PHONY: vendor
vendor:
	go mod tidy
	go mod vendor

.PHONY: gen
gen:
	go generate ./...

.PHONY: web
web:
	docker build -f Dockerfile -t web .
	docker run -e ENV=local -p 80:8080 -v ~/.config/gcloud:/root/.config/gcloud web

.PHONY: deploy
deploy:
	docker build -f Dockerfile -t asia.gcr.io/but-la/test/web .
	docker push asia.gcr.io/but-la/test/web
	gcloud run services update web-test --platform=managed --image=asia.gcr.io/but-la/test/web --region=australia-southeast2
