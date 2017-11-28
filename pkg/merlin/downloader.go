package merlin

import (
	"image"
)

type Downloader interface {
	DownloadImage(path string) (image.Image, string, error)
}
