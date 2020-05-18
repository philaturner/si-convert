# SI Convert

A simple api service written in Go to convert between SI units

## Coverage

```
cm -> [m, km, yd, ft]
in -> [cm, m, km, yd, ft]
ft -> [m, cm, km, yd]
m -> [cm, km, yd, ft]
```

## Usage

`/convert/cm/m/2500`
```json
{
  "message": "Conversion successful",
  "data": 25,
  "status": 200
}
```

## Docker

`docker build -t si-convert .`

`docker run -d --name si-convert-service -p8001:8080 si-convert`

## Tests

To run all tests from project root
```
go test ./... -v
```