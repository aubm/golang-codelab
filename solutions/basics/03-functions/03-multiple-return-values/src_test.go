package basics_test

import (
	"testing"

	. "github.com/aubm/golang-codelab/solutions/basics/03-functions/03-multiple-return-values"
)

func TestComputePolygonCenter(t *testing.T) {
	for _, d := range []struct {
		Coordinates          [][2]float64
		ExpectedX, ExpectedY float64
	}{
		{
			Coordinates: [][2]float64{{23.55, 54.16}, {43.24, 54.16}, {43.24, 61.43}, {23.55, 61.43}},
			ExpectedX:   33.395,
			ExpectedY:   57.795,
		},
		{
			Coordinates: [][2]float64{{36.650390625, 57.37393841871411}, {37.529296875, 50.62507306341435}, {51.064453125, 49.89463439573421}, {63.896484375, 50.792047064406866}, {63.10546874999999, 59.31076795603884}, {56.689453125, 64.51064316846676}, {37.353515625, 64.28275952823394}},
			ExpectedX:   49.46986607142857,
			ExpectedY:   56.68426622785844,
		},
	} {
		if actualX, actualY := ComputePolygonCenter(d.Coordinates); actualX != d.ExpectedX || actualY != d.ExpectedY {
			t.Errorf("when coordinates are %v, expected center to be [%v, %v], but got [%v, %v]",
				d.Coordinates, d.ExpectedX, d.ExpectedY, actualX, actualY)
		}
	}
}
