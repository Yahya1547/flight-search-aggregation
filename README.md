# flight-search-aggregation
Flight Search Aggregation Coding Test BookCabin

## Setup and Run the Application
Here's the step to run the application:
1. Run "go mod tidy"
This command is necessary to download all dependencies used in the project

2. Run "go run main.go"
If the project running successfully, then "Server running at ..." will be visible in the terminal.

3. Hit API Flight Search
http://localhost:8080/flights/search?origin=CGK&destination=DPS&departure_date=2025-12-15&passengers=1&cabin_class=economy&airlines=AirAsia,Batik Air&max_price=1000000&min_price=50000&max_duration=5&min_duration=1&number_of_stops=0,1,2&sort_by=price&sort_direction=desc

Path: /flights/search
Request Query:
- origin
- destination
- departure_date
- passengers
- cabin_class
- max_price (Support Price Range)
- min_price (Support Price Range)
- max_duration (Support Duration Range)
- min_duration (Support Duration Range)
- Airlines (Support filter by airlines): string separated by ","
- sort_by : "price", "duration", "departure_time", "arrival_time"
- sort_direction : "asc", "desc"

## Explanation

### Design Pattern
In this application, the main design pattern I used is Adapter Pattern. This pattern is required to face the challenges of handling multiple data format and normalized into one model that will be used in the application for Flight data. For instance, I create 1 AirlineProvider interface and then implement the interface by creating 4 Different Providers for each Airline provided by the mock APIs. In each specific Airline Provider, the data fetched through the mock API will be transformed into the normalized object.

As to fetching data for each providers, here I create the Flight Aggregator which will call every Airline Provider via the interface. I also implemented concurrent fetching so the queries run asynchronously to increase the performance.

### Caching
Caching implemented in this application using go-cache library, an in-memory caching mecahnism. The flight data are cached based on origin, destination, and date. Cache mechanism located right before the aggregator called the airline providers to reduce duplicated call by the same request within the time range. By doing so, the API can quickly get the data without having to call to external API and can proceed to filter the data based on user's request.

### Structure
This repo is divided into different packages based on the functionality:
- API package
The purpose of this package is to handle user's request and do the filtering and sorting based on the validated data by the aggregator

- Aggregator Package
This package is supposed to handle the caching mechanism and concurrent call to external airline providers.

- Provider Package
This package contains the airline providers, responsible for calling to each external API and send the normalized data to aggregator.

- Service Package
Service Package will serve the purpose to provide supporting function to filter, sorting, or any other functionality required by the API.

- Models Package
This package is basically containing the struct or model across the package, including normalized and specific external model.

- Utils Package
This is a supporting package to cover general functionality, not related to business logic.




