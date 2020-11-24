# Seminario Go
###### Cryptocurrency

## Getting started

Console command line
```
go run cmd/cryptocurrency/cryptocurrency.go -config ./config/config.yaml
```

## Rest

| Methods       | Endpoint                          | 
| ------------- |:--------------------------------: |
| GET           | localhost:8080/cryptocurrency     |
| GET           | localhost:8080/cryptocurrency/:id |
| POST          | localhost:8080/cryptocurrency/    |
| PUT           | localhost:8080/cryptocurrency/    |
| DELETE        | localhost:8080/cryptocurrency/:id |

#### Post body parameters

```
{
  "type": "Bitcoin",
  "quantity": 12
}
```

#### Put body parameters

```
{
  "ID": 1
  "type": "Bitcoin",
  "quantity": 18
}
```
