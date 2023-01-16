### Local Mode

Traefik also offers a developer mode that can be used for temporary testing of plugins not hosted on GitHub.
To use a plugin in local mode, the Traefik static configuration must define the module name (as is usual for Go packages) and a path to a [Go workspace](https://golang.org/doc/gopath_code.html#Workspaces), which can be the local GOPATH or any directory.

The plugins must be placed in `./plugins-local` directory,
which should be in the working directory of the process running the Traefik binary.
The source code of the plugin should be organized as follows:

```
 ├── docker-compose.yml
 └── plugins-local
    └── src
        └── github.com
            └── ghnexpress
                └── traefik-compress
                    ├── main.go
                    ├── vendor
                    ├── go.mod
                    └── ...

```

### K8s

```yaml
# cache-middleware.yaml
apiVersion: traefik.containo.us/v1alpha1
kind: Middleware
metadata:
  annotations: {}
  name: ghn-compress
  namespace: default
spec:
  plugin:
    plugin-compress:
      excludedContentTypes:
        - text/event-stream
      minResponseBodyBytes: 1024
```

