package style

import (
	"github.com/vibrantgio/textdraw"

	"github.com/vibrantgio/font/roboto/regular/bold"
	"github.com/vibrantgio/font/roboto/regular/light"
	"github.com/vibrantgio/font/roboto/regular/medium"
	"github.com/vibrantgio/font/roboto/regular/normal"
	"github.com/vibrantgio/font/roboto/regular/thin"
)

var (
	H1          = textdraw.TextStyle{Font: thin.Font, Alignment: textdraw.Start, Size: 96, MaxLines: 1, Truncator: "…"}
	H2          = textdraw.TextStyle{Font: light.Font, Alignment: textdraw.Start, Size: 96, MaxLines: 1, Truncator: "…"}
	H3          = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 48, MaxLines: 1, Truncator: "…"}
	H4          = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 34, MaxLines: 1, Truncator: "…"}
	H5          = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 24, MaxLines: 1, Truncator: "…"}
	H6          = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 20, MaxLines: 1, Truncator: "…"}
	Subtitle1   = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 16, MaxLines: 1, Truncator: "…"}
	Subtitle2   = textdraw.TextStyle{Font: medium.Font, Alignment: textdraw.Start, Size: 14, MaxLines: 1, Truncator: "…"}
	BodyText1   = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 16, MaxLines: 0, Truncator: "…"}
	BodyText2   = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 14, MaxLines: 0, Truncator: "…"}
	Button      = textdraw.TextStyle{Font: medium.Font, Alignment: textdraw.Start, Size: 14, MaxLines: 1, Truncator: "…"}
	Caption     = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 12, MaxLines: 1, Truncator: "…"}
	SmallButton = textdraw.TextStyle{Font: bold.Font, Alignment: textdraw.Start, Size: 12, MaxLines: 1, Truncator: "…"}
	Overline    = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 10, MaxLines: 1, Truncator: "…"}
)

func FontFaces() []textdraw.FontFace {
	return []textdraw.FontFace{
		thin.FontFace(),
		light.FontFace(),
		normal.FontFace(),
		medium.FontFace(),
		bold.FontFace(),
	}
}
