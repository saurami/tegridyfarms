test_all:
	echo "Executing all tests"
	go test server.go server_test.go -v

test_hello:
	echo "Executing tests for /hello endpoint"
	go test -run TestHelloHandler -v

test_health:
	echo "Executing tests for /health endpoint"
	go test -run TestHealthHandler -v

test_outdoor:
	echo "Executing tests for /outdoor endpoint"
	go test -run TestOutdoorHandler -v
