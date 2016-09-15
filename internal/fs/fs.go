package fs

import (
	"io"
	"os"
	"path/filepath"
)

type Exportable interface {
	ExportTo(destination string) error
}

type WriterTo interface {
	WriteAllTo(io.Writer) error
}

type File struct {
	Name     string
	Contents WriterTo
}

func (f *File) ExportTo(destination string) error {
	filename := filepath.Join(destination, f.Name)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	err = f.Contents.WriteAllTo(file)
	if closeErr := file.Close(); err == nil {
		err = closeErr
	}
	return err
}

type Directory struct {
	Name     string
	Children []Exportable
}

func (d *Directory) ExportTo(destination string) error {
	dir := filepath.Join(destination, d.Name)
	if err := os.RemoveAll(dir); err != nil {
		return err
	}
	if err := os.Mkdir(dir, 0766); err != nil {
		return err
	}
	for _, child := range d.Children {
		if err := child.ExportTo(dir); err != nil {
			return err
		}
	}
	return nil
}

var _ Exportable = (*Directory)(nil)
var _ Exportable = (*File)(nil)
