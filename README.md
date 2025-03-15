# Go-task (Clean Architecture)
A simple task management api built using Golang, Gin and GORM with simple and clean architecture

## Architecture Explanation (EN)

- `/api/v1` : The directory for things related to API like all available endpoints (route) and the handlers for each endpoints (controller). Subdirectory `/v1` is used for easy version control in case of several development phase.

  - `/controller` : The directory for things related to the Controller layer which is the part of program that handle requests and return responses.
  - `/router` : The directory for things related with routing. Therefore filled with every available supported routes / endpoints along with the request method and used middleware.

- `/common` : The directory for common things that are frequently used all over the architectures.

  - `/constants` : The directory for base things such as variables, constants, and functions to be used in other directories. It consists of things like response, request, and model base structure.
  - `/interfaces` : The directory for interfaces to give structure to request and response
  - `/middleware` : The directory for Middlewares which are mechanism that intercept a HTTP request and response process before handled directly by the controller of an endpoint.
  - `/util` : The directory to store utility / helper functions that can be used in other directories.

- `/models` : The directory for things related to entities / models which are available on the database via migration that are represented by structs.

- `/dto` : The directory to store DTO (Data Transfer Object) which is a placeholder for other objects, mainly to store data for requests and responses.

- `/services` : The directory for things related to the Service layer which is the layer that is responsible for the flow / business logic of the app.

- `/database`: The directory for things related to the database for example migrations and seeders.

  - `/seeder` : The directory for things related to database seeding for each entity.

- `/docs`: The directory for things related to the docs for example swagger.


## Pre-requisites
1. Create the database in PostgreSQL with the name equal to the value of DB_NAME in `.env`
2. Download [Air](https://github.com/air-verse/air) for hot reloading

## How to Run?

1. Use the command `make tidy` to adjust the dependencies accordingly
2. Seed the data using the command `go run ./database/seed.go `
2. Use the command `air` to run the application