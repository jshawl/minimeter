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

```
$ minimeter ls

COUNTERS
http.requests                    47,392
http.requests (status=200)       43,621
http.requests (status=404)        2,341
user.signups                        127
deployment                            8
errors                               94

MEASUREMENTS                 avg      min      max
http.response_time (ms)      156       12    2,340
db.query_time (ms)            23        3      487
build.time (s)               342      298      401
memory.usage (MB)            247      198      312
```
