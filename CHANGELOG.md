# CHANGELOG

### 1.2.1
- Fix [routing for paths with no trailing slash](https://github.com/noahfriedman-ca/server/blob/v1.2.1/router.go#L21)

### 1.2.0
- Completely rewrite [project routing](https://github.com/noahfriedman-ca/server/blob/v1.2.0/router.go)
- [Route a `dev` folder](https://github.com/noahfriedman-ca/server/blob/v1.2.0/router.go#L16) alongside a `projects` folder

### 1.1.2
- Update [project routing](https://github.com/noahfriedman-ca/server/blob/v1.1.2/router.go#L18)

### 1.1.1
- Change [LICENSE](https://github.com/noahfriedman-ca/server/blob/v1.1.1/LICENSE) from MPL v2.0 to AGPL v3.0

### 1.1.0
- Add [`LICENSE` route](https://github.com/noahfriedman-ca/server/blob/v1.1.0/router.go#L18)

### 1.0.1
- Update [README.md](https://github.com/noahfriedman-ca/server/blob/v1.0.1/README.md)

### 1.0.0
- [Route `/` to `./static/build/`](https://github.com/noahfriedman-ca/server/blob/v1.0.0/router.go#L15)
- [Route `/sitemap.xml` to `./sitemap.xml`](https://github.com/noahfriedman-ca/server/blob/v1.0.0/router.go#L14)
- [Route `/projects/*` to `./projects/*/build/`](https://github.com/noahfriedman-ca/server/blob/v1.0.0/router.go#L11)
