package frontend

import (
	"embed"
	"io/fs"
	"path"
)

//go:embed dist/*
var Dist embed.FS

func NewCustomFS() CustomFS {
	return CustomFS{
		content: Dist,
	}
}

func NewAssetFs() CustomAssetsFS {
	return CustomAssetsFS{
		content: Dist,
	}
}

type CustomFS struct {
	content embed.FS
}

func (c CustomFS) Open(name string) (fs.File, error) {
	return c.content.Open(path.Join("dist", name))
}

type CustomAssetsFS struct {
	content embed.FS
}

func (c CustomAssetsFS) Open(name string) (fs.File, error) {
	return c.content.Open(path.Join("dist/assets", name))
}
