package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// It will just get the data from the json file to populate these structs

type TileMapData struct {
	Data   []int `json"data"`
	Width  int   `json:"width"`
	Height int   `json:"height"`
}

type TileSetData struct {
	GID    int    `json:"firstgid"`
	Source string `json:"source"`
}

type TileMap struct {
	Layers  []TileMapData `json:"layers"`
	TileSet []TileSetData `json:"tilesets"`
}

func NewTileMap(filepath string) (*TileMap, error) {
	var tileMap TileMap
	if err := tileMap.Load(filepath); err != nil {
		return &TileMap{}, fmt.Errorf("Failed to load json file at tilemap.go: %w", err)
	}

	return &tileMap, nil
}

func (tm *TileMap) Load(file string) error {
	data, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Reading file error: %s - %w", file, err)
	}

	reader := bytes.NewReader(data)
	decoder := json.NewDecoder(reader)

	if err := decoder.Decode(tm); err != nil {
		return fmt.Errorf("Failed to decode json content: %w", err)
	}

	return nil
}
