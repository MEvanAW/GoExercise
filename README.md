# GoExercise
This is a repository for my assignments in "Scalable Web Service with Golang" course from Hacktiv8 × Kominfo.
## Contents
- [Biodata](#biodata)
- [Order API](#order-api)
  - [Order API Swagger](#order-api-swagger)
- [Water and Wind Indicator](#water-and-wind-indicator)
- [MyGram](#mygram)
  - [Entity Relationship Diagram](#entity-relationship-diagram)
  - [MyGram Swagger](#mygram-swagger)
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
### Order API Swagger
The Swagger specification defines a set of files required to describe such an API. These files can then be used by the Swagger-UI project to display the API and Swagger-Codegen to generate clients in various languages. Additional utilities can also take advantage of the resulting files, such as testing tools.<br/>Swagger is available for this API server. To open swagger UI, open http://localhost:8080/swagger/index.html.
![OrderApiSwagger](https://user-images.githubusercontent.com/50491841/198496355-e49233a8-65c4-4304-93de-4946082fc37f.png)
## Water and Wind Indicator
Serves a page to display water height and wind speed. Data generated pseudo-randomly.
### Run it
1. Open "WaterWind" folder
2. `go run .` in Terminal
3. Open localhost:8080 in browser
### Screenshot
![WaterAndWind](https://user-images.githubusercontent.com/50491841/195780089-c55b1ccc-5832-4f76-ae80-cb44576fce42.png)
## MyGram
A social media API to save and comment photos.
### Entity Relationship Diagram
<p align="center">
  <img src="https://github.com/Faqihyugos/mygram-go/blob/main/assets/images/drawSQL-export-2022-10-16_13_06.png" height="520">
</p>

### Prerequisite
1. Install postgresql if you haven't. Alternatively you can use other RDBMS, but you would need to replace gorm postgres driver with its respective gorm driver.
2. Create a database with a name of "mygram"
### Run the Server
1. Open "MyGram" folder.
2. `go run .` in Terminal.
### MyGram Swagger
Swagger is available for this API server. To open swagger UI, open http://localhost:8080/swagger/index.html. Request paths and bodies strictly follows kode.id materials.
![MyGramSwagger](https://user-images.githubusercontent.com/50491841/198496160-a46c743c-3729-4fa3-b667-f46bf6e671a1.png)
#### Authorize Swagger
1. Register user if you haven't.
2. Execute login user.
![LoginUserInSwagger](https://user-images.githubusercontent.com/50491841/198511046-bbbfeeab-a84a-4145-bca3-6a589b633ea5.png)
3. Copy the token response.
4. Authorize with value "Bearer yourToken".
<p align="center">
  <img src="https://user-images.githubusercontent.com/50491841/198511545-c09ef917-b271-47db-a374-b3a9098c811d.png" height="270">
</p>
