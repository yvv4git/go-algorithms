# STYLE_GUIDE

## README.md — Task File Format

Every task must include a `README.md` file with
the following sections in order:

1. `# [number]. [Task Name]`
2. `## Info` — brief description in Russian:
   core idea, key solution approach, final complexity.
3. `## Level - [easy|medium|hard]`
4. `## Task` — original problem statement in English.
5. `## Объяснение` — detailed explanation in Russian.
6. `## Example 1`, `## Example 2`, ... — examples
   with `Input:` and `Output:` blocks.
7. `## Constraints` — input constraints.
8. `## См. также` — links to approach files.

### README.md Rules

1. **Heading** — `# [number]. [Name]`
   (dot after the number, name in English).
2. **Info** — 2–3 sentences: task summary,
   key idea of the solution, final complexity.
3. **Level** — allowed values:
   `easy`, `medium`, `hard`.
4. **Task** — original problem statement
   in English without modifications.
5. **Объяснение** — detailed Russian explanation:
   what we do, why the approach works,
   list of approaches with complexity.
6. **Examples** — numbered starting from 1.
   Use fenced code blocks with `text` language.
   Include `Input:` and `Output:` lines.
   Optional `Explanation:` line.
7. **Constraints** — same as in the original source.
8. **См. также** — links to approach files
   using relative paths.

### Linting

Before committing, validate every `README.md`:

```bash
markdownlint README.md
```

Fix all reported issues before committing.
If markdownlint is not installed:

```bash
npm install -g markdownlint-cli
```

---

## Solution Files — Naming and Format

### File Naming

Solution files follow this pattern:

```text
ver[N]_[approach].go
```

Where:

- `[N]` — sequential approach number
  (1, 2, 3, ...), starting from the simplest.
- `[approach]` — short description of the approach
  in English, snake_case.

Examples:

- `ver1_bruteforce.go`
- `ver2_dp.go`
- `ver3_divide_and_conquer.go`
- `ver4_two_pointers.go`

Tests follow the same pattern with `_test` suffix:

```text
ver[N]_[approach]_test.go
```

### Solution File Structure (ver*.go)

Each solution file must contain:

- `package main`
- Function with comment block before it

The comment block must include:

- `TASK:` — task description in Russian (1–3 sentences).
- `METHOD:` — approach name and idea description.
- `TIME COMPLEXITY:` — time complexity with justification.
- `SPACE COMPLEXITY:` — space complexity with justification.

Function naming:

- ver1 uses base name from the problem
  (e.g., `maxCoins`).
- Subsequent versions append suffix
  `V2`, `V3`, etc. (e.g., `maxCoinsV2`).

### Solution File Rules

1. **Package** — always `package main`.
2. **Function name** — ver1 uses base name,
   subsequent versions add `V2`, `V3`, etc.
3. **Comment block** — placed directly before
   the function body, contains TASK, METHOD,
   TIME COMPLEXITY, SPACE COMPLEXITY sections.
4. **Inline comments** — brief, explain non-trivial
   steps. Do not restate what is obvious from code.
5. **Version order** — from simplest (bruteforce)
   to optimal. First version is always naive.

### Test File Structure (ver*_test.go)

Each test file must contain:

- `package main`
- `import "testing"`
- Table-driven test function

Test function naming: `Test_` + solution name
(e.g., `Test_maxCoins`, `Test_maxCoinsV2`).

### Test File Rules

1. **Test function name** — `Test_` + solution name.
2. **Table-driven** — all tests use a slice of structs
   with a `name` field.
3. **Required test cases**:
   - Examples from the task (Example 1, Example 2, ...).
   - Edge cases: empty input, single element,
     maximum size.
   - Special cases: all zeros, negative values
     (if allowed), duplicate elements.
4. **Comparison** — table-driven style:
   `if got := Function(tt.input); got != tt.want`.
