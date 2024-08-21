package rcon

import (
	"fmt"
	"strings"

	"github.com/zMoooooritz/go-let-loose/pkg/config"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (r *Rcon) GetCurrentMap() (hll.Layer, error) {
	data, _ := r.runBasicCommand("get map")
	return hll.ParseLayer(data), nil
}

func (r *Rcon) GetAllMaps() ([]hll.Layer, error) {
	layerStr, err := r.runListCommand("get mapsforrotation")
	if err != nil {
		return []hll.Layer{}, err
	}
	return toLayers(layerStr), nil
}

func (r *Rcon) GetCurrentMapRotation() ([]hll.Layer, error) {
	resp, err := r.runBasicCommand("rotlist")
	if err != nil {
		return []hll.Layer{}, err
	}
	return toLayers(strings.Split(resp, config.NEWLINE)), nil
}

func (r *Rcon) AddMapToRotation(layer hll.Layer) error {
	return runSetCommand(r, fmt.Sprintf("rotadd %s", layer.ID))
}

func (r *Rcon) RemoveMapFromRotation(layer hll.Layer) error {
	return runSetCommand(r, fmt.Sprintf("rotdel %s", layer.ID))
}

func (r *Rcon) SetCurrentMap(layer hll.Layer) error {
	return runSetCommand(r, fmt.Sprintf("map %s", layer.ID))
}

func (r *Rcon) IsMapShuffleActive() (bool, error) {
	resp, err := r.runBasicCommand("querymapshuffle")
	if err != nil {
		return false, err
	}
	return strings.Contains(resp, "TRUE"), nil
}

func (r *Rcon) ToggleMapShuffle() error {
	_, err := r.runBasicCommand("togglemapshuffle")
	return err
}

func (r *Rcon) GetCurrentMapSequence() ([]hll.Layer, error) {
	resp, err := r.runBasicCommand("listcurrentmapsequence")
	if err != nil {
		return []hll.Layer{}, err
	}
	return toLayers(strings.Split(resp, config.NEWLINE)), nil
}

func toLayers(mapStrings []string) []hll.Layer {
	layers := []hll.Layer{}
	for _, entry := range mapStrings {
		layers = append(layers, hll.ParseLayer(entry))
	}
	return layers
}
