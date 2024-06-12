package protobuf

import (
	"embed"
	"io/fs"
	"net/http"
)

//go:embed dist
var dist embed.FS

func SwagServe(prefix string) http.Handler {
	fSys, err := fs.Sub(dist, "dist")
	if err != nil {
		panic(err)
	}
	staticServer := http.FileServer(http.FS(fSys))
	return http.StripPrefix(prefix, staticServer)
}
