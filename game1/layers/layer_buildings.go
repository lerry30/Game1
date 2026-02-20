package layers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/fs"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"game1/utils"
)

type TileSetCollectionData struct {
	ID        int    `json:"id"`
	ImgPath   string `json:"image"`
	ImgWidth  int    `json:"imagewidth"`
	ImgHeight int    `json:"imageheight"`
}

type TileSetCollection struct {
	Tile []TileSetCollectionData `json:"tiles"`
}

type TileSetBuildings struct {
	imgs []*ebiten.Image
	gid  int
}

func NewLayer() *TileSetBuildings {
	var tileSetBuildings TileSetBuildings
	return &tileSetBuildings
}

func (t *TileSetBuildings) Img(id int) *ebiten.Image {
	id -= t.gid

	return t.imgs[id]
}

func (t *TileSetBuildings) GetImageObject(path string) error {
	// 1. Extract the image paths in the tileset json file
	data, err := fs.ReadFile(gameFs, path)
	if err != nil {
		return fmt.Errorf("Failed to read tileset file: %w", err)
	}

	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)

	var tileSetCollection TileSetCollection

	if err := decoder.Decode(&tileSetCollection); err != nil {
		return fmt.Errorf("Failed to parse json(TileSetCollection) at layer_interface.go: %w", err)
	}

	// 2. Create a slice of ebiten images
	t.imgs = make([]*ebiten.Image, 0)

	// 3. Loops the tileset image paths
	for _, tileset := range tileSetCollection.Tile {
		// 4. Prepare the image path string
		imagePath := utils.ParsePath1(tileset.ImgPath)
		imagePath = filepath.Join("assets/", imagePath)

		// 5. Create ebiten image with the image file path from tileset json file
		img, _, err := ebitenutil.NewImageFromFile(imagePath)
		if err != nil {
			return err
		}

		// 6. Save the ebiten image object
		t.imgs = append(t.imgs, img)
	}

	return nil
}
