# [noahfriedman.ca](https://noahfriedman.ca)/server [![CI/CD](https://github.com/noahfriedman-ca/server/workflows/CI/CD/badge.svg)](https://github.com/noahfriedman-ca/server/actions?query=workflow%3ACI%2FCD) [![Go Reference](https://pkg.go.dev/badge/github.com/noahfriedman-ca/server.svg)](https://pkg.go.dev/github.com/noahfriedman-ca/server)
The main server used on [noahfriedman.ca](https://noahfriedman.ca).

### Routes:
- `/`: If the request is for the base URL, route to the `./static/build/` path
- `/sitemap.xml`: If the request is for the sitemap, it serves the sitemap (located at path `./sitemap.xml`)
- `/projects/*`: If the request is for a subdirectory of `projects`, it returns the content located at `./projects/__requested__/build/`

All other requests result in a `404`.

