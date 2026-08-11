# style

The type scale of [Vibrant Gio](https://github.com/vibrantgio), a design system
for native desktop applications on macOS, Windows and Linux, written in pure Go
on [Gio](https://gioui.org) — fourteen named text styles and the font
collection they are drawn with, in one file of about forty lines.

An application drawing its own text has two problems that have nothing to do
with each other and always arrive together: *which faces does my shaper load*,
and *how big is a heading*. style answers both in one import. `FontFaces()`
returns the Roboto faces to build a `*text.Shaper` from, and `H1` … `Overline`
are ready-made `textdraw.TextStyle` values — a font, a size, an alignment, a
line limit and a truncator — to hand to
[textdraw](https://github.com/vibrantgio/textdraw)'s drawing calls. That is
why the example programs in the organization that still draw text this way
import exactly these two modules together — the workbench applications used
to as well, until F1 moved them onto the theme's typography.

**This module is frozen.** It is the Material Design 2 scale, and Vibrant Gio
targets MD3; more importantly the scale is a package-level table rather than a
theme value, so it cannot vary with the theme, the platform or the user's
preferences. ADR-003 supersedes it with the
[`theme/tokens`](https://github.com/vibrantgio/theme) `Typography` theme
token, and every symbol here carries a `Deprecated:` marker naming its MD3
replacement. The module was kept for the deprecation window so existing
imports kept compiling; nothing here will be extended, and only correctness
fixes land — see Status before building on it.

**The window is closed.** F3.4 of the
[org plan](https://github.com/vibrantgio/.github) landed on 2026-08-06 and
ended it. G-F1 had already moved the last in-org consumers off this module,
so there was nothing left to sweep: the deprecation shims other repositories
carried were deleted in their own major releases, and style, which is
nothing *but* deprecated surface, is instead **archived at v0.0.6**. It is
frozen there and will never be re-tagged. The freeze ends by archival, not
deletion — the v0.0.6 tag and the import path keep resolving forever; the
repository just stops changing. There will be no v0.0.7.

## Where it sits

Tier 0 of the stack — `mvu → theme → components → effects → patterns → markdown` —
and the one intra-tier edge ADR-001's table has to admit by name: style imports
[font](https://github.com/vibrantgio/font) and
[textdraw](https://github.com/vibrantgio/textdraw), both also tier 0. The
layering check permits that single edge rather than renumbering the tier. The
[organization page](https://github.com/vibrantgio) has the full tier table.

Nothing inside the design system imports style — not components, not patterns, not
markdown. Its consumers were all applications, and most of them are gone: F1
migrated every [workbench](https://github.com/vibrantgio/workbench)
application and every `mvu/example` program off it. What remains is eleven
example mains in the support repositories — four under
`ivg/raster/gio/example`, three under `svg/driver/gio/example` and the four
`traer/gio` demos. F3.4 archived the repository with those still pointing at
v0.0.6, which keeps resolving after archival — which is the whole point of
archiving rather than deleting.

```sh
go get github.com/vibrantgio/style
```

Every module in the organization is on gioui.org v0.10.1 and Go 1.25.1.

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
`unit.Sp`, not `unit.Dp` — it scales with the user's text-size preference, and
the twelve heading and label styles cap at one line.

## Usage

Nothing new should be written against this module — a Vibrant Gio application
takes its typography from the theme (`Typography.Shaper()` and the MD3 roles;
see the [theme](https://github.com/vibrantgio/theme) README). This
section documents the pattern the remaining consumers use.

Build the shaper once per window, at layer-building scope, and hand it to
everything that draws. It is the same line, character for character, in all
eleven programs that call `FontFaces()`:

```go
shaper := text.NewShaper(text.WithCollection(style.FontFaces()))
```

Note what is not passed: `text.NoSystemFonts()`. Leaving system fonts loaded
keeps a fallback for glyphs Roboto lacks and for named families such as the
generic `monospace` that markdown's code spans resolve through. And do not
append gofont to the collection — that loads two type systems into one window,
and the mixing shows up only on the glyphs that fall through.

The styles are ready-made `textdraw.TextStyle` values, and drawing through
[textdraw](https://github.com/vibrantgio/textdraw) is the pairing this module
exists for — the style carries the font and size, `MeasureText` gives the
line height, `FillText`/`Text` put glyphs on screen. From
`traer/gio/gravity`:

```go
layout.UniformInset(12).Layout(gtx, textdraw.Text(shaper, style.H3, 0.0, 0.0, Grey900, "Gravity Well"))
layout.UniformInset(12).Layout(gtx, textdraw.Text(shaper, style.H4, 1.0, 1.0, Grey900, fmt.Sprint(fps, "fps")))
```

## For coding assistants

Read the canonical guide before writing code against this module — the module
inventory with current tags, the application skeleton, MVU and rx semantics,
typography, and the pitfalls that are not guessable:

<https://raw.githubusercontent.com/vibrantgio/.github/master/llms.txt>

[`AGENTS.md`](./AGENTS.md) in this repository has the build and test commands.

## Status

Honest about what does not work yet. Every number below is measured against the
source, not estimated. None of it will be fixed: v0.0.6 is the last tag.

- **This module is archived at v0.0.6 and frozen there.** ADR-003 opened a
  deprecation window on it; F3.4 closed the window by archiving the
  repository rather than sweeping it, because by then G-F1 had moved the
  last in-org consumer off it and there was nothing left to remove that
  removing the whole module did not already cover. Every item below is
  therefore permanent, not pending.
- **This module was never wired into the component stack, and now never will
  be.** No library source file in components, effects, patterns, markdown or theme
  imports style; the components style their text from the
  `theme/tokens.Typography` theme token Phase C shipped — typeface, weight,
  size, line height and tracking per MD3 role, plus the one shared shaper —
  and C1.4 marked every symbol here `Deprecated:` with that replacement. An
  application that styles its own text with `H5` and drops a `patterns`
  heading beside it is mixing two type systems; F1 removed the last such
  mixture from the workbench applications.
- **`H2` was 96 sp until C1.4 — the same as `H1`.** MD2's H2 is 60, and the
  two differed only in weight — Thin and Light — so a document using both got
  no size hierarchy at all. It is 60 now, which means a program pinned to a
  version before the fix draws `H2` at 96, not the 60 this README documents.
- **The scale cannot vary.** These are package-level `var`s, so there is one
  scale per binary: no theme, no density, no platform and no accessibility
  preference can change a size. `theme/a11y` publishes the OS text-scaling
  and reduce-motion preferences and nothing here reads them.
- **`FontFaces()` returns five of the twelve faces
  [font](https://github.com/vibrantgio/font) ships.** Thin, Light, Normal,
  Medium and Bold — no Black, and no italic at any weight. A shaper built from
  this collection alone cannot render italic Roboto; it falls through to a
  system face, or to nothing under `text.NoSystemFonts()`. There is no
  `FontFacesWithItalic`, and no way to extend the collection except building
  the slice yourself.
- **Eleven of the fourteen styles have no consumer anywhere in the
  organization.** After the F1 migration only `H3` (4 call sites), `H4` (4)
  and `H5` (7) are still drawn with, all in the support repositories'
  example programs; `FontFaces` keeps its 11 call sites in the same programs.
  Everything else — every subtitle, body, button, caption and overline style
  — is imported by nothing.
- **`SmallButton` is not an MD2 role.** The MD2 scale has thirteen styles;
  this is a fourteenth, 12 sp Bold, invented here. It has no counterpart in
  the MD3 `Typography` token either, so there is nothing to migrate it to.
- **Only `BodyText1` and `BodyText2` wrap.** Every other style sets
  `MaxLines: 1`, so a heading longer than its box is silently truncated with an
  ellipsis rather than flowing. That is usually what a heading wants and never
  what a paragraph does; if you are drawing a paragraph with anything but the
  two body styles, copy the value and clear `MaxLines`.
- **There are no tests.** `go test ./...` reports "no test files".

## License

MIT — see [LICENSE](./LICENSE).
