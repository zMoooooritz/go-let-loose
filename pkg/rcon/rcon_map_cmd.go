package rcon

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
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

func (r *Rcon) GetCurrentMapObjectives() ([][]string, error) {
	baseCommand := "get objectiverow_"
	objectiveNames := [][]string{}
	for i := range hll.ObjectiveCount[hll.GmWarfare] {
		resp, err := r.runListCommand(baseCommand + strconv.Itoa(i))
		if err != nil {
			return [][]string{}, err
		}
		objectiveNames = append(objectiveNames, resp)
	}
	return objectiveNames, nil
}

// 0    => random objective in that row
// 1-3  => specific objective
func (r *Rcon) SetGameLayoutIndexed(objs []int) error {
	if len(objs) != hll.ObjectiveCount[hll.GmWarfare] {
		return errors.New("Incorrect number of objectives provided")
	}

	for i := range hll.ObjectiveCount[hll.GmWarfare] {
		if objs[i] < 0 || objs[i] > hll.OptionsPerObjective[hll.GmWarfare] {
			return errors.New("Provided index is invalid (0 (random) and 1-3 are valid)")
		}
	}

	allObjNames, err := r.GetCurrentMapObjectives()
	if err != nil {
		return err
	}

	objNames := []string{}
	for i := range hll.ObjectiveCount[hll.GmWarfare] {
		objectiveIndex := objs[i]
		if objectiveIndex == 0 {
			objectiveIndex = rand.Intn(3) + 1
		}

		objNames = append(objNames, allObjNames[i][objectiveIndex-1])
	}
	return r.SetGameLayout(objNames)
}

func (r *Rcon) SetGameLayout(objs []string) error {
	if len(objs) != hll.ObjectiveCount[hll.GmWarfare] {
		return errors.New("Incorrect number of objectives provided")
	}
	return runSetCommand(r, fmt.Sprintf("gamelayout %s %s %s %s %s", objs[0], objs[1], objs[2], objs[3], objs[4]))
}

func toLayers(mapStrings []string) []hll.Layer {
	layers := []hll.Layer{}
	for _, entry := range mapStrings {
		layers = append(layers, hll.ParseLayer(entry))
	}
	return layers
}
