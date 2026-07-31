# AGENTS.md — style

The MD2 type scale as a table of `textdraw.TextStyle` values — `H1` through
`H6`, `Subtitle1` and `Subtitle2`, `BodyText1` and `BodyText2`, `Button`,
`SmallButton`, `Caption` and `Overline` — plus `FontFaces()`, which returns
the five Roboto faces those styles name.

**Layer.** Tier 0 of ADR-001's table, and the one intra-tier edge that
table has to admit by name: `style` imports `font` and `textdraw`, both
also tier 0, so `check-layers.sh` (B2.3) permits that single edge rather
than renumbering the tier. Nothing in the design system imports style; its
consumers are example mains under mvu, ivg, svg and traer, and four of the
workbench applications.

**Read the canonical guide before you write code against this module.** It is
the organization's one agent guide — the module inventory with current tags,
the application skeleton, the MVU loop and rx semantics, typography, and the
pitfalls that are not guessable. It lives exactly once, in `vibrantgio/.github`,
and this file links it rather than copying it:

    https://raw.githubusercontent.com/vibrantgio/.github/master/llms.txt

**Module.** `github.com/vibrantgio/style`, one module at the repository
root.

**Build and test.** From the repository root:

    go build ./... && go test ./...

**Frozen — do not build on this.** ADR-003 moves the typeface into the theme:
`Typography` becomes a theme token carrying a full `TextStyle` per MD3 role —
typeface, weight, size, line height, tracking — plus the face collection and a
lazily built shaper. This module's MD2 scale is superseded by it. `style` is
frozen rather than deleted, so every existing import keeps compiling through
the deprecation window; the exported symbols get `Deprecated:` markers and the
`spectrum/tokens.Typography` replacement in task C1.4, and the shims go in the
major bump, F3.3.

It also carries a real bug that C1.4 fixes rather than inherits: `H1` and `H2`
are both 96 dp, where MD2's H2 is 60. Do not copy those numbers forward.

Nothing in the design system imports `style` — that is the point of ADR-003.
Its consumers are demo mains under mvu, ivg, svg and traer, plus the workbench
applications `todos`, `iconbrowser`, `launcher` and `mindchat`, which Phase F
migrates off it.
