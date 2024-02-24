# hotel-rates-api
This API interacts with the Hotelbeds API to retrieve rate information 


## Local Set up

1. Clone this repo using https 
    `git clone https://github.com/antenehhaile/hotel-rates-api.git`
2. Run `go mod download` to download all the dependencies
3. Set up you local env   
`export API_KEY=${API_KEY}`   
`export SECRET=${API_SECRET}` 

## Running the API locally

#### Run the App directly
```go run main.go```
#### Run the App using docker
```docker build -t ${image_name} .```   
```docker run -p 8080:8080 ${image_name}```  

Then the API should run on port `8080`

