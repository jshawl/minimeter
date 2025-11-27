# minimeter

POST /api/measure

```json
{
  "name": "metric.name.here",
  "value": 234 // optional, defaults to 1
}
```

GET /api/metrics

```json
[
  {
    "name": "metric.name.here",
    "count": 41,
    "sum": 41, // same as count when value == 1
    "average": 156,
    "min": 12,
    "max": 2340
  }
]
```
