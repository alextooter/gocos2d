package gocos2d

type (
	Rectangle struct {
		*Anchor
		width, height int
	}
	Position struct {
		x, y int
	}
	Camera struct {
	}
	Grid struct {
	}

	Anchor   Position
	Rotation float32
	Scale    struct {
		width, height int
		Factor        float32
	}
	Skew          Position
	ShaderProgram string
	BoundingBox   Rectangle
	Tag           string
	ZOrder        float32
)
