# Example sheets

Each example is a single `.tsvt` spreadsheet — a TAB-separated grid whose cells are literal values or `=formulas` that address other cells in A1 notation (`B2`, `D2:D5`). Open one in the browser editor and edit any cell live (edits recompute through the same engine):

```sh
tsvsheet serve examples/grades.tsvt
# then open http://127.0.0.1:8080
```

Every example also renders straight to stdout — handy for a terminal demo or piping into other tools:

```sh
tsvsheet render examples/invoice.tsvt | column -t
```

| Sheet | Demonstrates |
| --- | --- |
| [grades](grades.tsvt) | Per-row aggregates (`round(avg(B2:D2), 1)`) and a conditional flag (`if(E2 >= 70, 1, 0)`) that reads the average computed earlier in the same row. |
| [invoice](invoice.tsvt) | Per-row arithmetic (`Amount = Qty × Price`, `=B2*C2`) and a `Total` row summing the amount column over a range (`=sum(D2:D5)`). |
| [math](math.tsvt) | Error-value propagation: dividing by a zero denominator yields `#DIV/0!`, which flows through any expression that reads the cell. |

## A note on the language

A `.tsvt` **is** the spreadsheet: there is no separate data file. Each cell is a literal value, or — when it begins with `=` — a formula over the expression sublanguage (arithmetic, comparisons, and the builtins `sum`, `avg`, `min`, `max`, `count`, `round`, `abs`, `len`, `concat`, `if`). Formulas reference other cells by A1 address, exactly like a conventional spreadsheet; a reference off the grid resolves to `#REF!`, a cycle to `#CIRC!`.

Two things worth knowing when you edit these:

- References are A1-absolute (`B2`, `$B$2`, ranges `D2:D5`). A reference that leaves the grid is `#REF!` — the intended out-of-grid result, not a bug.
- Text results are not yet supported: a quoted string in a formula (`"Pass"`) is parsed as a named-column reference and currently resolves to `#REF!`, so these examples use numeric results. This is a known gap tracked against the A1 model, to be closed when the reference/string semantics are folded back into the grammar.

The full language is specified in [uplang/tsvsheet](https://github.com/uplang/tsvsheet).
