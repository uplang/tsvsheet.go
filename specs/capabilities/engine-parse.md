# engine-parse — typed AST over the generated parser

## Goal

Turn raw `.tsvt` bytes into an immutable, typed AST (or a precise syntax error), hiding every ANTLR type from the rest of the program.

## Requirements

- R1: `tsvt.Parse` accepts the full grammar: section markers (`=header(n)`, `=body`, `=final`), structural commands (`=A<`), and cell rows in all three equivalent §7 forms (the three `testdata/form*.tsvt` files from the grammar repo parse to equivalent addressed meaning).
- R2: The complete §5 reference algebra is represented: letter/numeric/named/last columns, relative/after/absolute/last/wildcard rows, ranges, matrices, grouped ranges — and §6 modifiers.
- R3: The §11 expression sublanguage is represented with the specified precedence (grouping, unary sign, `* / %`, `+ -`, comparisons) and function calls with case-preserved names.
- R4: A syntax error surfaces as `constants.ErrSyntax` (matchable with `errors.Is`) carrying line, column, and offending text; nothing is printed to stderr by the parser (custom `antlr.ErrorListener`).
- R5: No ANTLR type appears in any exported-or-shared signature outside `internal/tsvt`; the generated package is imported by `internal/tsvt` only.

## Acceptance Criteria

- All four grammar-repo `testdata/*.tsvt` files parse; the three forms yield ASTs that normalize to the same placements.
- Malformed input (unbalanced parens, bad numeric ref, stray modifier) yields `ErrSyntax` with the correct line/column, asserted by tests.
- The wrapper package holds 100% statement coverage including every error path.
