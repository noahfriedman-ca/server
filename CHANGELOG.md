# CHANGELOG

### unreleased
- Add [`api`](https://github.com/noahfriedman-ca/server/blob/v1.2.0/api) package
- Add [API router](https://github.com/noahfriedman-ca/server/blob/v1.2.0/api/api.go#L26)
- Add [`ListAvailable`](https://github.com/noahfriedman-ca/server/blob/v1.2.0/api/api.go#L45) API function
- Remove `sitemap.xml`

### 1.1.0
- Add [`LICENSE` route](https://github.com/noahfriedman-ca/server/blob/v1.1.0/router.go#L18)

### 1.0.1
- Update [README.md](https://github.com/noahfriedman-ca/server/blob/v1.0.1/README.md)

### 1.0.0
- [Route `/` to `./static/build/`](https://github.com/noahfriedman-ca/server/blob/v1.0.0/router.go#L15)
- [Route `/sitemap.xml` to `./sitemap.xml`](https://github.com/noahfriedman-ca/server/blob/v1.0.0/router.go#L14)
- [Route `/projects/*` to `./projects/*/build/`](https://github.com/noahfriedman-ca/server/blob/v1.0.0/router.go#L11)
