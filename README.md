# Timezone API
Simple REST API for getting current time in different timezones. This is the first assignment from [Rest-Based Microservices API Development in Golang Course](https://www.udemy.com/course/rest-based-microservices-api-development-in-go-lang/).

## Run Program
```Go
go run main.go
```

## Example Endpoints

### `GET /api/time` (UTC as default timezone)

Response

```JSON
{
    "current_time": "2021-10-31 13:14:55.0414006 +0000 UTC"
}
```

### `GET /api/time?tz=Asia/Jakarta` (Single timezone)

Response

```JSON
{
    "Asia/Jakarta": "2021-10-31 20:16:26.2831456 +0700 WIB"
}
```

### `GET /api/time?tz=Asia/Jakarta,America/New_York` (Multiple timezones)

Response

```JSON
{
    "America/New_York": "2021-10-31 09:17:17.4514516 -0400 EDT",
    "Asia/Jakarta": "2021-10-31 20:17:17.4514159 +0700 WIB"
}
```

### `GET /api/time?tz=Asia/Bandung` (Invalid timezone)

Response

```JSON
{
    "error_message": "invalid timezone"
}
```