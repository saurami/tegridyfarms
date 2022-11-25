test_app:
	echo "Executing application tests (collectively)"
	go test -v ./pkg/controller/app.go ./pkg/controller/app_test.go

test_app_health:
	echo "Executing tests for /health endpoint"
	go test -v ./pkg/controller -run TestHealthHandler

test_hello:
	echo "Executing tests for /hello endpoint"
	go test -v ./pkg/controller -run TestHelloHandler

test_home:
	echo "Executing tests for /home endpoint"
	go test -v ./pkg/controller -run TestHomeHandler

test_outdoor:
	echo "Executing tests for /outdoor endpoint"
	go test -v ./pkg/controller -run TestOutdoorHandler

test_get_products:
	echo "Executing tests for /products endpoint"
	go test -v ./pkg/controller -run TestGetProducts

test_toy_database:
	echo "Executing tests for local SQLite database"
	go test -v ./config/dbconn.go ./config/dbconn_test.go

test_coverage:
	echo "Running test coverage for project"
	go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

generate_coverage_report:
	echo "Generate coverage report (HTML)"
	go test -v -coverprofile cover.out ./...
	go tool cover -html cover.out -o cover.html
