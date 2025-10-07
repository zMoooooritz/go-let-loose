package rcon

import (
	"fmt"
	"slices"
	"strings"

	"github.com/zMoooooritz/go-let-loose/pkg/hll"
	"github.com/zMoooooritz/go-let-loose/pkg/logger"
)

type rconVerification struct {
	enabled bool
}

func WithVerification() RconOption {
	return func(r *Rcon) {
		r.verification.enabled = true
	}
}

func (v *rconVerification) verifyLayers(r *Rcon) {
	if !v.enabled {
		return
	}

	resp, err := r.GetCommandDetails("AddMapToRotation")
	if err != nil {
		return
	}

	apiLayers := strings.Split(resp.DialogueParameters[0].ValueMember, ",")

	appLayers := []string{}
	for _, layer := range hll.AllLayers() {
		appLayers = append(appLayers, layer.ID)
	}

	appLayersNotInAPI := []string{}
	for _, layer := range appLayers {
		if !slices.Contains(apiLayers, layer) {
			appLayersNotInAPI = append(appLayersNotInAPI, layer)
		}
	}
	apiLayersNotInApp := []string{}
	for _, layer := range apiLayers {
		if !slices.Contains(appLayers, layer) {
			apiLayersNotInApp = append(apiLayersNotInApp, layer)
		}
	}

	if len(appLayersNotInAPI) > 0 {
		logger.Warn("There are layers in the application that are not present in the API:")
		logger.Warn("This may cause issues with the API, please report this to the developer.")
		printStr := fmt.Sprintf("Layers not in API (%d):", len(appLayersNotInAPI))
		logger.Warn(printStr)
		for _, layer := range appLayersNotInAPI {
			logger.Warn("\t- ", layer)
		}
	} else {
		logger.Debug("All layers in the application are present in the API.")

	}

	if len(apiLayersNotInApp) > 0 {
		logger.Warn("There are layers in the API that are not present in the application:")
		logger.Warn("This may cause issues with the API, please report this to the developer.")
		printStr := fmt.Sprintf("Layers not in the application (%d):", len(apiLayersNotInApp))
		logger.Warn(printStr)
		for _, layer := range apiLayersNotInApp {
			logger.Warn("\t- ", layer)
		}
	} else {
		logger.Debug("All layers in the API are present in the application.")
	}
}
