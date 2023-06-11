# SHORT1URL

<img align="center" src="public/architecture.png" alt="architecture_img">

## How to run in local

1. Clone this repo

```bash
git clone https://github.com/thuongtruong1009/short1url.git
```

2. Run Docker container

```bash
docker-compose up -d
```

3. Testing API

```bash
# with testing GUI (postman, hoppscotch, thunder client, etc)
POST http://localhost:3000/api
body: {
    "url": "<your_original_url>"
}

# with curl
curl -X POST 'http://localhost:3000/api'
     -H 'Content-Type: application/json' 
     -d '{"url": "<your_original_url>"}"
```

## References

[Logo](https://icon-icons.com/icon/link/102162)
[Ref1](https://liamhieuvu.com/url-shortener-with-golang-and-mysql)