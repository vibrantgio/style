# style

The type scale of [Vibrant Gio](https://github.com/vibrantgio), a design system
for native desktop applications on macOS, Windows and Linux, written in pure Go
on [Gio](https://gioui.org) — fourteen named text styles and the font
collection they are drawn with, in one file.

An application drawing its own text has two problems that have nothing to do
with each other and always arrive together: *which faces does my shaper load*,
and *how big is a heading*. style answers both in one import. `FontFaces()`
returns the Roboto faces to build a `*text.Shaper` from, and `H1` … `Overline`
are ready-made `textdraw.TextStyle` values — a font, a size, an alignment, a
line limit and a truncator — to hand to
[textdraw](https://github.com/vibrantgio/textdraw)'s drawing calls.

**This module is frozen.** It is the Material Design 2 scale, a
package-level table rather than a theme value, so it cannot vary with the
theme, the platform or the user's preferences. Type in a Vibrant Gio
application comes from the
[`theme/tokens`](https://github.com/vibrantgio/theme) `Typography` theme
token, which every symbol here names in its `Deprecated:` marker. Nothing
here will be extended — see Status before building on it.

## Where it sits

Tier 0 of the stack — `mvu → theme → components → effects → patterns → markdown` —
and the two intra-tier edges the tier table has to admit by name: style imports
[font](https://github.com/vibrantgio/font) and
[textdraw](https://github.com/vibrantgio/textdraw), both also tier 0. The
layering check permits those two edges rather than renumbering the tier. The
[organization page](https://github.com/vibrantgio) has the full tier table.

```sh
go get github.com/vibrantgio/style
```

The Gio and Go versions are the ones `go.mod` names.

## Packages

One package, at the module root.

| Symbol | |
| --- | --- |
| `FontFaces()` | The five upright Roboto weights the scale names — Thin, Light, Normal, Medium, Bold — as `[]textdraw.FontFace`, which is a type *alias* for `gioui.org/font.FontFace`, so it drops straight into `text.WithCollection`. |
| `H1` … `H6` | Headings, 96/60/48/34/24/20 sp, single-line, ellipsis-truncated. `H1` is Thin, `H2` is Light, the rest are Normal. |
| `Subtitle1`, `Subtitle2` | 16 sp Normal and 14 sp Medium, single-line. |
| `BodyText1`, `BodyText2` | 16 sp and 14 sp Normal, and the only two styles with `MaxLines: 0` — the only two that wrap. |
| `Button`, `SmallButton` | 14 sp Medium and 12 sp Bold, single-line. |
| `Caption`, `Overline` | 12 sp and 10 sp Normal, single-line. |

Every style is `Alignment: textdraw.Start` and `Truncator: "…"`. `Size` is
`unit.Sp`, not `unit.Dp`, so Gio converts it through the window's `PxPerSp`
metric, and every style but the two body styles caps at one line.

## Usage

Nothing new should be written against this module — a Vibrant Gio application
takes its typography from the theme (`Typography.Shaper()` and the typography roles;
see the [theme](https://github.com/vibrantgio/theme) README). This
section documents how a program that already uses this module draws with it.

Build the shaper once per window, at layer-building scope, and hand it to
everything that draws:

```go
shaper := text.NewShaper(text.WithCollection(style.FontFaces()))
```

Note what is not passed: `text.NoSystemFonts()`. Leaving system fonts loaded
keeps a fallback for glyphs Roboto lacks and for named families such as the
generic `monospace`. And do not
append gofont to the collection — that loads two type systems into one window,
and the mixing shows up only on the glyphs that fall through.

The styles are ready-made `textdraw.TextStyle` values, and drawing through
[textdraw](https://github.com/vibrantgio/textdraw) is the pairing this module
exists for — the style holds the font and size, `MeasureText` gives the
line height, `FillText`/`Text` put glyphs on screen:

```go
layout.UniformInset(12).Layout(gtx, textdraw.Text(shaper, style.H3, 0.0, 0.0, Grey900, "Gravity Well"))
layout.UniformInset(12).Layout(gtx, textdraw.Text(shaper, style.H4, 1.0, 1.0, Grey900, fmt.Sprint(fps, "fps")))
```

## For coding assistants

Read this repository's [AGENTS.md](./AGENTS.md) before you write code against this module.

Build and test the root module with `go build ./...` and `go test ./...`.

## Status

Honest about what does not work yet. The module is frozen, so none of it
will be fixed.

- **This is a second type system.** The `theme/tokens.Typography` theme
  token — typeface, weight, size, line height and tracking per typography role, plus
  the one shared shaper — which every symbol here names in its `Deprecated:`
  marker, is a separate scale. An application that styles some text with `H5`
  and other text from `Typography` is mixing two type systems.
- **The scale cannot vary.** These are package-level `var`s, so there is one
  scale per binary: no theme, no density, no platform and no accessibility
  preference can change a size. `theme/a11y` publishes the OS text-scaling
  and reduce-motion preferences and nothing here reads them.
- **`FontFaces()` returns five of the twelve Roboto faces
  [font](https://github.com/vibrantgio/font) ships.** Thin, Light, Normal,
  Medium and Bold — no Black, and no italic at any weight. A shaper built from
  this collection alone cannot render italic Roboto; it falls through to a
  system face, or to nothing under `text.NoSystemFonts()`. There is no
  `FontFacesWithItalic`, and no way to extend the collection except building
  the slice yourself.
- **`SmallButton` is not an MD2 style.** The MD2 scale has thirteen styles;
  this is a fourteenth, 12 sp Bold. It has no counterpart in
  the `Typography` token either.
- **Only `BodyText1` and `BodyText2` wrap.** Every other style sets
  `MaxLines: 1`, so a heading longer than its box is silently truncated with an
  ellipsis rather than flowing. That is usually what a heading wants and never
  what a paragraph does; if you are drawing a paragraph with anything but the
  two body styles, copy the value and clear `MaxLines`.
- **There are no tests.** `go test ./...` reports "no test files".

## License

MIT — see [LICENSE](./LICENSE).
