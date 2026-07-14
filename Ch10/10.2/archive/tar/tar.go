// Package tar plugs POSIX tar support into the archive package.
// Importing it for its side effect registers the format:
//
//	import _ "gobook/Ch10/10.2/archive/tar"
package tar

import (
	"archive/tar"
	"bytes"
	"io"

	"gobook/Ch10/10.2/archive"
)

func init() {
	archive.RegisterFormat("tar", match, decode)
}

// POSIX (ustar) tar archives carry the magic string "ustar" at byte
// offset 257 of the first header block. Pre-POSIX (v7) tar files have
// no such signature, so they can't be reliably sniffed by content alone.
func match(data []byte) bool {
	return len(data) >= 262 && string(data[257:262]) == "ustar"
}

func decode(data []byte) ([]archive.File, error) {
	tr := tar.NewReader(bytes.NewReader(data))
	var files []archive.File
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
		if hdr.Typeflag != tar.TypeReg {
			continue
		}
		body, err := io.ReadAll(tr)
		if err != nil {
			return nil, err
		}
		files = append(files, archive.File{Name: hdr.Name, Body: body})
	}
	return files, nil
}
