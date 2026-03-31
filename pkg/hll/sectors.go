package hll

type SectorsIdentifier string

const (
	SECTORS_CARENTAN_LARGE        SectorsIdentifier = "carentan_large"
	SECTORS_CARENTAN_SMALL        SectorsIdentifier = "carentan_small"
	SECTORS_DRIEL_LARGE           SectorsIdentifier = "driel_large"
	SECTORS_DRIEL_SMALL           SectorsIdentifier = "driel_small"
	SECTORS_ELALAMEIN_LARGE       SectorsIdentifier = "elalamein_large"
	SECTORS_ELALAMEIN_SMALL       SectorsIdentifier = "elalamein_small"
	SECTORS_ELSENBORNRIDGE_LARGE  SectorsIdentifier = "elsenbornridge_large"
	SECTORS_ELSENBORNRIDGE_SMALL  SectorsIdentifier = "elsenbornridge_small"
	SECTORS_FOY_LARGE             SectorsIdentifier = "foy_large"
	SECTORS_HILL400_LARGE         SectorsIdentifier = "hill400_large"
	SECTORS_HILL400_SMALL         SectorsIdentifier = "hill400_small"
	SECTORS_HURTGENFOREST_LARGE   SectorsIdentifier = "hurtgenforest_large"
	SECTORS_KHARKOV_LARGE         SectorsIdentifier = "kharkov_large"
	SECTORS_KURSK_LARGE           SectorsIdentifier = "kursk_large"
	SECTORS_MORTAIN_LARGE         SectorsIdentifier = "mortain_large"
	SECTORS_MORTAIN_SMALL         SectorsIdentifier = "mortain_small"
	SECTORS_OMAHABEACH_LARGE      SectorsIdentifier = "omahabeach_large"
	SECTORS_PURPLEHEARTLANE_LARGE SectorsIdentifier = "purpleheartlane_large"
	SECTORS_PURPLEHEARTLANE_SMALL SectorsIdentifier = "purpleheartlane_small"
	SECTORS_REMAGEN_LARGE         SectorsIdentifier = "remagen_large"
	SECTORS_REMAGEN_SMALL         SectorsIdentifier = "remagen_small"
	SECTORS_SMOLENSK_LARGE        SectorsIdentifier = "smolensk_large"
	SECTORS_SMOLENSK_SMALL        SectorsIdentifier = "smolensk_small"
	SECTORS_STALINGRAD_LARGE      SectorsIdentifier = "stalingrad_large"
	SECTORS_STALINGRAD_SMALL      SectorsIdentifier = "stalingrad_small"
	SECTORS_STMARIEDUMONT_LARGE   SectorsIdentifier = "stmariedumont_large"
	SECTORS_STMARIEDUMONT_SMALL   SectorsIdentifier = "stmariedumont_small"
	SECTORS_STMEREEGLISE_LARGE    SectorsIdentifier = "stmereeglise_large"
	SECTORS_STMEREEGLISE_SMALL    SectorsIdentifier = "stmereeglise_small"
	SECTORS_TOBRUK_LARGE          SectorsIdentifier = "tobruk_large"
	SECTORS_TOBRUK_SMALL          SectorsIdentifier = "tobruk_small"
	SECTORS_UTAHBEACH_LARGE       SectorsIdentifier = "utahbeach_large"
)

func (s SectorsIdentifier) Sectors() []Sector {
	if sectors, ok := sectorsMap[s]; ok {
		return sectors
	}
	return []Sector{}
}

type Strongpoint struct {
	ID     string
	Name   string
	Center Position
	Radius float64
}

func (s Strongpoint) IsInside(pos Position) bool {
	dx := pos.X - s.Center.X
	dy := pos.Y - s.Center.Y
	dz := pos.Z - s.Center.Z
	distanceSquared := dx*dx + dy*dy + dz*dz
	return distanceSquared <= s.Radius*s.Radius
}

type CaptureZone struct {
	From        GridCoordinate
	To          GridCoordinate
	Strongpoint Strongpoint
}

type Sector struct {
	From         GridCoordinate
	To           GridCoordinate
	CaptureZones []CaptureZone
}

