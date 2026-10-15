# [Type]: [Short Description]

> **File:** `pr_stories/[NNN]-[type]-[kebab-case-name].md`
>
> **Use this template for:** Bug fixes, small features, single-file updates, refactoring tasks
>
> **Target length:** 500–800 words  
> **Diagrams:** 0–1 (optional)

---

## Summary

<!-- Brief 2-3 sentence overview. What was changed and why? -->

Write a concise summary of what this PR addresses. For bugfixes, state the problem and solution. For features, state the new capability added.

**Example for bugfix:**
> Fixed NaN comparison edge case in `internal/dsp/notes_test.go` where `res.Frequency != NaN` would fail (Go's NaN ≠ NaN). Added explicit `math.IsNaN()` check for proper edge-case handling.

**Example for feature:**
> Added `CentsDeviation()` helper function to `internal/dsp/notes.go` enabling precise frequency-to-cents conversion. Critical for tuner UI cents display and in-tune/out-of-tune state logic.

---

## What Changed

### Files Modified

| File | Lines Added | Lines Removed | Type | Purpose |
|------|-------------|---------------|------|---------|
| `internal/dsp/notes.go` | 5 | 1 | Fix | Corrected function call from `math.isNaN` → `math.IsNaN` |
| `internal/dsp/notes_test.go` | 10 | 3 | Test | Added NaN edge-case handling in `TestFrequencyToNote_EdgeCases` |

### Code Impact

- **Before:** NaN frequency comparison would fail silently
- **After:** Explicit NaN handling via `math.IsNaN()` ensures edge cases pass

---

## Quality Checks

- [ ] `task check` passes locally (tidy, lint, test)
- [ ] All unit tests pass for changed modules
- [ ] No new compiler warnings
- [ ] Code follows project conventions (AGENTS.md)
- [ ] Commit messages follow Conventional Commits format

---

## Testing

### Manual Testing Steps

1. Run unit tests for affected module:
   ```bash
   go test -v ./internal/dsp/...
   ```

2. Verify no regressions:
   ```bash
   task check
   ```

---

## Notes

<!-- Any additional context, design decisions, or gotchas? -->

- Edge cases with NaN/Inf now handled gracefully
- No breaking changes to public API
- Performance impact: negligible (single conditional check)

---

