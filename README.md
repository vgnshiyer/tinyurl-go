# tinyurl-go
An Educational project: A Tiny URL service implementation in golang.

**Tools Used:**
- go-fiber
- redis
- docker, docker-compose

## API Documentation

#### Endpoints

- /api/v1

**Method:** POST

**Sample Body:**
```
{
  "url": "https://www.google.com/asd2352f"
}
```

**Sample Response:**
```
{
  "url": "https://www.google.com/asd2352f",
  "CustomShort": "<DOMAINNAME>/afv23x",
  "Expiry": 24,
  "XRateRemaining": 10,
  "XRateLimitReset": 30
}
```

- /`id`

**Method:** GET

**Response:** Redirects to actual url
