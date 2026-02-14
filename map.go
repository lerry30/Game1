package main

import (
	"fmt"
	"path"

	"game1/constants"
	"game1/layers"

	"github.com/hajimehoshi/ebiten/v2"
)

type Map struct {
	tileMap *TileMap
	layers  []layers.Layer
}

func NewMap(jsonPath string) (*Map, error) {
	// Create map data
	tileMap, err := NewTileMap(jsonPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to create tile map data at map.go: %w", err)
	}

	tileDataLayers := make([]layers.Layer, 0)
	for _, tileset := range tileMap.TileSet {
		tilesetPath := path.Join("assets/maps", tileset.Source)
		layer, err := layers.GetLayer(tilesetPath, tileset.GID)
		if err != nil {
			return nil, fmt.Errorf("Failed to get layers at map.go: %w", err)
		}

		tileDataLayers = append(tileDataLayers, layer)
	}

	return &Map{
		tileMap: tileMap,
		layers:  tileDataLayers,
	}, nil
}

func (m *Map) RenderMap(screen *ebiten.Image, drawOpts *DrawOption) {
	//opts.GeoM.Reset()
	for indexLayer, layer := range m.tileMap.Layers {
		for index, id := range layer.Data {
			if id == 0 {
				continue
			}

			x := index % layer.Width
			y := index / layer.Width

			x *= 16
			y *= 16

			img := m.layers[indexLayer].Img(id)

			drawOpts.Translate(float64(x), float64(y))
			drawOpts.opts.GeoM.Translate(0.0, -float64(img.Bounds().Dy()+16))
			screen.DrawImage(img, drawOpts.opts)

			drawOpts.opts.GeoM.Reset()
		}
	}
}

func (m *Map) BuildingsCoords() []*Rect {
	if len(m.tileMap.Layers) <= 1 {
		return nil
	}

	buildingsLayer := 1
	var buildingCoordData TileMapData = m.tileMap.Layers[buildingsLayer]
	coord := make([]*Rect, 0)
	yOffset := 60
	buildingBody := 20

	for index, id := range buildingCoordData.Data {
		if id == 0 {
			continue
		}

		x := (index % buildingCoordData.Width) * constants.TileSize
		y := (index / buildingCoordData.Width) * constants.TileSize

		img := m.layers[buildingsLayer].Img(id)
		buildingWidth := img.Bounds().Dx()
		buildingHeight := img.Bounds().Dy()

		coord = append(coord, &Rect{
			X:      x,
			Y:      y - yOffset,
			Width:  buildingWidth,
			Height: buildingHeight - buildingBody,
		})
	}

	return coord
}
