// Package style is the standalone Material Design 2 type scale of Vibrant
// Gio: fourteen named textdraw.TextStyle values and FontFaces, the five
// upright Roboto faces they are drawn with.
//
// Deprecated: this package is frozen. Its MD2 scale is superseded by the MD3
// Typography theme token of github.com/vibrantgio/theme/tokens; only
// correctness fixes land here through the deprecation window.
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
	// H1 is the MD2 headline 1: 96 sp Thin, single line.
	//
	// Deprecated: use the DisplayLarge role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	H1 = textdraw.TextStyle{Font: thin.Font, Alignment: textdraw.Start, Size: 96, MaxLines: 1, Truncator: "…"}

	// H2 is the MD2 headline 2: 60 sp Light, single line.
	//
	// Deprecated: use the DisplayMedium role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	H2 = textdraw.TextStyle{Font: light.Font, Alignment: textdraw.Start, Size: 60, MaxLines: 1, Truncator: "…"}

	// H3 is the MD2 headline 3: 48 sp Normal, single line.
	//
	// Deprecated: use the DisplaySmall role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	H3 = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 48, MaxLines: 1, Truncator: "…"}

	// H4 is the MD2 headline 4: 34 sp Normal, single line.
	//
	// Deprecated: use the HeadlineMedium role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	H4 = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 34, MaxLines: 1, Truncator: "…"}

	// H5 is the MD2 headline 5: 24 sp Normal, single line.
	//
	// Deprecated: use the HeadlineSmall role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	H5 = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 24, MaxLines: 1, Truncator: "…"}

	// H6 is the MD2 headline 6: 20 sp Normal, single line.
	//
	// Deprecated: use the TitleLarge role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	H6 = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 20, MaxLines: 1, Truncator: "…"}

	// Subtitle1 is the MD2 subtitle 1: 16 sp Normal, single line.
	//
	// Deprecated: use the TitleMedium role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	Subtitle1 = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 16, MaxLines: 1, Truncator: "…"}

	// Subtitle2 is the MD2 subtitle 2: 14 sp Medium, single line.
	//
	// Deprecated: use the TitleSmall role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	Subtitle2 = textdraw.TextStyle{Font: medium.Font, Alignment: textdraw.Start, Size: 14, MaxLines: 1, Truncator: "…"}

	// BodyText1 is the MD2 body 1: 16 sp Normal, wrapping.
	//
	// Deprecated: use the BodyLarge role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	BodyText1 = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 16, MaxLines: 0, Truncator: "…"}

	// BodyText2 is the MD2 body 2: 14 sp Normal, wrapping.
	//
	// Deprecated: use the BodyMedium role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	BodyText2 = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 14, MaxLines: 0, Truncator: "…"}

	// Button is the MD2 button label: 14 sp Medium, single line.
	//
	// Deprecated: use the LabelLarge role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	Button = textdraw.TextStyle{Font: medium.Font, Alignment: textdraw.Start, Size: 14, MaxLines: 1, Truncator: "…"}

	// Caption is the MD2 caption: 12 sp Normal, single line.
	//
	// Deprecated: use the BodySmall role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	Caption = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 12, MaxLines: 1, Truncator: "…"}

	// SmallButton is 12 sp Bold, single line. It is not an MD2 role — this
	// scale invented it — and it has no counterpart in the MD3 scale.
	//
	// Deprecated: use github.com/vibrantgio/theme/tokens.Typography
	// instead; it has no exact counterpart there, and the nearest role is
	// LabelMedium, 12 dp Medium.
	SmallButton = textdraw.TextStyle{Font: bold.Font, Alignment: textdraw.Start, Size: 12, MaxLines: 1, Truncator: "…"}

	// Overline is the MD2 overline: 10 sp Normal, single line.
	//
	// Deprecated: use the LabelSmall role of
	// github.com/vibrantgio/theme/tokens.Typography instead.
	Overline = textdraw.TextStyle{Font: normal.Font, Alignment: textdraw.Start, Size: 10, MaxLines: 1, Truncator: "…"}
)

// FontFaces returns the five upright Roboto faces the scale names — Thin,
// Light, Normal, Medium and Bold — as a collection to build a text shaper
// from.
//
// Deprecated: use github.com/vibrantgio/theme/tokens.DefaultTypography.Faces
// — or github.com/vibrantgio/font/roboto.FontFaces for the full twelve-face
// Roboto collection it holds — instead.
func FontFaces() []textdraw.FontFace {
	return []textdraw.FontFace{
		thin.FontFace(),
		light.FontFace(),
		normal.FontFace(),
		medium.FontFace(),
		bold.FontFace(),
	}
}
