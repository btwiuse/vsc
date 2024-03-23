# vsc

[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/btwiuse/vsc?tab=doc)
[![Go 1.18+](https://img.shields.io/github/go-mod/go-version/btwiuse/vsc)](https://golang.org/dl/)
[![License](https://img.shields.io/github/license/btwiuse/vsc?color=%23000&style=flat-round)](https://github.com/btwiuse/vsc/blob/main/LICENSE)

```
$ go install github.com/btwiuse/vsc/cmd/vsc@latest
```

```
$ vsc -h
Usage of serveWebArgs:
  -connection-token string
        A secret that must be included with all requests.
  -connection-token-file string
        A file containing a secret that must be included with all requests.
  -extensions-dir string
        Set the root path for extensions.
  -host string
        Host to listen on, defaults to '127.0.0.1' (default "127.0.0.1")
  -log string
        Log level: {off,critical,error,warn,info,debug,trace}, defaults to 'off'. (default "off")
  -port int
        Port to listen on, defaults to 0. If 0 is passed a random free port is picked.
  -quality string
        Quality: {insider,stable,exploration}, defaults to 'insider' (default "insider")
  -relay value
        Relay URL, can be specified multiple times.
  -server-base-path string
        Specifies the path under which the web UI and the code server is provided.
  -server-data-dir string
        Specifies the directory that server data is kept in.
  -socket-path string
        The path to a socket file for the server to listen to.
  -user-data-dir string
        Specifies the directory that user data is kept in. Can be used to open multiple distinct instances of Code.
  -verbose
        Verbose logging.
  -version
        Show version.
  -without-connection-token
        Run without a connection token. Only use this if the connection is secured by other means.
```
