// Package zip plugs ZIP support into the archive package.
// Importing it for its side effect registers the format:
//
//	import _ "gobook/Ch10/10.2/archive/zip"
package zip

import (
	"archive/zip"
	"bytes"
	"io"

	"gobook/Ch10/10.2/archive"
)

func init() {
	archive.RegisterFormat("zip", match, decode)
}

// ZIP archives begin with the local file header signature "PK\x03\x04".
func match(data []byte) bool {
	return len(data) >= 4 && string(data[:4]) == "PK\x03\x04"
}

func decode(data []byte) ([]archive.File, error) {
	zr, err := zip.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		return nil, err
	}
	var files []archive.File
	for _, zf := range zr.File {
		rc, err := zf.Open()
		if err != nil {
			return nil, err
		}
		body, err := io.ReadAll(rc)
		rc.Close()
		if err != nil {
			return nil, err
		}
		files = append(files, archive.File{Name: zf.Name, Body: body})
	}
	return files, nil
}
