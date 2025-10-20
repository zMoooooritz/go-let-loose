package hll

import (
	"fmt"
	"math"
)

const (
	MAP_SIZE_CM = 200000.0
	MAP_MIN     = -MAP_SIZE_CM / 2
	MAP_MAX     = MAP_SIZE_CM / 2
)

type Position struct {
	X float64
	Y float64
	Z float64
}

func (p Position) String() string {
	return fmt.Sprintf("(%.1f, %.1f, %.1f)", p.X, p.Y, p.Z)
}

func (p Position) IsActive() bool {
	return p.X != 0 || p.Y != 0 || p.Z != 0
}

// the distance is measured in 1 unit = 1 cm on the 2x2km map
func (p Position) SpacialDistanceTo(pos Position) int {
	diffX := p.X - pos.X
	diffY := p.Y - pos.Y
	diffZ := p.Z - pos.Z
	return int(math.Sqrt(diffX*diffX + diffY*diffY + diffZ*diffZ))
}

// the distance is measured in 1 unit = 1 cm on the 2x2km map
func (p Position) PlanarDistanceTo(pos Position) int {
	diffX := p.X - pos.X
	diffY := p.Y - pos.Y
	return int(math.Sqrt(diffX*diffX + diffY*diffY))
}

func (p Position) BearingTo(pos Position) float64 {
	deltaX := pos.X - p.X
	deltaY := pos.Y - p.Y

	radians := math.Atan2(deltaY, deltaX)
	degrees := radians * 180 / math.Pi

	// Convert to compass bearing (0-360 degrees, 0 = North)
	bearing := 90 - degrees
	if bearing < 0 {
		bearing += 360
	}
	return bearing
}

func (p Position) DirectionTo(pos Position) string {
	bearing := p.BearingTo(pos)
	directions := []string{"N", "NE", "E", "SE", "S", "SW", "W", "NW"}
	index := int((bearing+22.5)/45) % 8
	return directions[index]
}

func (p Position) IsWithinMapBounds() bool {
	return p.X >= MAP_MIN && p.X <= MAP_MAX &&
		p.Y >= MAP_MIN && p.Y <= MAP_MAX
}

func (p Position) ToGridReference() string {
	if !p.IsActive() {
		return "N/A"
	}

	if !p.IsWithinMapBounds() {
		return "OUT_OF_BOUNDS"
	}

	// Convert to 0-1 range
	normalizedX := (p.X - MAP_MIN) / MAP_SIZE_CM
	normalizedY := (p.Y - MAP_MIN) / MAP_SIZE_CM

	// Convert to grid (assuming 10x10 grid)
	gridX := int(normalizedX * 10)
	gridY := int(normalizedY * 10)

	return fmt.Sprintf("%c%d", 'A'+gridX, gridY+1)
}
