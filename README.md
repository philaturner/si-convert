# SI Convert

A simple api service written in Go to convert between SI units

*Url* https://si.philturner.dev/convert/cm/m/2500

## Coverage

```
cm -> [m, km, yd, ft, in]
in -> [cm, m, km, yd, ft]
ft -> [m, cm, km, yd, in]
m -> [cm, km, yd, ft]
km -> [cm, m, yd, ft]
°C -> [°F, K]
°F -> [°C, K]
K -> [°C, °F]
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

*Convert 10°C to °F*

*GET* `/convert/c/f/10`
```json
{
  "message": "Conversion successful",
  "result": 50,
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
=== RUN   TestKilometerConversions
--- PASS: TestKilometerConversions (0.00s)
=== RUN   TestTemperatureConversions
--- PASS: TestTemperatureConversions (0.00s)
PASS
```