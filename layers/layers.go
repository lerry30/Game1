package layers

import "io/fs"

var gameFs fs.FS

func Init(f fs.FS) {
	gameFs = f
}
