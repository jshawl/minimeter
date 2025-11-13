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

```
$ minimeter get http.response_time --last 24h --chart

http.response_time (ms)                                Last 24 hours

400ms│                                          /\
350ms│                                         /  \
300ms│                                  __    /    \
250ms│                                 /  \__/      \_
200ms│                        __      /               \
150ms│              __       /  \____/                 \_
100ms│     __      /  \____/                             \
 50ms│____/  \____/                                       \___
   0 └──────────────────────────────────────────────────────────────────
       12p  2p  4p  6p  8p  10p  12a  2a  4a  6a  8a  10a  12p
       Nov 11                              Nov 12

       Avg: 156ms    Min: 12ms    Max: 2,340ms    p95: 420ms
```
