package server

import (
	"io"
	"os"

	"github.com/crazy-max/ftpgrab/v7/internal/model"
)

// Handler is a server interface
type Handler interface {
	Common() model.ServerCommon
	ReadDir(source string) ([]os.FileInfo, error)
	Retrieve(path string, dest io.Writer) error
	Close() error
}

// Client represents an active server object
type Client struct {
	Handler
}
