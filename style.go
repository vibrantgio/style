package style

import (
	text "github.com/vibrantgio/text"

	"github.com/vibrantgio/font/roboto/regular/bold"
	"github.com/vibrantgio/font/roboto/regular/light"
	"github.com/vibrantgio/font/roboto/regular/medium"
	"github.com/vibrantgio/font/roboto/regular/normal"
	"github.com/vibrantgio/font/roboto/regular/thin"
)

var (
	H1          = text.TextStyle{Font: thin.Font, Alignment: text.Start, Size: 96, MaxLines: 1, Truncator: "…"}
	H2          = text.TextStyle{Font: light.Font, Alignment: text.Start, Size: 96, MaxLines: 1, Truncator: "…"}
	H3          = text.TextStyle{Font: normal.Font, Alignment: text.Start, Size: 48, MaxLines: 1, Truncator: "…"}
	H4          = text.TextStyle{Font: normal.Font, Alignment: text.Start, Size: 34, MaxLines: 1, Truncator: "…"}
	H5          = text.TextStyle{Font: normal.Font, Alignment: text.Start, Size: 24, MaxLines: 1, Truncator: "…"}
	H6          = text.TextStyle{Font: normal.Font, Alignment: text.Start, Size: 20, MaxLines: 1, Truncator: "…"}
	Subtitle1   = text.TextStyle{Font: normal.Font, Alignment: text.Start, Size: 16, MaxLines: 1, Truncator: "…"}
	Subtitle2   = text.TextStyle{Font: medium.Font, Alignment: text.Start, Size: 14, MaxLines: 1, Truncator: "…"}
	BodyText1   = text.TextStyle{Font: normal.Font, Alignment: text.Start, Size: 16, MaxLines: 0, Truncator: "…"}
	BodyText2   = text.TextStyle{Font: normal.Font, Alignment: text.Start, Size: 14, MaxLines: 0, Truncator: "…"}
	Button      = text.TextStyle{Font: medium.Font, Alignment: text.Start, Size: 14, MaxLines: 1, Truncator: "…"}
	Caption     = text.TextStyle{Font: normal.Font, Alignment: text.Start, Size: 12, MaxLines: 1, Truncator: "…"}
	SmallButton = text.TextStyle{Font: bold.Font, Alignment: text.Start, Size: 12, MaxLines: 1, Truncator: "…"}
	Overline    = text.TextStyle{Font: normal.Font, Alignment: text.Start, Size: 10, MaxLines: 1, Truncator: "…"}
)

func FontFaces() []text.FontFace {
	return []text.FontFace{
		thin.FontFace(),
		light.FontFace(),
		normal.FontFace(),
		medium.FontFace(),
		bold.FontFace(),
	}
}
