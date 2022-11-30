echo "Creating new hemp products"

curl -vX POST 127.0.0.1:8080/product \
  --data @hemp_products.json \
  --header "Content-Type: application/json"


echo "Deleting organic house blend"

curl -vX DELETE 127.0.0.1:8080/products/Organic%20House%20Blend \
  --header "Content-Type: application/json"
