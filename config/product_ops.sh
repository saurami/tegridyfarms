echo "Get all products"

curl -X GET 127.0.0.1:8080/products


echo "Get single product ('Tegridy Weed')"

curl -X GET 127.0.0.1:8080/products/Tegridy%20Weed


echo "Creating new hemp product"

curl -i -X POST 127.0.0.1:8080/product \
  --header "Content-Type: application/json" \
  --data '{"Name":"Hemp Milk","Updated":"2022-11-22 11:12:11"}'


echo "Deleting 'Organic House Blend'"

curl -vX DELETE 127.0.0.1:8080/products/Organic%20House%20Blend \
  --header "Content-Type: application/json"


echo "Updating date for 'Hemp Milk'"

curl -vX PUT 127.0.0.1:8080/products/Hemp%20Milk \
  --header "Content-Type: application/json" \
  --data '{"Name":"Hemp Milk","Updated":"2021-11-10 00:11:11"}'
