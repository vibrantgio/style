# AGENTS.md — style

The MD2 type scale as a table of `textdraw.TextStyle` values — `H1` through
`H6`, `Subtitle1` and `Subtitle2`, `BodyText1` and `BodyText2`, `Button`,
`SmallButton`, `Caption` and `Overline` — plus `FontFaces()`, which returns
the five Roboto faces those styles name.

**Layer.** Tier 0 of ADR-001's table, and the one intra-tier edge that
table admits by name: `style`, `font` and `textdraw` are all tier 0, and a
tier may not otherwise import itself. ADR-003 freezes style rather than
deleting it, so `check-layers.sh` (B2.3) allows that single pair by name
instead of renumbering the tier. Its root module imports `font` and
`textdraw`. No other repository's root module imports it; outside the tier
table it is imported by the adapter modules `ivg/raster/gio`,
`svg/driver/gio` and `traer/gio`. Both directions are measured rather than
typed — `scripts/check-layers.sh --edges` reports the graph and
`scripts/sync-agents.sh` renders these sentences from it — so correcting
them here changes nothing.

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

**Frozen — do not build on this.** ADR-003 moved the typeface into the theme:
`spectrum/tokens.Typography` is a theme token carrying a full `TextStyle` per
MD3 role — typeface, weight, size, line height, tracking — plus the face
collection and a lazily built shaper. This module's MD2 scale is superseded by
it. `style` was frozen rather than deleted, so every existing import still
compiles; every exported symbol carries a `Deprecated:` marker naming its
`spectrum/tokens.Typography` replacement (C1.4). F3.3's major-bump sweep took
the alias shims out of prism and spectrum and deliberately left this module
standing — that is ADR-003's arrangement, not an oversight to finish.

One real bug was fixed rather than inherited: `H1` and `H2` were both 96 sp —
`TextStyle.Size` is `unit.Sp`, not `unit.Dp` — where MD2's H2 is 60. They
differed only in weight, so a document using both got no size hierarchy. C1.4
set `H2` to 60. Do not copy these numbers forward.

Nothing in the design system imports `style`, and that is ADR-003's point
rather than an accident. What does import it is measured and listed in the
Layer paragraph above; this note deliberately does not repeat it, because the
copy that used to live here named four workbench applications that had already
been migrated off and a demo main under mvu that never imported it at all.
