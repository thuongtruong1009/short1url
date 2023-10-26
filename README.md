<div align="center">
    <h1><img src="public/logo.png" alt="logo"> SHORT1URL</h1>
    <img alt="Build status" src="https://img.shields.io/github/actions/workflow/status/thuongtruong1009/short1url/build.yml?logo=GitHub&label=build">
    <img alt="Test status" src="https://img.shields.io/github/actions/workflow/status/thuongtruong1009/short1url/test.yml?logo=GitHub&label=test">
    <img alt="CircleCI status" src="https://circleci.com/gh/circleci/circleci-docs.svg?style=svg">
    <a href="https://github.com/thuongtruong1009/short1url/pkgs/container/short1url-api"><img alt="Automated api build" src="https://img.shields.io/docker/automated/thuongtruong1009/short1url-api?logo=Docker&label=server"></a>
    <a href="https://github.com/thuongtruong1009/short1url/pkgs/container/short1url-client"><img alt="Automated client build" src="https://img.shields.io/docker/automated/thuongtruong1009/short1url-client?logo=Docker&label=client"></a>
    <img alt="GitHub code size in bytes" src="https://img.shields.io/github/languages/code-size/thuongtruong1009/short1url">
    <a href="https://github.com/thuongtruong1009/short1url/blob/main/LICENSE"><img alt="License" src="https://img.shields.io/github/license/thuongtruong1009/short1url"></a>
    <a href="https://paypal.me/thuongtruong1009" rel="nofollow"><img src="https://camo.githubusercontent.com/30c9a9ce3120b1eeb5cac34c303f02145c7f6997b4cd6d8faa049e98e1714ae0/68747470733a2f2f696d672e736869656c64732e696f2f62616467652f446f6e6174652d50617950616c2d6666336635392e737667" data-canonical-src="https://img.shields.io/badge/Donate-PayPal-ff3f59.svg" style="max-width: 100%;"></a>
     <!-- <img alt="api image size" src="https://img.shields.io/docker/image-size/thuongtruong1009/short1url-api/latest">
    <img alt="client image size" src="https://img.shields.io/docker/image-size/thuongtruong1009/short1url-client/latest"> -->
</div>

## Description

This is a simple URL shortener service. It is written in Golang and uses Redis as database. Other hand, it also provides some services such as QR code generator, barcode generator, etc.

## Preview

![Preview image](public/preview.png)

## What's new

- [x] Shorten URL
- [x] Redirect to original URL
- [x] Expiration time
- [x] Statistics
- [x] Rate limit
- [x] QR code generator (custom color, download image)
- [x] Barcode generator
- [x] Microservices Dockerize
- [x] Auto build and deploy image
- [x] Reverse proxy
- [ ] Unit test
- [ ] Caching

## Architecture

![](public/architecture.png)

## Getting started

1. Clone this repo

```bash
git clone https://github.com/thuongtruong1009/short1url.git
```

2. Fill in environment variables

```bash
# client
cd api && cp .env.example .env

# server
cd client && cp .env.example .env
```

3. Run Docker container

```bash
docker-compose up -d
```

4. Testing API

```bash
# with browser
Open http://localhost:81/s
```

```bash
# with Postman or browser
POST http://localhost:81/s
body: {
    "url": "<your_original_url>"
}
```

```bash
# with curl
curl -X POST 'http://localhost:81/s'
     -H 'Content-Type: application/json'
     -d '{"url": "<your_original_url>"}"
```

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change. If you like my work, please star 🌟 this repository.

## License

**Short1url** is an [MIT-licensed](LICENSE) open source project.

Copyright of <a href="https://github.com/thuongtruong1009">thuongtruong1009</a>

<!-- ## References

[Ref1](https://liamhieuvu.com/url-shortener-with-golang-and-mysql)
[Go on K8s](https://www.callicoder.com/deploy-multi-container-go-redis-app-kubernetes/)
[Nginx cache](https://vietnix.vn/cau-hinh-cache-nginx/)
[Nginx refs](https://github.dev/veryacademy/yt-nginx-mastery-series)
-->
