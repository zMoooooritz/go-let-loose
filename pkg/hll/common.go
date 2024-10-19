package hll

type TeamData struct {
	Allies int
	Axis   int
}

type GameState struct {
	PlayerCount      TeamData
	GameScore        TeamData
	RemainingSeconds int
	CurrentMap       Layer
	NextMap          Layer
}
