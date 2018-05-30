REPO=github.com/t-tashiro/yakiniku

## generate swagger
gen-swagger:
	# make swagger file
	swagger -apiPackage="$(REPO)" -mainApiFile="(REPO)/main.go" -output=./swagger-ui
	sed -i -e "s/package main/package swaggerui/g" ./swagger-ui/docs.go
    sed -i -e "s/\"basePath\": \"http:\/\/localhost:8081\/swagger-ui\"/\"basePath\": \"http:\/\/localhost:8081\"/g" swagger-ui/docs.go
    sed -i -e "1,/\"basePath\": \"http:\/\/localhost:8081\"/s/\"basePath\": \"http:\/\/localhost:8081\"/\"basePath\": \"http:\/\/localhost:8081\/swagger-ui\"/" swagger-ui/docs.go
	