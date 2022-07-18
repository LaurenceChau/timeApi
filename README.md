# Time API

return the time of the given timezones.

## Example 1
```
curl http://localhost:8080/api/time

# response
{
    "current_time": "2021-08-09 11:18:06 +0000 UTC"
}
```

## Example 2

```
curl  http://localhost:8080/api/time?tz=America/New_York,Asia/Kolkata

# response
{
    "Asia/Kolkata": "2021-08-09 01:23:42 +0530 IST",
    "America/New_York": "2021-08-09 01:23:42 -0400 EDT"
}

```