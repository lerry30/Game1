package layers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"io/fs"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"game1/utils"
)

type TileSetFloorData struct {
	Img string `json:"image"`
}

type TileSetFloor struct {
	img *ebiten.Image
	gid int
}

func (t *TileSetFloor) Img(id int) *ebiten.Image {
	id -= t.gid

	srcX := id % 22
	srcY := id / 22

	srcX *= 16
	srcY *= 16

	return t.img.SubImage(
		image.Rect(srcX, srcY, srcX+16, srcY+16),
	).(*ebiten.Image)
}

func (t *TileSetFloor) GetImageObject(path string) error {
	// 1. Extract the image path in the tileset json file
	data, err := fs.ReadFile(gameFs, path)
	if err != nil {
		return fmt.Errorf("Failed to read file(%s): %w", path, err)
	}

	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)

	var tileSetFloorData TileSetFloorData

	if err := decoder.Decode(&tileSetFloorData); err != nil {
		return fmt.Errorf("Failed to parse json(TileSetFloorData) at layer_interface.go: %w", err)
	}

	// 2. Prepare file path string
	imagePath := utils.ParsePath1(tileSetFloorData.Img)
	imagePath = filepath.Join("assets/", imagePath)

	// 3. Initialize a new ebiten image object with the image file path
	img, _, err := ebitenutil.NewImageFromFile(imagePath)
	if err != nil {
		return err
	}

	// 4. Save to struct
	t.img = img

	return nil
}
