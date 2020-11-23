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
| UPDATE        | localhost:8080/cryptocurrency/    |

#### Post body parameters

```
{
  "type": "Bitcoin",
  "quantity": 12
}
```
