package controller

import (
	"encoding/json"
	"github.com/saurabh/tegridyfarms/pkg/model"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestGetProducts(t *testing.T) {
	request, err := http.NewRequest("GET", "/products", nil)
	if err != nil {
		t.Fatalf("Unable to reach endpoint: %v", err)
	}

	response := httptest.NewRecorder()

	handler := http.HandlerFunc(GetProducts)
	handler.ServeHTTP(response, request)

	t.Run("Response Code", func(t *testing.T) {
		got := response.Code
		want := http.StatusOK
		if got != want {
			t.Errorf("Incorrect response code ... got %v; want %v", got, want)
		}
	})

	t.Run("Header", func(t *testing.T) {
		got := response.Header().Get("Content-Type")
		want := "application/json"
		if got != want {
			t.Errorf("Incorrect response header ... got %v; want %v", got, want)
		}
	})

	t.Run("Content", func(t *testing.T) {
		err := json.NewDecoder(response.Body).Decode(&model.Products)
		if err != nil {
			t.Errorf("Error parsing response data into defined product structure: %v", err)
		}
	})
}

func TestRetrieveExistingProduct(t *testing.T) {
	request, err := http.NewRequest("GET", "/products/Tegridy%20Weed", nil)
	if err != nil {
		t.Fatalf("Unable to reach endpoint: %v", err)
	}

	response := httptest.NewRecorder()

	handler := http.HandlerFunc(RetrieveProduct)
	handler.ServeHTTP(response, request)

	t.Run("Response Code", func(t *testing.T) {
		got := response.Code
		want := http.StatusOK
		if got != want {
			t.Errorf("Incorrect response header: got %v; want %v", got, want)
		}
	})

	t.Run("Header", func(t *testing.T) {
		got := response.Header().Get("Content-Type")
		want := "application/json"
		if got != want {
			t.Errorf("Unexpected content type: got %v; want %v", got, want)
		}
	})

	t.Run("Content", func(t *testing.T) {
		var singleProduct model.Product
		err := json.NewDecoder(response.Body).Decode(&singleProduct)
		if err != nil {
			t.Errorf("Error parsing response data into product structure: %v", err)
		}
	})
}

func TestCreateProduct(t *testing.T) {
	hempFile, err := os.Open("../../config/hemp_products.json")
	if err != nil {
		t.Errorf("Error opening file: %v", err)
	}
	defer hempFile.Close()

	request, err := http.NewRequest("POST", "/product", hempFile)
	if err != nil {
		t.Errorf("Unable to reach endpoint: %v", err)
	}

	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	response := httptest.NewRecorder()

	handler := http.HandlerFunc(CreateProduct)
	handler.ServeHTTP(response, request)

	t.Run("Response Code", func(t *testing.T) {
		got := response.Code
		want := http.StatusOK
		if got != want {
			t.Errorf("Incorrect response code ... got %v; want %v", got, want)
		}
	})

	t.Run("Header", func(t *testing.T) {
		got := response.Header().Get("Content-Type")
		want := "application/json"
		if got != want {
			t.Errorf("Invalid header ... got %v; want %v", got, want)
		}
	})

	t.Run("Response Content", func(t *testing.T) {
		var newProduct model.Product
		err := json.NewDecoder(response.Body).Decode(&newProduct)
		if err != nil {
			t.Errorf("Error parsing response data into the defined structure: %v", err)
		}
	})
}
