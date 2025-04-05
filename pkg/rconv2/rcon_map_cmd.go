package rconv2

import (
	"fmt"

	"github.com/zMoooooritz/go-let-loose/internal/socketv2/api"
	"github.com/zMoooooritz/go-let-loose/pkg/hll"
)

func (r *Rcon) GetCurrentMap() (hll.GameMap, error) {
	resp, err := getSessionInfo(r)
	if err != nil {
		return hll.GameMap{}, err
	}
	return hll.LogMapNameToMap(resp.MapName), nil
}

func (r *Rcon) GetCurrentLayer() (hll.Layer, error) {
	return hll.Layer{}, fmt.Errorf("not implemented")
}

func (r *Rcon) GetGameMode() (string, error) {
	resp, err := getSessionInfo(r)
	if err != nil {
		return "", err
	}
	return resp.GameMode, nil
}

func (r *Rcon) GetAllMaps() ([]hll.Layer, error) {
	return []hll.Layer{}, fmt.Errorf("not implemented")
}

func (r *Rcon) GetCurrentMapRotation() ([]hll.Layer, error) {
	layers := []hll.Layer{}
	resp, err := getMapRotation(r)
	if err != nil {
		return layers, err
	}
	for _, entry := range resp.Maps {
		layers = append(layers, hll.ParseLayer(entry.ID))
	}
	return layers, nil
}

func (r *Rcon) AddMapToRotation(layer hll.Layer, index int) error {
	_, err := runCommand[api.AddMapToRotation, any](r,
		api.AddMapToRotation{
			MapName: layer.ID,
			Index:   index,
		},
	)
	return err
}

func (r *Rcon) RemoveMapFromRotation(index int) error {
	_, err := runCommand[api.RemoveMapFromRotation, any](r,
		api.RemoveMapFromRotation{
			Index: index,
		},
	)
	return err
}

func (r *Rcon) SetCurrentMap(layer hll.Layer) error {
	_, err := runCommand[api.MapChange, any](r,
		api.MapChange{
			MapName: layer.ID,
		},
	)
	return err
}

func (r *Rcon) ShuffleMapSequence(enabled bool) error {
	_, err := runCommand[api.ShuffleMapSequence, any](r,
		api.ShuffleMapSequence{
			Enable: enabled,
		},
	)
	return err
}

func (r *Rcon) GetCurrentMapSequence() ([]hll.Layer, error) {
	layers := []hll.Layer{}
	resp, err := getMapSequence(r)
	if err != nil {
		return layers, err
	}
	for _, entry := range resp.Maps {
		layers = append(layers, hll.ParseLayer(entry.ID))
	}
	return layers, nil
}

func (r *Rcon) MoveMapInSequence(from, to int) error {
	_, err := runCommand[api.MoveMapInSequence, any](r,
		api.MoveMapInSequence{
			CurrentIndex: from,
			NewIndex:     to,
		},
	)
	return err
}

func (r *Rcon) GetCurrentMapObjectives() ([][]string, error) {
	return [][]string{}, fmt.Errorf("not implemented")
}

func (r *Rcon) SetGameLayoutIndexed(objs []int) error {
	return fmt.Errorf("not implemented")
}

func (r *Rcon) SetGameLayout(objs []string) error {
	// TODO: validation
	_, err := runCommand[api.ChangeSectorLayout, any](r,
		api.ChangeSectorLayout{
			SectorOne:   objs[0],
			SectorTwo:   objs[1],
			SectorThree: objs[2],
			SectorFour:  objs[3],
			SectorFive:  objs[4],
		},
	)
	return err
}
