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