var sectorsMap = map[SectorsIdentifier][]Sector{
	SECTORS_CARENTAN_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "BLACTOT",
						Name:   "Blactot",
						Center: Position{X: -65543.41, Y: -39731.965, Z: 1359.9531},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "502ND START",
						Name:   "502nd Start",
						Center: Position{X: -67076.41, Y: 4670.035, Z: 123.953125},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "FARM RUINS",
						Name:   "Farm Ruins",
						Center: Position{X: -68814.41, Y: 37720.035, Z: 365.95312},
						Radius: 5000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "PUMPING STATION",
						Name:   "Pumping Station",
						Center: Position{X: -36748.406, Y: -29821.965, Z: 146.95312},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "RUINS",
						Name:   "Ruins",
						Center: Position{X: -26183.406, Y: 2343.0352, Z: 101.953125},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "DERAILED TRAIN",
						Name:   "Derailed Train",
						Center: Position{X: -39381.406, Y: 28975.035, Z: 279.95312},
						Radius: 5000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "CANAL CROSSING",
						Name:   "Canal Crossing",
						Center: Position{X: 5892.5938, Y: -39387.965, Z: 279.95312},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "TOWN CENTER",
						Name:   "Town Center",
						Center: Position{X: 1021.59375, Y: -1021.96484, Z: 104.953125},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "TRAIN STATION",
						Name:   "Train Station",
						Center: Position{X: 246.59375, Y: 27698.035, Z: 176.95312},
						Radius: 5000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "CUSTOMS",
						Name:   "Customs",
						Center: Position{X: 40816.594, Y: -34224.965, Z: 279.95312},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "RAIL CROSSING",
						Name:   "Rail Crossing",
						Center: Position{X: 44171.594, Y: -6296.965, Z: 279.95312},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "MONT HALAIS",
						Name:   "Mount Halais",
						Center: Position{X: 33828.594, Y: 51343.035, Z: 2518.9531},
						Radius: 3973.6313,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "CANAL LOCKS",
						Name:   "Canal Locks",
						Center: Position{X: 66826.59, Y: -26456.965, Z: 279.95312},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "RAIL CAUSEWAY",
						Name:   "Rail Causeway",
						Center: Position{X: 75611.59, Y: 5968.035, Z: 279.95312},
						Radius: 5495.63,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "LA MAISON DES ORMES",
						Name:   "La Maison Des Ormes",
						Center: Position{X: 72222.59, Y: 38476.035, Z: 103.53516},
						Radius: 5000.0,
					},
				},
			},
		},
	},
	SECTORS_CARENTAN_SMALL: {
		{
			From:         GridCoordinate{X: -5, Y: -4},
			To:           GridCoordinate{X: -2, Y: 3},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -1, Y: -4},
			To:   GridCoordinate{X: 0, Y: 3},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "Town Center",
						Name:   "Town Center",
						Center: Position{X: 510.0, Y: -170.0, Z: -850.0},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: 1, Y: -4},
			To:           GridCoordinate{X: 4, Y: 3},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_DRIEL_LARGE: {
		{
			From: GridCoordinate{X: -3, Y: -5},
			To:   GridCoordinate{X: 2, Y: -4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -5},
					To:   GridCoordinate{X: -2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "OOSTERBEEK APPROACH",
						Name:   "Oosterbeek Approach",
						Center: Position{X: -36028.543, Y: -79955.25, Z: -412.13928},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -5},
					To:   GridCoordinate{X: 0, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "ROSANDER POLDER",
						Name:   "Roseander Polder",
						Center: Position{X: 2809.0745, Y: -78795.875, Z: -159.22235},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -5},
					To:   GridCoordinate{X: 2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "KASTEEL ROSANDE",
						Name:   "Kasteel Rosande",
						Center: Position{X: 38371.14, Y: -76418.7, Z: 86.83746},
						Radius: 7000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: 2, Y: -2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "BOATYARD",
						Name:   "Boatyard",
						Center: Position{X: -38518.715, Y: -33980.625, Z: -205.24878},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "BRIDGEWAY",
						Name:   "Bridgeway",
						Center: Position{X: 3880.9673, Y: -39449.43, Z: -249.79703},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "RIJN BANKS",
						Name:   "Rijn Banks",
						Center: Position{X: 39177.535, Y: -42960.49, Z: -309.9341},
						Radius: 7000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -1},
			To:   GridCoordinate{X: 2, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "BRICK FACTORY",
						Name:   "Brick Factory",
						Center: Position{X: -39703.027, Y: 6122.7656, Z: -205.24878},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "RAILWAY BRIDGE",
						Name:   "Railway Bridge",
						Center: Position{X: 2882.3755, Y: -3877.2988, Z: -205.24878},
						Radius: 9000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "GUN EMPLACEMENTS",
						Name:   "Gun Emplacements",
						Center: Position{X: 43301.99, Y: -2530.0012, Z: -205.24878},
						Radius: 5500.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 1},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "RIETVELD",
						Name:   "Rietveld",
						Center: Position{X: -40615.844, Y: 40909.707, Z: -375.4043},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "SOUTH RAILWAY",
						Name:   "South Railway",
						Center: Position{X: 3826.8418, Y: 42206.754, Z: -429.63922},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "MIDDEL ROAD",
						Name:   "Middel Road",
						Center: Position{X: 41461.46, Y: 38457.824, Z: -205.24878},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 3},
			To:   GridCoordinate{X: 2, Y: 4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 3},
					To:   GridCoordinate{X: -2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "ORCHARDS",
						Name:   "Orchards",
						Center: Position{X: -39533.195, Y: 77266.98, Z: -329.62988},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 3},
					To:   GridCoordinate{X: 0, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "SCHADUWWOLKEN FARM",
						Name:   "Schaduwwolken Farm",
						Center: Position{X: -2113.1738, Y: 75816.06, Z: -357.01},
						Radius: 6500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 3},
					To:   GridCoordinate{X: 2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "FIELDS",
						Name:   "Fields",
						Center: Position{X: 41461.46, Y: 75453.516, Z: -205.25464},
						Radius: 6000.0,
					},
				},
			},
		},
	},
	SECTORS_DRIEL_SMALL: {
		{
			From:         GridCoordinate{X: -4, Y: -5},
			To:           GridCoordinate{X: 3, Y: -2},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -4, Y: -1},
			To:   GridCoordinate{X: 3, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "UNDERPASS",
						Name:   "Underpass",
						Center: Position{X: 2540.0, Y: 83820.0, Z: 450.0},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: -4, Y: 1},
			To:           GridCoordinate{X: 3, Y: 4},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_ELALAMEIN_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "VEHICLE DEPOT",
						Name:   "Vehicle Depot",
						Center: Position{X: -68233.38, Y: -37264.52, Z: 968.23914},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "ARTILLERY GUNS",
						Name:   "Artillery Guns",
						Center: Position{X: -71609.83, Y: -8175.4565, Z: -228.91907},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "MITEIRIYA RIDGE",
						Name:   "Miteiriya Ridge",
						Center: Position{X: -79261.695, Y: 36680.625, Z: 1629.1664},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "HAMLET RUINS",
						Name:   "Hamlet Ruins",
						Center: Position{X: -37466.633, Y: -37732.38, Z: -1402.2225},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "EL MREIR",
						Name:   "El Mreir",
						Center: Position{X: -37776.816, Y: -2887.5278, Z: -1248.4254},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "WATCHTOWER",
						Name:   "Watchtower",
						Center: Position{X: -40818.59, Y: 37838.586, Z: 648.23755},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "DESERT RAT TRENCHES",
						Name:   "Desert Rat Trenches",
						Center: Position{X: 4880.006, Y: -40988.05, Z: 831.8468},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "OASIS",
						Name:   "Oasis",
						Center: Position{X: -2900.921, Y: -851.27783, Z: -1248.4254},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "VALLEY",
						Name:   "Valley",
						Center: Position{X: 1970.4421, Y: 35186.074, Z: -787.2982},
						Radius: 8190.72,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "FUEL DEPOT",
						Name:   "Fuel Depot",
						Center: Position{X: 43333.848, Y: -35426.484, Z: -1862.4927},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "AIRFIELD COMMAND",
						Name:   "Airfield Command",
						Center: Position{X: 38495.92, Y: -4155.8906, Z: -1057.3627},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "AIRFIELD HANGARS",
						Name:   "Airfield Hangars",
						Center: Position{X: 41085.367, Y: 32927.33, Z: -1248.4176},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "CLIFFSIDE VILLAGE",
						Name:   "Cliffside Village",
						Center: Position{X: 68942.24, Y: -39028.402, Z: 550.42114},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "AMBUSHED CONVOY",
						Name:   "Ambushed Convoy",
						Center: Position{X: 72480.45, Y: -2526.4434, Z: -1248.4176},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "QUARRY",
						Name:   "Quarry",
						Center: Position{X: 78760.73, Y: 41540.4, Z: -69.86389},
						Radius: 6000.0,
					},
				},
			},
		},
	},
	SECTORS_ELALAMEIN_SMALL: {
		{
			From:         GridCoordinate{X: -5, Y: -4},
			To:           GridCoordinate{X: -2, Y: 3},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -1, Y: -4},
			To:   GridCoordinate{X: 0, Y: 3},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "OASIS",
						Name:   "Oasis",
						Center: Position{X: -15171.79688, Y: -4065.439, Z: -850.0},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: 1, Y: -4},
			To:           GridCoordinate{X: 4, Y: 3},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_ELSENBORNRIDGE_LARGE: {
		{
			From: GridCoordinate{X: -3, Y: -5},
			To:   GridCoordinate{X: 2, Y: -4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -5},
					To:   GridCoordinate{X: -2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "99TH COMMAND CENTRE",
						Name:   "99th Command Centre",
						Center: Position{X: -39637.496, Y: -67610.96, Z: 5383.203},
						Radius: 7500.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -5},
					To:   GridCoordinate{X: 0, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "GUN BATTERY",
						Name:   "Gun Battery",
						Center: Position{X: 420.0, Y: -69376.0, Z: 6308.203},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -5},
					To:   GridCoordinate{X: 2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "U.S. CAMP",
						Name:   "U.S. Camp",
						Center: Position{X: 50979.016, Y: -67675.0, Z: 6308.203},
						Radius: 7000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: 2, Y: -2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ELSENBORN RIDGE",
						Name:   "Elsenborn Ridge",
						Center: Position{X: -30950.0, Y: -41967.96, Z: 6858.203},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "FARAHILDE FARM",
						Name:   "Farahilde Farm",
						Center: Position{X: 10158.0, Y: -30210.0, Z: 5808.203},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "JENSIT PILLBOXES",
						Name:   "Jensit Pillboxes",
						Center: Position{X: 49674.99, Y: -28085.947, Z: 5108.2026},
						Radius: 7000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -1},
			To:   GridCoordinate{X: 2, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "ROAD TO ELSENBORN RIDGE",
						Name:   "Road To Elsenborn Ridge",
						Center: Position{X: -40964.0, Y: 3317.0, Z: 6608.203},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "DUGOUT TANKS",
						Name:   "Dug Out Tank",
						Center: Position{X: -9124.0, Y: 2404.0, Z: 5608.203},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CHECKPOINT",
						Name:   "Checkpoint",
						Center: Position{X: 40444.914, Y: 6529.1445, Z: 1716.1997},
						Radius: 5000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 1},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "ERELSDELL FARMHOUSE",
						Name:   "Erelsdell Farmhouse",
						Center: Position{X: -41672.0, Y: 38246.0, Z: 3633.2031},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "AA BATTERY",
						Name:   "AA Battery",
						Center: Position{X: 8607.678, Y: 33127.07, Z: 2796.5845},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "HINTERBERG",
						Name:   "Hinterburg",
						Center: Position{X: 39637.227, Y: 39888.688, Z: 3225.2002},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 3},
			To:   GridCoordinate{X: 2, Y: 4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 3},
					To:   GridCoordinate{X: -2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "SUPPLY CACHE",
						Name:   "Supply Cache",
						Center: Position{X: -25666.0, Y: 66300.0, Z: 2862.2031},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 3},
					To:   GridCoordinate{X: 0, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "FOXHOLES",
						Name:   "Foxholes",
						Center: Position{X: 12223.855, Y: 67172.27, Z: -364.80078},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 3},
					To:   GridCoordinate{X: 2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "FUEL DEPOT",
						Name:   "Fuel Depot",
						Center: Position{X: 38049.76, Y: 70408.04, Z: 2951.984},
						Radius: 7000.0,
					},
				},
			},
		},
	},
	SECTORS_ELSENBORNRIDGE_SMALL: {
		{
			From:         GridCoordinate{X: -4, Y: -5},
			To:           GridCoordinate{X: 3, Y: -2},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -4, Y: -1},
			To:   GridCoordinate{X: 3, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "DUG OUT TANK",
						Name:   "Dug Out Tank",
						Center: Position{X: -8510.0, Y: 1410.0, Z: -279.18115},
						Radius: 5000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: -4, Y: 1},
			To:           GridCoordinate{X: 3, Y: 4},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_FOY_LARGE: {
		{
			From: GridCoordinate{X: -3, Y: -5},
			To:   GridCoordinate{X: 2, Y: -4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -5},
					To:   GridCoordinate{X: -2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "ROAD TO RECOGNE",
						Name:   "Road To Recogne",
						Center: Position{X: -49755.0, Y: -74340.0, Z: -211.0},
						Radius: 2750.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -5},
					To:   GridCoordinate{X: 0, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "COBRU APPROACH",
						Name:   "Cobru Approach",
						Center: Position{X: 9952.0, Y: -74787.0, Z: -243.0},
						Radius: 3500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -5},
					To:   GridCoordinate{X: 2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "ROAD TO NOVILLE",
						Name:   "Road To Noville",
						Center: Position{X: 38286.176, Y: -76947.95, Z: -243.0},
						Radius: 5343.75,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: 2, Y: -2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "COBRU FACTORY",
						Name:   "Cobru Factory",
						Center: Position{X: -29988.0, Y: -44676.0, Z: -890.0},
						Radius: 5500.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "FOY",
						Name:   "Foy",
						Center: Position{X: -9586.0, Y: -34052.0, Z: -551.0},
						Radius: 3250.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "FLAK BATTERY",
						Name:   "Flak Battery",
						Center: Position{X: 45241.0, Y: -39594.0, Z: -964.0},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -1},
			To:   GridCoordinate{X: 2, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "WEST BEND",
						Name:   "West Bend",
						Center: Position{X: -53153.0, Y: -12966.0, Z: -634.0},
						Radius: 5500.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "SOUTHERN EDGE",
						Name:   "Southern Edge",
						Center: Position{X: -1114.0, Y: 589.0, Z: -102.0},
						Radius: 4738.37,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "DUGOUT BARN",
						Name:   "Dugout Barn",
						Center: Position{X: 46085.04, Y: -4721.094, Z: -1008.08936},
						Radius: 4139.884,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 1},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "N30 HIGHWAY",
						Name:   "N30 Highway",
						Center: Position{X: -38407.0, Y: 31775.0, Z: -142.0},
						Radius: 6250.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "BIZORY-FOY ROAD",
						Name:   "Bizory-Foy Road",
						Center: Position{X: 10035.0, Y: 39390.0, Z: -545.0},
						Radius: 3500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "EASTERN OURTHE",
						Name:   "Eastern Ourthe",
						Center: Position{X: 45845.0, Y: 27822.0, Z: -771.0},
						Radius: 4531.25,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 3},
			To:   GridCoordinate{X: 2, Y: 4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 3},
					To:   GridCoordinate{X: -2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "ROAD TO BASTOGNE",
						Name:   "Road To Bastogne",
						Center: Position{X: -52862.0, Y: 63773.0, Z: 112.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 3},
					To:   GridCoordinate{X: 0, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "BOIS JACQUES",
						Name:   "Bois Jacques",
						Center: Position{X: -5582.0, Y: 68237.0, Z: 1106.0},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 3},
					To:   GridCoordinate{X: 2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "FOREST OUTSKIRTS",
						Name:   "Forest Outskirts",
						Center: Position{X: 46279.0, Y: 67141.0, Z: 512.0},
						Radius: 5000.0,
					},
				},
			},
		},
	},
	SECTORS_HILL400_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "CONVOY AMBUSH",
						Name:   "Convoy Ambush",
						Center: Position{X: -65875.18, Y: -36966.816, Z: 6207.824},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "FEDERHECKE JUNCTION",
						Name:   "Federchecke Junction",
						Center: Position{X: -65367.926, Y: 2874.167, Z: 10826.176},
						Radius: 4250.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "STUCKCHEN FARM",
						Name:   "Stuckchen Farm",
						Center: Position{X: -63938.484, Y: 42413.004, Z: 8142.296},
						Radius: 3000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ROER RIVER HOUSE",
						Name:   "Roer River House",
						Center: Position{X: -38405.066, Y: -43380.766, Z: -342.3706},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "BERGSTEIN CHURCH",
						Name:   "Bergstein Church",
						Center: Position{X: -30580.357, Y: 8420.501, Z: 11575.116},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "KIRCHWEG",
						Name:   "Kirchweg",
						Center: Position{X: -41257.29, Y: 31282.14, Z: 8949.19},
						Radius: 3000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "FLAK PITS",
						Name:   "Flak Pits",
						Center: Position{X: 1384.4886, Y: -33584.805, Z: 8937.715},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "HILL 400",
						Name:   "Hill 400",
						Center: Position{X: -1408.8995, Y: 4698.0444, Z: 17213.738},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "SOUTHERN APPROACH",
						Name:   "Southern Approach",
						Center: Position{X: 948.21277, Y: 25170.994, Z: 12086.199},
						Radius: 3000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ESELSWEG JUNCTION",
						Name:   "Eselsweg Junction",
						Center: Position{X: 26549.63, Y: -41028.504, Z: 7713.7764},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "EASTERN SLOPE",
						Name:   "Eastern Slope",
						Center: Position{X: 29662.375, Y: -3406.8445, Z: 8725.453},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "TRAIN WRECK",
						Name:   "Trainwreck",
						Center: Position{X: 32129.537, Y: 43600.098, Z: 994.9547},
						Radius: 3000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ROER RIVER CROSSING",
						Name:   "Roer River Crossing",
						Center: Position{X: 64685.836, Y: -33321.977, Z: -2164.823},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "ZERKALL",
						Name:   "Zerkall",
						Center: Position{X: 78823.555, Y: -9569.677, Z: -2095.891},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "PAPER MILL",
						Name:   "Paper Mill",
						Center: Position{X: 69319.79, Y: 39032.61, Z: -2095.891},
						Radius: 3000.0,
					},
				},
			},
		},
	},
	SECTORS_HILL400_SMALL: {
		{
			From:         GridCoordinate{X: -5, Y: -4},
			To:           GridCoordinate{X: -2, Y: 3},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -1, Y: -4},
			To:   GridCoordinate{X: 0, Y: 3},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "HILL 400",
						Name:   "Hill 400",
						Center: Position{X: 0.0, Y: 1015.0, Z: -519.0},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: 1, Y: -4},
			To:           GridCoordinate{X: 4, Y: 3},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_HURTGENFOREST_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "MAUSBACH APPROACH",
						Name:   "The Masbauch Approach",
						Center: Position{X: -74423.0, Y: -46733.0, Z: 4336.0},
						Radius: 4625.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "RESERVE STATION",
						Name:   "Reserve Station",
						Center: Position{X: -78776.0, Y: 2238.0, Z: 3895.0},
						Radius: 4500.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "LUMBER YARD",
						Name:   "Lumber Yard",
						Center: Position{X: -77356.0, Y: 36029.0, Z: 4122.0},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "WEHEBACH OVERLOOK",
						Name:   "Wehebach Overlook",
						Center: Position{X: -38278.0, Y: -34416.0, Z: 6683.0},
						Radius: 5031.25,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "KALL TRAIL",
						Name:   "Kall Trail",
						Center: Position{X: -35755.0, Y: 2459.0, Z: 4007.0},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "THE RUIN",
						Name:   "The Ruin",
						Center: Position{X: -42793.0, Y: 26141.0, Z: 3879.0},
						Radius: 4400.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "NORTH PASS",
						Name:   "North Pass",
						Center: Position{X: 6540.0, Y: -49329.0, Z: 215.0},
						Radius: 4347.4674749999995,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "THE SCAR",
						Name:   "The Scar",
						Center: Position{X: -6935.0, Y: 3328.0, Z: 1327.0},
						Radius: 3015.2502,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "THE SIEGFRIED LINE",
						Name:   "The Siegfried Line",
						Center: Position{X: -3711.0, Y: 42305.0, Z: 2603.0},
						Radius: 4500.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "HILL 15",
						Name:   "Hill 15",
						Center: Position{X: 45628.0, Y: -34330.0, Z: 4504.0},
						Radius: -4500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "JACOB'S BARN",
						Name:   "Jacob's Barn",
						Center: Position{X: 37658.0, Y: 8531.0, Z: 6550.0},
						Radius: 3500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "SALIENT 42",
						Name:   "Salient 42",
						Center: Position{X: 40632.0, Y: 50244.0, Z: 6895.0},
						Radius: 3250.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "GROSSHAU APPROACH",
						Name:   "Grosshau Approach",
						Center: Position{X: 73663.0, Y: -38895.0, Z: 5297.0},
						Radius: 3750.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "HURTGEN APPROACH",
						Name:   "Hürtgen Approach",
						Center: Position{X: 67776.0, Y: 6558.0, Z: 6600.0},
						Radius: 3500.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "LOGGING CAMP",
						Name:   "Logging Camp",
						Center: Position{X: 64477.0, Y: 51502.0, Z: 6495.0},
						Radius: 3750.0,
					},
				},
			},
		},
	},
	SECTORS_KHARKOV_LARGE: {
		{
			From: GridCoordinate{X: -3, Y: -5},
			To:   GridCoordinate{X: 2, Y: -4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -5},
					To:   GridCoordinate{X: -2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "MARSH TOWN",
						Name:   "Marsh Town",
						Center: Position{X: -36517.52, Y: -70661.75, Z: -2300.2422},
						Radius: 3750.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -5},
					To:   GridCoordinate{X: 0, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "SOVIET VANTAGE POINT",
						Name:   "Soviet Vantage Point",
						Center: Position{X: 8032.91, Y: -70714.63, Z: 402.14062},
						Radius: 3750.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -5},
					To:   GridCoordinate{X: 2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "GERMAN FUEL DUMP",
						Name:   "German Fuel Dump",
						Center: Position{X: 41168.312, Y: -70231.15, Z: 3192.4492},
						Radius: 3750.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: 2, Y: -2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "BITTER SPRING",
						Name:   "Bitter Spring",
						Center: Position{X: -37433.285, Y: -38891.406, Z: -2293.7441},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "LUMBER WORKS",
						Name:   "Lumber Works",
						Center: Position{X: 7916.1367, Y: -39814.156, Z: 279.73047},
						Radius: 3750.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "WINDMILL HILLSIDE",
						Name:   "Windmill Hillside",
						Center: Position{X: 46877.23, Y: -41370.87, Z: 2556.1875},
						Radius: 3750.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -1},
			To:   GridCoordinate{X: 2, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "WATER MILL",
						Name:   "Water Mill",
						Center: Position{X: -36761.05, Y: -3563.8867, Z: -2120.2441},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "ST MARY",
						Name:   "St Mary",
						Center: Position{X: 6074.9873, Y: -633.23, Z: 911.6289},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "DISTILLERY",
						Name:   "Distillery",
						Center: Position{X: 44449.215, Y: -4542.487, Z: 2724.1992},
						Radius: 3750.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 1},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "RIVER CROSSING",
						Name:   "River Crossing",
						Center: Position{X: -27116.355, Y: 40003.023, Z: -890.95703},
						Radius: 3750.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "BELGOROD OUTSKIRTS",
						Name:   "Belgorod Outskirts",
						Center: Position{X: 8105.9688, Y: 38673.008, Z: 221.38281},
						Radius: 9000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "LUMBERYARD",
						Name:   "Lumberyard",
						Center: Position{X: 46774.79, Y: 37490.91, Z: 2052.8438},
						Radius: 3750.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 3},
			To:   GridCoordinate{X: 2, Y: 4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 3},
					To:   GridCoordinate{X: -2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "WEHRMACHT OUTLOOK",
						Name:   "Wehrmacht Overlook",
						Center: Position{X: -37313.91, Y: 72972.37, Z: -661.0508},
						Radius: 3750.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 3},
					To:   GridCoordinate{X: 0, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "HAY STORAGE",
						Name:   "Hay Storage",
						Center: Position{X: 4240.6523, Y: 71736.38, Z: -1926.957},
						Radius: 3750.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 3},
					To:   GridCoordinate{X: 2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "OVERPASS",
						Name:   "Overpass",
						Center: Position{X: 41180.39, Y: 70416.95, Z: -41.4375},
						Radius: 3750.0,
					},
				},
			},
		},
	},
	SECTORS_KURSK_LARGE: {
		{
			From: GridCoordinate{X: -3, Y: -5},
			To:   GridCoordinate{X: 2, Y: -4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -5},
					To:   GridCoordinate{X: -2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "ARTILLERY POSITION",
						Name:   "Artillery Position",
						Center: Position{X: -35117.0, Y: -68921.0, Z: 9323.0},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -5},
					To:   GridCoordinate{X: 0, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "GRUSHKI",
						Name:   "Grushki",
						Center: Position{X: 7070.0, Y: -68141.0, Z: 7093.0},
						Radius: 4960.5308,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -5},
					To:   GridCoordinate{X: 2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "GRUSHKI FLANK",
						Name:   "Grushki Flank",
						Center: Position{X: 47151.0, Y: -67169.0, Z: 5786.0},
						Radius: 4500.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: 2, Y: -2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "PANZER'S END",
						Name:   "Panzer's End",
						Center: Position{X: -35117.0, Y: -31958.0, Z: 8935.0},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "DEFENCE IN DEPTH",
						Name:   "Defence In Depth",
						Center: Position{X: 1604.0, Y: -34906.0, Z: 7647.0},
						Radius: 7022.2216,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "LISTENING POST",
						Name:   "Listening Post",
						Center: Position{X: 40413.0, Y: -36000.0, Z: 5889.0},
						Radius: 7673.426,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -1},
			To:   GridCoordinate{X: 2, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "THE WINDMILLS",
						Name:   "The Windmills",
						Center: Position{X: -26712.39, Y: -4842.251, Z: 9948.998},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "YAMKI",
						Name:   "Yamki",
						Center: Position{X: 9609.0, Y: 3974.0, Z: 8754.0},
						Radius: 6973.061,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "OLEG'S HOUSE",
						Name:   "Oleg's House",
						Center: Position{X: 39754.0, Y: 7774.0, Z: 6623.0},
						Radius: 4500.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 1},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "RUDNO",
						Name:   "Rudno",
						Center: Position{X: -27089.0, Y: 40069.0, Z: 9949.0},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "DESTROYED BATTERY",
						Name:   "Destroyed Battery",
						Center: Position{X: -990.0, Y: 39981.0, Z: 10190.0},
						Radius: 4500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "THE MUDDY CHURN",
						Name:   "The Muddy Churn",
						Center: Position{X: 41089.0, Y: 42772.0, Z: 7983.0},
						Radius: 4500.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 3},
			To:   GridCoordinate{X: 2, Y: 4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 3},
					To:   GridCoordinate{X: -2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "ROAD TO KURSK",
						Name:   "Road To Kursk",
						Center: Position{X: -31287.0, Y: 68120.0, Z: 9949.0},
						Radius: 4500.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 3},
					To:   GridCoordinate{X: 0, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "AMMO DUMP",
						Name:   "Ammo Dump",
						Center: Position{X: -1729.0, Y: 66294.0, Z: 9632.0},
						Radius: 5446.865,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 3},
					To:   GridCoordinate{X: 2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "EASTERN POSITION",
						Name:   "Eastern Position",
						Center: Position{X: 36100.0, Y: 65758.0, Z: 8227.0},
						Radius: 6000.0,
					},
				},
			},
		},
	},
	SECTORS_MORTAIN_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "HOTEL DE LA POSTE",
						Name:   "Hotel De La Poste",
						Center: Position{X: -71664.03, Y: -47217.445, Z: -480.30115},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "FORWARD BATTERY",
						Name:   "Forward Battery",
						Center: Position{X: -67949.4, Y: 6438.873, Z: -978.4287},
						Radius: 6500.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "SOUTHERN APPROACH",
						Name:   "Southern Approach",
						Center: Position{X: -70344.09, Y: 46402.33, Z: -5391.659},
						Radius: 7500.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "MORTAIN OUTSKIRTS",
						Name:   "Mortain Outskirts",
						Center: Position{X: -49136.977, Y: -39819.566, Z: 1842.5009},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "FORWARD MEDICAL AID STATION",
						Name:   "Forward Medical Aid Station",
						Center: Position{X: -35275.574, Y: -2194.1567, Z: 2411.5898},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "MORTAIN APPROACH",
						Name:   "Mortain Approach",
						Center: Position{X: -42775.5, Y: 33050.027, Z: -1580.7439},
						Radius: 7000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "HILL 314",
						Name:   "Hill 314",
						Center: Position{X: -2425.1055, Y: -38259.5, Z: 5891.1714},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "LA PETITE CHAPELLE SAINT-MICHEL",
						Name:   "La Petite Chapelle Saint-Michel",
						Center: Position{X: 1725.2772, Y: 5918.17, Z: 5670.377},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "U.S. SOUTHERN ROADBLOCK",
						Name:   "U.S. Southern Roadblock",
						Center: Position{X: -11254.05, Y: 49076.438, Z: -3064.87},
						Radius: 7000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "DESTROYED GERMAN CONVOY",
						Name:   "Destroyed German Convoy",
						Center: Position{X: 35469.836, Y: -42255.99, Z: 6050.5566},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "GERMAN RECON CAMP",
						Name:   "German Recon Camp",
						Center: Position{X: 40439.145, Y: -2510.7285, Z: 4398.437},
						Radius: 6000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "LES AUBRILS FARM",
						Name:   "Les Aubrils Farm",
						Center: Position{X: 48018.547, Y: 26619.574, Z: 1071.811},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ABANDONED GERMAN CHECKPOINT",
						Name:   "Abandoned German Checkpoint",
						Center: Position{X: 68651.26, Y: -40271.47, Z: 6068.428},
						Radius: 6500.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "GERMAN DEFENSIVE CAMP",
						Name:   "German Defensive Camp",
						Center: Position{X: 68294.46, Y: 1986.8845, Z: 3941.4448},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "LE FERME DU DESCHAMPS",
						Name:   "Le Ferme Du Deschamps",
						Center: Position{X: 71327.67, Y: 36841.695, Z: 1896.9764},
						Radius: 7000.0,
					},
				},
			},
		},
	},
	SECTORS_MORTAIN_SMALL: {
		{
			From:         GridCoordinate{X: -5, Y: -4},
			To:           GridCoordinate{X: -2, Y: 3},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -1, Y: -4},
			To:   GridCoordinate{X: 0, Y: 3},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "LA PETITE CHAPELLE SAINT-MICHEL",
						Name:   "La Petite Chapelle Saint-Michel",
						Center: Position{X: 1300.0, Y: 5500.0, Z: -1000.0},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: 1, Y: -4},
			To:           GridCoordinate{X: 4, Y: 3},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_OMAHABEACH_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "BEAUMONT ROAD",
						Name:   "Beaumont Road",
						Center: Position{X: -66508.0, Y: -34528.0, Z: 1461.6543},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CROSSROADS",
						Name:   "Crossroads",
						Center: Position{X: -63975.723, Y: 2684.23, Z: 1607.7931},
						Radius: 3713.8135,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "LES ISLES",
						Name:   "Les Isles",
						Center: Position{X: -65785.0, Y: 33673.0, Z: 1949.5156},
						Radius: 5000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "REAR BATTERY",
						Name:   "Rear Battery",
						Center: Position{X: -40364.508, Y: -47019.88, Z: 1471.3027},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CHURCH ROAD",
						Name:   "Church Road",
						Center: Position{X: -36692.0, Y: -9308.0, Z: 1471.3047},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "THE ORCHARDS",
						Name:   "The Orchards",
						Center: Position{X: -44319.355, Y: 27163.912, Z: 1705.2588},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "WEST VIERVILLE",
						Name:   "West Vierville",
						Center: Position{X: 4665.0, Y: -40540.0, Z: 1262.0},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "VIERVILLE SUR MER",
						Name:   "Vierville Sur Mer",
						Center: Position{X: -2661.8896, Y: -2895.0942, Z: 889.7754},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "ARTILLERY BATTERY",
						Name:   "Artillery Battery",
						Center: Position{X: 2342.277, Y: 31510.633, Z: 1730.7812},
						Radius: 5000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "WN73",
						Name:   "WN73",
						Center: Position{X: 54259.0, Y: -44498.0, Z: 164.8125},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "WN71",
						Name:   "WN71",
						Center: Position{X: 55132.387, Y: -5791.973, Z: 1015.4961},
						Radius: 3750.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "WN70",
						Name:   "WN70",
						Center: Position{X: 46516.0, Y: 30340.0, Z: 1368.2734},
						Radius: 5000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "DOG GREEN",
						Name:   "Dog Green",
						Center: Position{X: 67602.0, Y: -31262.0, Z: -3134.6387},
						Radius: 6250.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "THE DRAW",
						Name:   "The Draw",
						Center: Position{X: 71322.0, Y: -7432.0, Z: -2748.504},
						Radius: 3750.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "DOG WHITE",
						Name:   "Dog White",
						Center: Position{X: 71817.0, Y: 30284.0, Z: -3019.504},
						Radius: 5000.0,
					},
				},
			},
		},
	},
	SECTORS_PURPLEHEARTLANE_LARGE: {
		{
			From: GridCoordinate{X: -3, Y: -5},
			To:   GridCoordinate{X: 2, Y: -4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -5},
					To:   GridCoordinate{X: -2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "BLOODY BEND",
						Name:   "Bloody Bend",
						Center: Position{X: -53699.133, Y: -68803.984, Z: 6831.375},
						Radius: 2750.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -5},
					To:   GridCoordinate{X: 0, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "DEAD MAN'S CORNER",
						Name:   "Dead Man's Corner",
						Center: Position{X: 740.8672, Y: -65433.984, Z: 6831.375},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -5},
					To:   GridCoordinate{X: 2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "FORWARD BATTERY",
						Name:   "Forward Battery",
						Center: Position{X: 33330.867, Y: -66643.984, Z: 6831.375},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: 2, Y: -2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "JOURDAN CANAL",
						Name:   "Jourdan Canal",
						Center: Position{X: -41489.133, Y: -38108.99, Z: 6831.375},
						Radius: 2750.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "DOUVE BRIDGE",
						Name:   "Douve Bridge",
						Center: Position{X: -1434.0474, Y: -26826.268, Z: 7010.586},
						Radius: 4250.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "DOUVE RIVER BATTERY",
						Name:   "Douve River Battery",
						Center: Position{X: 33572.38, Y: -36601.72, Z: 6840.3447},
						Radius: 3500.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -1},
			To:   GridCoordinate{X: 2, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "GROULT PILLBOX",
						Name:   "Groult Pillbox",
						Center: Position{X: -37607.4, Y: -5672.991, Z: 6730.6123},
						Radius: 5500.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CARENTAN CAUSEWAY",
						Name:   "Carentan Causeway",
						Center: Position{X: 787.74744, Y: 1346.289, Z: 6969.521},
						Radius: 3500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "FLAK POSITION",
						Name:   "Flak Position",
						Center: Position{X: 45592.906, Y: -4116.6772, Z: 6732.951},
						Radius: -4750.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 1},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "MADELEINE FARM",
						Name:   "Madeleine Farm",
						Center: Position{X: -33264.676, Y: 30204.594, Z: 7391.552},
						Radius: 3250.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "MADELEINE BRIDGE",
						Name:   "Madeleine Bridge",
						Center: Position{X: 1928.2188, Y: 39878.098, Z: 6973.26},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "AID STATION",
						Name:   "Aid Station",
						Center: Position{X: 47043.207, Y: 32172.8, Z: 6753.122},
						Radius: 3250.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 3},
			To:   GridCoordinate{X: 2, Y: 4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 3},
					To:   GridCoordinate{X: -2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "INGOUF CROSSROADS",
						Name:   "Ingouf Crossroads",
						Center: Position{X: -36344.676, Y: 66489.59, Z: 7391.552},
						Radius: 3250.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 3},
					To:   GridCoordinate{X: 0, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "ROAD TO CARENTAN",
						Name:   "Road To Carentan",
						Center: Position{X: 2953.2188, Y: 63908.098, Z: 6973.26},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 3},
					To:   GridCoordinate{X: 2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "CABBAGE PATCH",
						Name:   "Cabbage Patch",
						Center: Position{X: 46253.22, Y: 62363.098, Z: 6973.26},
						Radius: 2500.0,
					},
				},
			},
		},
	},
	SECTORS_PURPLEHEARTLANE_SMALL: {
		{
			From:         GridCoordinate{X: -4, Y: -5},
			To:           GridCoordinate{X: 3, Y: -2},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -4, Y: -1},
			To:   GridCoordinate{X: 3, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CARENTAN CAUSEWAY",
						Name:   "Carentan Causeway",
						Center: Position{X: -960.0, Y: 285.0, Z: 46.0},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: -4, Y: 1},
			To:           GridCoordinate{X: 3, Y: 4},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_REMAGEN_LARGE: {
		{
			From: GridCoordinate{X: -3, Y: -5},
			To:   GridCoordinate{X: 2, Y: -4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -5},
					To:   GridCoordinate{X: -2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "ALTE LIEBE BARSCH",
						Name:   "Alte Liebe Barsch",
						Center: Position{X: -41114.0, Y: -69583.0, Z: 6515.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -5},
					To:   GridCoordinate{X: 0, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "BEWALDET KREUZUNG",
						Name:   "Bewaldet Kreuzung",
						Center: Position{X: -891.0, Y: -69550.0, Z: 12708.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -5},
					To:   GridCoordinate{X: 2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "DAN RADART 512",
						Name:   "Dan Radart 512",
						Center: Position{X: 41625.0, Y: -69063.0, Z: 16150.0},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: 2, Y: -2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ERPEL",
						Name:   "Erpel",
						Center: Position{X: -39275.0, Y: -40853.0, Z: 1774.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ERPELER LEY",
						Name:   "Erpeler Ley",
						Center: Position{X: 9697.0, Y: -42679.0, Z: 13960.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "KASBACH OUTLOOK",
						Name:   "Kasbach Outlook",
						Center: Position{X: 38436.418, Y: -41098.23, Z: 9033.0},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -1},
			To:   GridCoordinate{X: 2, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "ST. SEVERIN CHAPEL",
						Name:   "St Severin Chapel",
						Center: Position{X: -39275.0, Y: -12967.0, Z: 766.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "LUDENDORFF BRIDGE",
						Name:   "Ludendorff Bridge",
						Center: Position{X: 3032.2412, Y: 7.0210953, Z: 1261.005},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "BAUERNHOF AM RHEIN",
						Name:   "Bauernhof Am Rhein",
						Center: Position{X: 38817.02, Y: 15613.944, Z: 104.0},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 1},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "REMAGEN",
						Name:   "Remagen",
						Center: Position{X: -35925.75, Y: 39434.0, Z: -27.363525},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "MÖBELFABRIK",
						Name:   "Möbelfabrik",
						Center: Position{X: -1000.0, Y: 40824.0, Z: -35.674072},
						Radius: 5000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "SCHLIEFFEN AUSWEG",
						Name:   "Schlieffen Ausweg",
						Center: Position{X: 39053.0, Y: 38264.0, Z: 104.0},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 3},
			To:   GridCoordinate{X: 2, Y: 4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 3},
					To:   GridCoordinate{X: -2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "WALDBURG",
						Name:   "Waldburg",
						Center: Position{X: -40954.977, Y: 80279.71, Z: -125.40869},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 3},
					To:   GridCoordinate{X: 0, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "MÜHLENWEG",
						Name:   "Mühlenweg",
						Center: Position{X: 3742.6152, Y: 72094.91, Z: -121.5036},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 3},
					To:   GridCoordinate{X: 2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "HAGELKREUZ",
						Name:   "Hagelkreuz",
						Center: Position{X: 37607.746, Y: 68933.32, Z: 104.0},
						Radius: 4000.0,
					},
				},
			},
		},
	},
	SECTORS_REMAGEN_SMALL: {
		{
			From:         GridCoordinate{X: -4, Y: -5},
			To:           GridCoordinate{X: 3, Y: -2},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -4, Y: -1},
			To:   GridCoordinate{X: 3, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "LUDENDORFF BRIDGE",
						Name:   "LUDENDORFF BRIDGE",
						Center: Position{X: 3228.9722, Y: -570.1361, Z: -315.0},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: -4, Y: 1},
			To:           GridCoordinate{X: 3, Y: 4},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_SMOLENSK_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "PANZER LOADING STATION",
						Name:   "Panzer Loading Station",
						Center: Position{X: -68850.08, Y: -40044.953, Z: 512.03125},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "TRAM DEPOT",
						Name:   "Tram Depot",
						Center: Position{X: -67850.08, Y: -5084.953, Z: 512.03125},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "SMOLENSK OUTSKIRTS",
						Name:   "Smolensk Outskirts",
						Center: Position{X: -68850.08, Y: 40680.047, Z: 512.03125},
						Radius: 6500.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "SMOLENSK HAUPTBAHNHOF",
						Name:   "Smolensk Hauptbahnhof",
						Center: Position{X: -38450.08, Y: -40044.953, Z: 512.03125},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "LUMBER YARD",
						Name:   "Lumber Yard",
						Center: Position{X: -36050.08, Y: 8915.047, Z: 512.03125},
						Radius: 9000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "DNIEPER WEST CROSSING",
						Name:   "Dnieper West Crossing",
						Center: Position{X: -40450.08, Y: 39680.047, Z: 512.03125},
						Radius: 7000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "PYATNITSKII OVERPASS",
						Name:   "Pyatnitskii Overpass",
						Center: Position{X: 1000.0, Y: -38544.953, Z: 512.03125},
						Radius: 6500.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "ZHELYABOVA SQUARE",
						Name:   "Zhelyabova Square",
						Center: Position{X: 1000.0, Y: 0.0, Z: 0.0},
						Radius: 6500.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "84TH BATTALION BRIDGE",
						Name:   "84th Battalion Bridge",
						Center: Position{X: 1000.0, Y: 40680.047, Z: 512.03125},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ZADNEPROVIE DISTRICT",
						Name:   "Zadneprovie District",
						Center: Position{X: 39264.92, Y: -40444.953, Z: 512.03125},
						Radius: 6500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "MOSKOVSKAYA STREET",
						Name:   "Moskovskaya Street",
						Center: Position{X: 39764.92, Y: -1084.9531, Z: 512.03125},
						Radius: 6500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "SMOLENSK CITADEL",
						Name:   "Smolensk Citadel",
						Center: Position{X: 39264.92, Y: 40680.047, Z: 512.03125},
						Radius: 7000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "RAILYARD STORAGE",
						Name:   "Railyard Storage",
						Center: Position{X: 68709.92, Y: -40044.953, Z: 512.03125},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "APARTMENT BLOCK",
						Name:   "Apartment Block",
						Center: Position{X: 69209.92, Y: -1084.9531, Z: 512.03125},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "BOMBARDED RIVERFRONT",
						Name:   "Bombarded Riverfront",
						Center: Position{X: 69209.92, Y: 40080.047, Z: 512.03125},
						Radius: 7000.0,
					},
				},
			},
		},
	},
	SECTORS_SMOLENSK_SMALL: {
		{
			From:         GridCoordinate{X: -5, Y: -4},
			To:           GridCoordinate{X: -2, Y: 3},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -1, Y: -4},
			To:   GridCoordinate{X: 0, Y: 3},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "ZHELYABOVA SQUARE",
						Name:   "Zhelyabova Square",
						Center: Position{X: 1770.0, Y: 0.0, Z: -1735.0},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: 1, Y: -4},
			To:           GridCoordinate{X: 4, Y: 3},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_STALINGRAD_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "MAMAYEV APPROACH",
						Name:   "Mamayev Approach",
						Center: Position{X: -69500.0, Y: -47966.0, Z: 5684.0},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "NAIL FACTORY",
						Name:   "Nail Factory",
						Center: Position{X: -71016.0, Y: 11068.0, Z: 7295.0},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "CITY OVERLOOK",
						Name:   "City Overlook",
						Center: Position{X: -69346.0, Y: 48417.0, Z: 7445.0},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "DOLGIY RAVINE",
						Name:   "Dolgiy Ravine",
						Center: Position{X: -39681.0, Y: -48845.0, Z: 4095.0},
						Radius: 7500.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "YELLOW HOUSE",
						Name:   "Yellow House",
						Center: Position{X: -39693.438, Y: -1.544678, Z: 7515.0},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "KOMSOMOL HQ",
						Name:   "Komsomol HQ",
						Center: Position{X: -39683.0, Y: 39676.0, Z: 7310.0},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "RAILWAY CROSSING",
						Name:   "Railway Crossing",
						Center: Position{X: 7.0, Y: -39673.0, Z: 4505.0},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CARRIAGE DEPOT",
						Name:   "Carriage Depot",
						Center: Position{X: -15.0, Y: 13.0, Z: 4961.0},
						Radius: 8500.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "TRAIN STATION",
						Name:   "Train Station",
						Center: Position{X: 6.0, Y: 39678.0, Z: 4942.0},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "HOUSE OF THE WORKERS",
						Name:   "House Of The Workers",
						Center: Position{X: 36591.0, Y: -40602.0, Z: 4355.0},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "PAVLOV'S HOUSE",
						Name:   "Pavlov's House",
						Center: Position{X: 48586.0, Y: 1452.0, Z: 4035.0},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "THE BREWERY",
						Name:   "The Brewery",
						Center: Position{X: 39674.0, Y: 41970.0, Z: 4077.0},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "L-SHAPED HOUSE",
						Name:   "L-Shaped House",
						Center: Position{X: 68875.0, Y: -35043.0, Z: 4195.0},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "GRUDININ'S MILL",
						Name:   "Grudinin's Mill",
						Center: Position{X: 70063.0, Y: -32.0, Z: 3965.0},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "VOLGA BANKS",
						Name:   "Volga Banks",
						Center: Position{X: 70121.0, Y: 43351.0, Z: 3965.0},
						Radius: 8000.0,
					},
				},
			},
		},
	},
	SECTORS_STALINGRAD_SMALL: {
		{
			From:         GridCoordinate{X: -5, Y: -4},
			To:           GridCoordinate{X: -2, Y: 3},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -1, Y: -4},
			To:   GridCoordinate{X: 0, Y: 3},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CARRIAGE DEPOT",
						Name:   "Carriage Depot",
						Center: Position{X: 300.0, Y: -220.0, Z: -1388.4307},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: 1, Y: -4},
			To:           GridCoordinate{X: 4, Y: 3},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_STMARIEDUMONT_LARGE: {
		{
			From: GridCoordinate{X: -3, Y: -5},
			To:   GridCoordinate{X: 2, Y: -4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -5},
					To:   GridCoordinate{X: -2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "WINTERS LANDING",
						Name:   "Winters Landing",
						Center: Position{X: -39503.258, Y: -78343.03, Z: 809.0},
						Radius: 5754.7485,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -5},
					To:   GridCoordinate{X: 0, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "LE GRAND CHEMIN",
						Name:   "Le Grand Chemin",
						Center: Position{X: -367.0, Y: -76667.0, Z: 809.0},
						Radius: 4500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -5},
					To:   GridCoordinate{X: 2, Y: -4},
					Strongpoint: Strongpoint{
						ID:     "THE BARN",
						Name:   "The Barn",
						Center: Position{X: 44896.0, Y: -73822.0, Z: -89.0},
						Radius: 5691.033,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: 2, Y: -2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "BRECOURT BATTERY",
						Name:   "Brecourt Battery",
						Center: Position{X: -39380.0, Y: -39702.0, Z: 809.0},
						Radius: 6078.6405,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "CATTLESHEDS",
						Name:   "Cattlesheds",
						Center: Position{X: 2961.319, Y: -41557.402, Z: 809.0},
						Radius: 5400.8775,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "RUE DE LA GARE",
						Name:   "Rue De La Gare",
						Center: Position{X: 35565.562, Y: -39370.594, Z: 809.0},
						Radius: 5892.7635,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -1},
			To:   GridCoordinate{X: 2, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "THE DUGOUT",
						Name:   "The Dugout",
						Center: Position{X: -37170.906, Y: -151.18701, Z: 460.30884},
						Radius: 5814.819,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "AA NETWORK",
						Name:   "AA Network",
						Center: Position{X: 1716.0, Y: 2530.0, Z: 422.64746},
						Radius: 6530.8635,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "PIERRE'S FARM",
						Name:   "Pierre's Farm",
						Center: Position{X: 37508.1, Y: 1336.6963, Z: 438.40137},
						Radius: 5207.283,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 1},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "HUGO'S FARM",
						Name:   "Hugo's Farm",
						Center: Position{X: -38001.0, Y: 38089.0, Z: 809.0},
						Radius: 6046.29,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "THE HAMLET",
						Name:   "The Hamlet",
						Center: Position{X: -2158.7668, Y: 42649.312, Z: 597.02246},
						Radius: 4500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "STE MARIE DU MONT",
						Name:   "Ste Marie Du Mont",
						Center: Position{X: 47022.125, Y: 50258.56, Z: 1336.7599},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: 3},
			To:   GridCoordinate{X: 2, Y: 4},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: 3},
					To:   GridCoordinate{X: -2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "THE CORNER",
						Name:   "The Corner",
						Center: Position{X: -34620.76, Y: 69152.766, Z: 809.0},
						Radius: 5105.9565,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 3},
					To:   GridCoordinate{X: 0, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "HILL 6",
						Name:   "Hill 6",
						Center: Position{X: 142.14453, Y: 76822.93, Z: 467.88086},
						Radius: 4500.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 3},
					To:   GridCoordinate{X: 2, Y: 4},
					Strongpoint: Strongpoint{
						ID:     "THE FIELDS",
						Name:   "The Fields",
						Center: Position{X: 39750.15, Y: 78234.78, Z: 1152.4478},
						Radius: 4500.0,
					},
				},
			},
		},
	},
	SECTORS_STMARIEDUMONT_SMALL: {
		{
			From:         GridCoordinate{X: -4, Y: -5},
			To:           GridCoordinate{X: 3, Y: -2},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -4, Y: -1},
			To:   GridCoordinate{X: 3, Y: 0},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CATTLESHEDS",
						Name:   "Cattlesheds",
						Center: Position{X: 3670.0, Y: -86553.4, Z: 125.0},
						Radius: 10000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: -4, Y: 1},
			To:           GridCoordinate{X: 3, Y: 4},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_STMEREEGLISE_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "FLAK POSITION",
						Name:   "Flak Position",
						Center: Position{X: -69311.0, Y: -40772.0, Z: -81.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "VAULAVILLE",
						Name:   "Vaulaville",
						Center: Position{X: -62223.0, Y: -3146.0, Z: -1163.0},
						Radius: 2507.38,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "LA PRAIRIE",
						Name:   "La Prairie",
						Center: Position{X: -67517.0, Y: 35037.0, Z: 1050.0},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ROUTE DU HARAS",
						Name:   "Route Du Haras",
						Center: Position{X: -40886.0, Y: -37779.0, Z: -688.63086},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "WESTERN APPROACH",
						Name:   "Western Approach",
						Center: Position{X: -32652.0, Y: -14761.0, Z: -400.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "RUE DE GAMBOSVILLE",
						Name:   "Rue De Gambosville",
						Center: Position{X: -34553.0, Y: 41733.0, Z: -699.0},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "HOSPICE",
						Name:   "Hospice",
						Center: Position{X: -1100.0, Y: -46000.0, Z: -400.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "SAINTE-MÈRE-ÉGLISE",
						Name:   "Sainte-Mère-Église",
						Center: Position{X: 5949.0, Y: -7436.0, Z: -718.7539},
						Radius: 4741.136,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "CHECKPOINT",
						Name:   "Checkpoint",
						Center: Position{X: 467.0, Y: 32490.0, Z: -1330.0},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ARTILLERY BATTERY",
						Name:   "Artillery Battery",
						Center: Position{X: 39652.0, Y: -34374.0, Z: -400.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "THE CEMETERY",
						Name:   "The Cemetery",
						Center: Position{X: 28858.0, Y: 5593.0, Z: -742.64844},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "MAISON DU CRIQUE",
						Name:   "Maison Du Crique",
						Center: Position{X: 25884.0, Y: 30530.0, Z: -400.0},
						Radius: 4000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "LES VIEUX VERGERS",
						Name:   "Les Vieux Vergers",
						Center: Position{X: 70168.0, Y: -28861.0, Z: -743.6719},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CROSS ROADS",
						Name:   "Cross Roads",
						Center: Position{X: 72279.0, Y: 1393.0, Z: -676.0},
						Radius: 4000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "RUISSEAU DE FERME",
						Name:   "Russeau De Ferme",
						Center: Position{X: 72138.0, Y: 38912.0, Z: -1125.0},
						Radius: 4000.0,
					},
				},
			},
		},
	},
	SECTORS_STMEREEGLISE_SMALL: {
		{
			From:         GridCoordinate{X: -5, Y: -4},
			To:           GridCoordinate{X: -2, Y: 3},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -1, Y: -4},
			To:   GridCoordinate{X: 0, Y: 3},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "SAINTE-MÈRE-ÉGLISE",
						Name:   "Sainte-Mère-Église",
						Center: Position{X: 165.0, Y: -113.0, Z: -76.0},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: 1, Y: -4},
			To:           GridCoordinate{X: 4, Y: 3},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_TOBRUK_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "GUARD ROOM",
						Name:   "Guard Room",
						Center: Position{X: -68855.0, Y: -27530.0, Z: -6107.8867},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "TANK GRAVEYARD",
						Name:   "Tank Graveyard",
						Center: Position{X: -69405.0, Y: -2075.0, Z: -7493.6714},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "DIVISION HEADQUARTERS",
						Name:   "Division Headquarters",
						Center: Position{X: -69835.0, Y: 45005.0, Z: -8821.878},
						Radius: 7000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "WEST CREEK",
						Name:   "West Creek",
						Center: Position{X: -40077.0, Y: -44015.0, Z: -5846.329},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "ALBERGO RISTORANTE MODERNO",
						Name:   "Albergo Ristorante Moderno",
						Center: Position{X: -29536.549, Y: 2241.0, Z: -7323.1304},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "KING SQUARE",
						Name:   "King Square",
						Center: Position{X: -29485.152, Y: 39744.316, Z: -7213.4697},
						Radius: 8000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "DESERT RAT CAVES",
						Name:   "Desert Rat Caves",
						Center: Position{X: 31.221313, Y: -39665.25, Z: -5381.714},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CHURCH GROUNDS",
						Name:   "Church Grounds",
						Center: Position{X: 59.92363, Y: 11770.049, Z: -7002.9688},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "ADMIRALTY HOUSE",
						Name:   "Admiralty House",
						Center: Position{X: 7919.4326, Y: 48901.832, Z: -7326.4785},
						Radius: 9000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ABANDONED AMMO CACHE",
						Name:   "Abandoned Ammo Cache",
						Center: Position{X: 39687.508, Y: -39659.996, Z: -4662.4526},
						Radius: 8000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "8TH ARMY MEDICAL HOSPITAL",
						Name:   "8th Army Medical Hospital",
						Center: Position{X: 40124.035, Y: -2845.0, Z: -7074.448},
						Radius: 9000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "SUPPLY DUMP",
						Name:   "Supply Dump",
						Center: Position{X: 39820.547, Y: 43918.36, Z: -7314.533},
						Radius: 9000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "ROAD TO SENUSSI MINE",
						Name:   "Road To Senussi Mine",
						Center: Position{X: 69522.8, Y: -40790.0, Z: -4568.7607},
						Radius: 7000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "MAKESHIFT AID STATION",
						Name:   "Makeshift Aid Station",
						Center: Position{X: 69380.0, Y: 25.0, Z: -6849.7656},
						Radius: 9000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "CARGO WAREHOUSES",
						Name:   "Cargo Warehouses",
						Center: Position{X: 70047.33, Y: 41171.96, Z: -7616.6694},
						Radius: 8000.0,
					},
				},
			},
		},
	},
	SECTORS_TOBRUK_SMALL: {
		{
			From:         GridCoordinate{X: -5, Y: -4},
			To:           GridCoordinate{X: -2, Y: 3},
			CaptureZones: []CaptureZone{},
		},
		{
			From: GridCoordinate{X: -1, Y: -4},
			To:   GridCoordinate{X: 0, Y: 3},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "CHURCH GROUNDS",
						Name:   "Church Grounds",
						Center: Position{X: 0.0, Y: 7625.0, Z: -844.0},
						Radius: 6000.0,
					},
				},
			},
		},
		{
			From:         GridCoordinate{X: 1, Y: -4},
			To:           GridCoordinate{X: 4, Y: 3},
			CaptureZones: []CaptureZone{},
		},
	},
	SECTORS_UTAHBEACH_LARGE: {
		{
			From: GridCoordinate{X: -5, Y: -3},
			To:   GridCoordinate{X: -4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -5, Y: -3},
					To:   GridCoordinate{X: -4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "MAMMUT RADAR",
						Name:   "Mammut Radar",
						Center: Position{X: -65158.0, Y: -51522.0, Z: -2401.0},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: -1},
					To:   GridCoordinate{X: -4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "FLOODED HOUSE",
						Name:   "Flooded House",
						Center: Position{X: -66464.0, Y: -2944.0, Z: -2388.0},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: -5, Y: 1},
					To:   GridCoordinate{X: -4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "SAINTE MARIE APPROACH",
						Name:   "Sainte Marie Approach",
						Center: Position{X: -64837.0, Y: 52705.0, Z: -2157.0},
						Radius: 3000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -3, Y: -3},
			To:   GridCoordinate{X: -2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -3, Y: -3},
					To:   GridCoordinate{X: -2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "SUNKEN BRIDGE",
						Name:   "Sunken Bridge",
						Center: Position{X: -30810.0, Y: -47853.0, Z: -2157.0},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: -1},
					To:   GridCoordinate{X: -2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "LA GRANDE CRIQUE",
						Name:   "La Grande Crique",
						Center: Position{X: -28986.0, Y: 508.0, Z: -2388.0},
						Radius: 3294.095,
					},
				},
				{
					From: GridCoordinate{X: -3, Y: 1},
					To:   GridCoordinate{X: -2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "DROWNED FIELDS",
						Name:   "Drowned Fields",
						Center: Position{X: -36400.0, Y: 43928.0, Z: -2157.0},
						Radius: 3000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: -1, Y: -3},
			To:   GridCoordinate{X: 0, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: -1, Y: -3},
					To:   GridCoordinate{X: 0, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "WN4",
						Name:   "WN4",
						Center: Position{X: 5359.0, Y: -42267.0, Z: -2283.0},
						Radius: 3742.354,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: -1},
					To:   GridCoordinate{X: 0, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "THE CHAPEL",
						Name:   "The Chapel",
						Center: Position{X: 10650.0, Y: -7786.0, Z: -2157.0},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: -1, Y: 1},
					To:   GridCoordinate{X: 0, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "WN7",
						Name:   "WN7",
						Center: Position{X: 727.0, Y: 50408.0, Z: -2157.0},
						Radius: 5000.0,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 1, Y: -3},
			To:   GridCoordinate{X: 2, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 1, Y: -3},
					To:   GridCoordinate{X: 2, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "AA BATTERY",
						Name:   "AA Battery",
						Center: Position{X: 35101.0, Y: -43589.0, Z: -2389.0},
						Radius: 3941.9941,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: -1},
					To:   GridCoordinate{X: 2, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "HILL 5",
						Name:   "Hill 5",
						Center: Position{X: 36131.0, Y: -1838.0, Z: -2157.0},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: 1, Y: 1},
					To:   GridCoordinate{X: 2, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "WN5",
						Name:   "WN5",
						Center: Position{X: 48763.0, Y: 44080.0, Z: -2247.0},
						Radius: 3194.091,
					},
				},
			},
		},
		{
			From: GridCoordinate{X: 3, Y: -3},
			To:   GridCoordinate{X: 4, Y: 2},
			CaptureZones: []CaptureZone{
				{
					From: GridCoordinate{X: 3, Y: -3},
					To:   GridCoordinate{X: 4, Y: -2},
					Strongpoint: Strongpoint{
						ID:     "TARE GREEN",
						Name:   "Tare Green",
						Center: Position{X: 63845.0, Y: -46581.0, Z: -2284.0},
						Radius: 3000.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: -1},
					To:   GridCoordinate{X: 4, Y: 0},
					Strongpoint: Strongpoint{
						ID:     "RED ROOF HOUSE",
						Name:   "Red Roof House",
						Center: Position{X: 64923.67, Y: 3144.1865, Z: -2206.6382},
						Radius: 3250.0,
					},
				},
				{
					From: GridCoordinate{X: 3, Y: 1},
					To:   GridCoordinate{X: 4, Y: 2},
					Strongpoint: Strongpoint{
						ID:     "UNCLE RED",
						Name:   "Uncle Red",
						Center: Position{X: 66675.0, Y: 45162.0, Z: -2157.0},
						Radius: 2823.963,
					},
				},
			},
		},
	},
}
