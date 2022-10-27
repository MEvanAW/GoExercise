# GoExercise
This is a repository for my assignments in "Scalable Web Service with Golang" course from Hacktiv8 × Kominfo.
## Contents
- [Biodata](#biodata)
- [Order API](#order-api)
  - [Swagger](#swagger)
- [Water and Wind Indicator](#water-and-wind-indicator)
- [MyGram](#mygram)
## Biodata
Prints name, address, job, and reason to join course of the student with id passed in args.
### Run the Code
1. Open "Biodata" folder.
2. `go run . <student id>` in Terminal.
valid student id: 0-49
## Order API
API server in Go for orders of items.
### Prerequisite
1. Install postgresql if you haven't. Alternatively you can use other RDBMS, but you would need to replace gorm postgres driver with its respective gorm driver.
2. Create a database with a name of "order-api"
### Run the Server
1. Open "OrderApi" folder.
2. `go run .` in Terminal.
### Swagger
The Swagger specification defines a set of files required to describe such an API. These files can then be used by the Swagger-UI project to display the API and Swagger-Codegen to generate clients in various languages. Additional utilities can also take advantage of the resulting files, such as testing tools.<br/>Swagger is available for this API server. To open swagger UI, open http://localhost:8080/swagger/index.html.
![image](https://user-images.githubusercontent.com/50491841/195003066-b1418298-c253-4b8f-a06f-51cd57913abf.png)
## Water and Wind Indicator
Serves a page to display water height and wind speed. Data generated pseudo-randomly.
### Run it
1. Open "WaterWind" folder
2. `go run .` in Terminal
3. Open localhost:8080 in browser
### Screenshot
![image](https://user-images.githubusercontent.com/50491841/195780089-c55b1ccc-5832-4f76-ae80-cb44576fce42.png)
## MyGram
A social media
### Prerequisite
1. Install postgresql if you haven't. Alternatively you can use other RDBMS, but you would need to replace gorm postgres driver with its respective gorm driver.
2. Create a database with a name of "mygram"
### Run the Server
1. Open "OrderApi" folder.
2. `go run .` in Terminal.
### Swagger
Swagger is available for this API server. To open swagger UI, open http://localhost:8080/swagger/index.html. The API implementation is already complete, although the swagger docs is still work in progress. Request paths and bodies strictly follows kode.id materials though, so the rest of the format can be read there.
