package convert

import (
	"coord-converter/types"
	"github.com/twpayne/go-proj/v10"
	"log/slog"
)

func Convert(from string, to string, coordinates []types.Coordinate, logger *slog.Logger) []types.Coordinate {
	pj, err := proj.NewCRSToCRS(from, to, nil)
	if err != nil {
		panic(err)
	}
	var results []types.Coordinate
	for _, coordinate := range coordinates {
		src := proj.NewCoord(coordinate.X, coordinate.Y, 0, 0)

		dest, err := pj.Forward(src)
		if err != nil {
			logger.Error("invalid coordinates", from, to, coordinate)
		} else {
			results = append(results, types.Coordinate{X: dest.X(), Y: dest.Y()})
		}
	}
	return results
}
