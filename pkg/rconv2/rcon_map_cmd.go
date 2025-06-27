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

func (r *Rcon) AddMapToSequence(layer hll.Layer, index int) error {
	_, err := runCommand[api.AddMapToSequence, any](r,
		api.AddMapToSequence{
			MapName: layer.ID,
			Index:   index,
		},
	)
	return err
}

func (r *Rcon) RemoveMapToSequence(index int) error {
	_, err := runCommand[api.RemoveMapFromSequence, any](r,
		api.RemoveMapFromSequence{
			Index: index,
		},
	)
	return err
}

func (r *Rcon) SetCurrentMap(layer hll.Layer) error {
	_, err := runCommand[api.ChangeMap, any](r,
		api.ChangeMap{
			MapName: layer.ID,
		},
	)
	return err
}

func (r *Rcon) ShuffleMapSequence(enabled bool) error {
	_, err := runCommand[api.SetShuffleMapSequence, any](r,
		api.SetShuffleMapSequence{
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
	// TODO: should be possible using the ClientReferenceData-API?
	return [][]string{}, fmt.Errorf("not implemented")
}

func (r *Rcon) SetGameLayoutIndexed(objs []int) error {
	// TODO: should be possible using the ClientReferenceData-API?
	return fmt.Errorf("not implemented")
}

func (r *Rcon) SetGameLayout(objs []string) error {
	// TODO: validation?
	_, err := runCommand[api.SetSectorLayout, any](r,
		api.SetSectorLayout{
			SectorOne:   objs[0],
			SectorTwo:   objs[1],
			SectorThree: objs[2],
			SectorFour:  objs[3],
			SectorFive:  objs[4],
		},
	)
	return err
}

func (r *Rcon) SetDynamicWeatherToggle(layer hll.Layer, enabled bool) error {
	_, err := runCommand[api.SetMapWeatherToggle, any](r,
		api.SetMapWeatherToggle{
			MapId:  layer.ID,
			Enable: enabled,
		},
	)
	return err
}

func (r *Rcon) SetMatchTimer(gameMode hll.GameMode, duration int) error {
	_, err := runCommand[api.SetMatchTimer, any](r,
		api.SetMatchTimer{
			GameMode:    string(gameMode),
			MatchLength: duration, // in minutes
		},
	)
	return err
}

func (r *Rcon) RemoveMatchTimer(gameMode hll.GameMode) error {
	_, err := runCommand[api.RemoveMatchTimer, any](r,
		api.RemoveMatchTimer{
			GameMode: string(gameMode),
		},
	)
	return err
}

func (r *Rcon) SetWarmupTimer(gameMode hll.GameMode, duration int) error {
	_, err := runCommand[api.SetWarmupTimer, any](r,
		api.SetWarmupTimer{
			GameMode:     string(gameMode),
			WarmupLength: duration, // in minutes
		},
	)
	return err
}

func (r *Rcon) RemoveWarmupTimer(gameMode hll.GameMode) error {
	_, err := runCommand[api.RemoveWarmupTimer, any](r,
		api.RemoveWarmupTimer{
			GameMode: string(gameMode),
		},
	)
	return err
}
