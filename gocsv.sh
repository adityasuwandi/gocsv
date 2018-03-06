#!/bin/bash

# (curl --write-out %{http_code} --silent --output /dev/null localhost:3000/index)
# (curl --write-out %{http_response} --silent --output /dev/null localhost:3000/index)
# response=$(curl -sb -H "Accept: application/json" "http://host:8080/some/resource")
# response=$(curl -sb -H "Accept: application/json" "http://0.0.0.0:3000/index")
# database="inventory.db"
# echo $response

# get response from goventory
response=$(curl -sb -H "Accept: application/json" "http://0.0.0.0:3000/index")

# check goventory, it should be running on Port :3000
if [ "$response" != "Inventory REST API." ]; then
  echo "Something wrong dude, goventory is not running on Port :3000."
  exit
fi

# When somebody press Ctrl-C
trap '{ echo "Hey, you pressed Ctrl-C. Time to quit." ; exit 1; }' INT

# start gocsv microservice
exec `PORT=4000 go run main.go`
