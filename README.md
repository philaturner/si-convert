# SI Convert

A simple api service written in Go to convert between SI units

*Url* https://si.philturner.dev/convert/cm/m/2500

## Coverage

```
cm -> [m, km, yd, ft]
in -> [cm, m, km, yd, ft]
ft -> [m, cm, km, yd]
m -> [cm, km, yd, ft]
km -> [cm, m, yd, ft]
```

## Usage

*Convert 2500 cm to m*

*GET*`/convert/cm/m/2500`
```json
{
  "message": "Conversion successful",
  "result": 25,
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
```text
=== RUN   TestUnsupportedConversion
--- PASS: TestUnsupportedConversion (0.00s)
=== RUN   TestMeterConversions
--- PASS: TestMeterConversions (0.00s)
=== RUN   TestCentimeterConversions
--- PASS: TestCentimeterConversions (0.00s)
=== RUN   TestInchConversions
--- PASS: TestInchConversions (0.00s)
=== RUN   TestFootConversions
--- PASS: TestFootConversions (0.00s)
=== RUN   TestKilometerConversion
--- PASS: TestKilometerConversion (0.00s)
PASS
```