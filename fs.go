package main

import (
	"embed"
	"io/fs"
	"path"
)

type CustomFS struct {
	content embed.FS
}

func (c CustomFS) Open(name string) (fs.File, error) {
	return c.content.Open(path.Join("frontend/build", name))
}
