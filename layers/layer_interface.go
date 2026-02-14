package layers

import (
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type Layer interface {
	Img(id int) *ebiten.Image
}

func GetLayer(path string, gid int) (Layer, error) {
	if strings.Contains(path, "buildings") {
		tileSetBuildings := NewLayer()
		tileSetBuildings.gid = gid
		err := tileSetBuildings.GetImageObject(path)
		if err != nil {
			return nil, err
		}

		return tileSetBuildings, nil
	}

	var tileSetFloor TileSetFloor
	tileSetFloor.gid = gid
	err := tileSetFloor.GetImageObject(path)
	if err != nil {
		return nil, err
	}

	return &tileSetFloor, nil
}
