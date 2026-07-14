// Package archive defines a generic archive file-reading function.
// Support for individual formats (ZIP, tar, ...) is not built in;
// it is plugged in by blank-importing the relevant subpackage
// (e.g. gobook/Ch10/10.2/archive/zip), which registers itself with
// RegisterFormat in an init function. This mirrors the mechanism
// image.RegisterFormat uses for image codecs (gopl.io §10.5).
package archive

import (
	"fmt"
	"io"
)

// File is a single file extracted from an archive.
type File struct {
	Name string
	Body []byte
}

// MatchFunc reports whether data (a prefix, or the whole archive) is
// encoded in the associated format.
type MatchFunc func(data []byte) bool

// DecodeFunc decodes the contents of an archive.
type DecodeFunc func(data []byte) ([]File, error)

type format struct {
	name   string
	match  MatchFunc
	decode DecodeFunc
}

var formats []format

// RegisterFormat registers a decoder for an archive format so that
// Read and Sniff can recognize it.
func RegisterFormat(name string, match MatchFunc, decode DecodeFunc) {
	formats = append(formats, format{name, match, decode})
}

// Sniff reports the name of the registered format that recognizes
// data, or "" if none does.
func Sniff(data []byte) string {
	for _, f := range formats {
		if f.match(data) {
			return f.name
		}
	}
	return ""
}

// Read reads an entire archive of unspecified format from r and
// returns its files along with the name of the format that was
// used to decode it. The format must have been registered by
// blank-importing its package.
func Read(r io.Reader) (files []File, formatName string, err error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, "", err
	}
	formatName = Sniff(data)
	if formatName == "" {
		return nil, "", fmt.Errorf("archive: unrecognized format")
	}
	for _, f := range formats {
		if f.name == formatName {
			files, err = f.decode(data)
			return files, formatName, err
		}
	}
	panic("archive: inconsistent format registration")
}
