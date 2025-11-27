# minimeter

## Install

```
docker run --rm -v minimeter-data:/app/data -p 8080:8080 ghcr.io/jshawl/minimeter:main
```

## Usage

POST /api/measure

```json
{
  "name": "metric.name.here",
  "value": 234
}
```

`value` is optional. Defaults to `1`.

GET /api/metrics

```json
[
  {
    "name": "metric.name.here",
    "count": 41,
    "sum": 41,
    "average": 156,
    "min": 12,
    "max": 2340
  }
]
```
