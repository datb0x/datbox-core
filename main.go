package main

import (
	"io"
	"io/fs"
)

type ProviderInfo struct {
	Root        string
	Encrypted   bool
	StoredFiles uint
	StoredBytes uint
}

type Provider interface {
	Get(writer io.Writer) error
	Put(reader io.Reader) error

	List(path string) ([]fs.DirEntry, error)
	Move(src, dst string) error
	Delete(path string) error
	Mkdir(path string) error
	Copy(src, dst string) error

	Info() (ProviderInfo, error)
}
