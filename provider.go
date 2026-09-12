package datboxcore

import (
	"io"
	"io/fs"
	"time"
)

type Progress struct {
	StartTime     time.Time
	CurrentBytes  int64
	TotalBytes    int64
	CurrentChunks int
	TotalChunks   int
}

type ProviderInfo struct {
	Root        string // Root path of the file system of this provider
	Encrypted   bool   // Whether the provider is encrypting uploaded files
	StoredFiles uint   // Number of files stored within this provider
	StoredBytes uint   // Total bytes of files stored within this provider
}

type Provider interface {
	Get(path string, writer io.WriteCloser, progressCallback func(progress Progress)) error            // Get a file at the path and write to the writer
	Put(path string, size int64, reader io.ReadCloser, progressCallback func(progress Progress)) error // Put a file to the path using the reader

	List(path string) ([]fs.DirEntry, error)          // List files under a directory
	Move(src, dst string) error                       // Move file from src to dst
	Delete(path string, recursive, remote bool) error // Delete file at the path, optionally recursively and remotely
	Mkdir(path string, parents bool) error            // Create a directory, optionally create parents if they don't exist
	Copy(src, dst string) error                       // Copy file (reference) from src to dst

	Info() (ProviderInfo, error) // Fetch info of the provider
}
