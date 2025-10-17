package colorspace

import (
	"testing"
	"github.com/Diwice/color-picker/pkg/colorspace"	
)

// Is there REALLY any point in testing a struct creation?

func TestRGB(t *testing.T) {
	in := colorspace.RGB_obj{
		RED: 100,
		GREEN: 150,
		BLUE: 200,
	}	

	if in.RED != 100 || in.GREEN != 150 || in.BLUE != 200 {
		t.Errorf("Expected : R100/G150/B200; Got : R%d/G%d/B%d", in.RED, in.GREEN, in.BLUE)
	}

	in.RED += 10
	in.GREEN += 10
	in.BLUE += 10

	if in.RED != 110 || in.GREEN != 160 || in.BLUE != 210 {
		t.Errorf("Expected : R110/G160/B210 (Incremented); Got : R%d/G%d/B%d", in.RED, in.GREEN, in.BLUE) 
	}
}
