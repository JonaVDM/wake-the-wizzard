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
	return c.content.Open(path.Join("frontend/dist", name))
}

type CustomAssetsFS struct {
	content embed.FS
}

func (c CustomAssetsFS) Open(name string) (fs.File, error) {
	return c.content.Open(path.Join("frontend/dist/assets", name))
}
