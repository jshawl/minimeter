# minimeter

POST /api/measure

```json
{
  "name": "metric.name.here",
  "value": 234
}
```

value is optional, defaults to 1.

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
