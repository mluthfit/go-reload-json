# Objective

Build a microservice to change a JSON file every 15 seconds with random numbers between 1-100 for water and wind values and create a page to display data according to the rules and always up-to-date (auto reload).

`JSON file`

```json
{
  "status": {
    "water": 6,
    "wind": 8
  }
}
```

## Rules

- If the water is below 5 then the status is safe
- If the water is between 6 - 8 then the status is standby
- If the water is above 8 then the status is dangerous
- If the wind is below 6 then the status is safe
- If the wind is between 7 - 15 then the status is standby
- If the wind is above 15 then the status is dangerous
- Value of water in meters
- Value wind in meters per second

## Getting Started

This project developed using `go 1.19`.

To start running this project locally, you must follow these steps:

First, clone these repository to the your folder.

```
> https://github.com/mluthfit/go-reload-json.git
```

Then, open the folder and **install** all packages.

```
> go mod tidy
```

Then, start the server.

```
> go run main.go
```
