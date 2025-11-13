# minimeter

```
$ minimeter increment deployment environment=production
$ minimeter measure build.time 234
```

```
POST /api/increment
{
  "name": "deployment",
  "tags": {"environment": "production"},
}

POST /api/measure
{
  "name": "build.time",
  "value": 234,
  "tags": {"service": "api"},  // optional
}
```
