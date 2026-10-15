# [Type(Scope)]: [Feature Title in English]

> **File:** `pr_stories/[NNN]-[type]-[kebab-case-name].md`
>
> **Use this template for:** Major features, multi-layer changes, architectural updates, significant refactors
>
> **Target length:** 1500–3000 words  
> **Diagrams:** 1–3 (Mermaid or ASCII)  
> **Audience:** Code reviewers, technical leads, future maintainers

---

## Executive Summary

<!-- 2-3 sentences: What is this PR about? Why does it matter? -->

Provide a high-level overview that can be read standalone. Include the problem statement, the solution provided, and the business or technical value.

**Example:**
> Implemented complete YIN pitch detection algorithm (Alain de Cheveigné & Hideki Kawahara, 2002) to enable accurate guitar tuning detection. YIN eliminates octave errors that plague FFT-only approaches, especially on low-frequency strings (E2 ≈ 82.4 Hz). Achieved ±0.5 Hz accuracy with sub-10ms latency on synthetic sine waves. Enables real-time tuner UI to display frequency, cents deviation, and string detection.

---

## Problem Statement

### Current State

What was the limitation or missing capability before?

- Example: "Frequency-to-note conversion existed but only for concert pitch (440 Hz reference). No way to detect which guitar string was being tuned."
- Example: "Audio capture from microphone was not implemented; app could not receive live input."

### Why It Matters

- User impact: (e.g., "Users can't know which string they're tuning")
- Technical debt: (e.g., "YIN algorithm is industry standard, avoiding it limits accuracy")
- Business value: (e.g., "Real-time tuning feedback is core feature")

---

## Solution Overview

### Architecture Diagram

```mermaid
graph LR
    A["🎤 Microphone Input<br/>(44.1 kHz, Float32)"] 
    B["internal/audio<br/>(malgo wrapper)"]
    C["internal/dsp/yin<br/>(YIN Algorithm)"]
    D["internal/dsp/notes<br/>(Frequency → Note)"]
    E["internal/ui<br/>(Bubbletea TUI)"]
    
    A -->|"PCM Frames"| B
    B -->|"Buffer (float32)"| C
    C -->|"Freq (Hz) + Confidence"| D
    D -->|"TunerNote (name, octave, cents)"| E
    E -->|"Render Terminal"| F["🖥️ Display"]
```

### How It Works (High Level)

1. **Audio Capture**: `malgo` reads 4096-sample frames from microphone @ 44.1 kHz
2. **YIN Algorithm**: Detects fundamental frequency via autocorrelation + absolute thresholding
3. **Note Conversion**: Maps frequency → note name, octave, cents deviation
4. **String Detection**: Matches frequency to nearest guitar string (E2–E4)
5. **UI Rendering**: Bubbletea displays gauge, tuning status, real-time feedback

---

## Design Decisions

### Why YIN Over FFT?

| Criterion | FFT | YIN |
|-----------|-----|-----|
| Octave errors | High (harmonics confuse) | Rare (autocorrelation finds period) |
| Low frequencies | Poor (need ~88 Hz for E2) | Excellent (works down to 50 Hz) |
| Polyphony | Fails on chords | Handles well |
| Speed | Very fast | Moderate (acceptable for real-time) |
| Complexity | Simple | Medium |

**Decision:** YIN chosen for accuracy on guitar fundamentals (E2–E4 range).

### Tolerance Threshold (3 cents)

- Human ear threshold: ~1 cent perceptible
- Practical tolerance: ±3 cents accepted as "in tune"
- UI feedback: Green (in tune), Yellow/Red (out of tune)

---

## Files Changed

### New Files

| File | Lines | Type | Rationale |
|------|-------|------|-----------|
| `internal/dsp/yin.go` | 450 | Implementation | Core YIN algorithm (4 phases) |
| `internal/dsp/yin_test.go` | 600 | Tests | Synthetic sine waves, edge cases |
| `.plans/03-yin-pitch-detector.md` | 250 | Documentation | Algorithm walkthrough & math |

### Modified Files

| File | Added | Removed | Type | Reason |
|------|-------|---------|------|--------|
| `internal/dsp/notes.go` | 120 | 0 | Code | Added `FrequencyToNote()`, `ClosestGuitarString()` |
| `internal/dsp/notes_test.go` | 200 | 50 | Test | Edge cases, NaN handling |
| `go.mod` | 1 | 0 | Config | No new dependencies (pure Go) |

### Unchanged Core Files

- `cmd/tuner/main.go` — Not yet integrated (Vaihe 5)
- `internal/audio/` — Stub only
- `internal/ui/` — Stub only

---

## Testing Strategy

### Unit Tests (Synthetic Signals)

```go
// Test pure sine waves at known frequencies
TestYIN_PureSineWave_440Hz()       // A4 concert pitch
TestYIN_PureSineWave_82Hz()        // E2 (guitar low string)
TestYIN_PureSineWave_329Hz()       // E4 (guitar high string)

// Test harmonically rich signals
TestYIN_WithHarmonics_E2Plus3rdHarmonic()

// Edge cases
TestYIN_EdgeCases_NaN()
TestYIN_EdgeCases_ZeroBuffer()
TestYIN_EdgeCases_InfinityFreq()
```

### Coverage

- Target: > 95% code coverage for `internal/dsp`
- Lines covered: 1,200+ / 1,280 lines (ysin + notes combined)
- Gap analysis: Only error paths with panic() remain untested

### Performance Metrics

- Speed: < 10ms per 4096-sample frame @ 44.1 kHz ✅
- Memory: ~100 KB heap allocation per frame (recycled) ✅
- Accuracy: ±0.5 Hz on pure sine waves (better than audible threshold) ✅

---

## Quality Checklist

### Code Quality

- [x] `task check` passes (tidy, vet, lint, test)
- [x] No compiler warnings or errors
- [x] Code reviewed against AGENTS.md conventions
- [x] Commit messages follow Conventional Commits
- [x] No unused imports or dead code

### Testing

- [x] All unit tests pass
- [x] Coverage > 95% for DSP modules
- [x] Edge cases handled gracefully (no panics)
- [x] Performance benchmarks acceptable

### Documentation

- [x] Inline code comments explain non-obvious logic
- [x] `.plans/03-yin-pitch-detector.md` details algorithm
- [x] Function signatures documented (Go doc comments)
- [x] This PR story complete and reviewable

### Git Hygiene

- [x] Branch: `feat/yin-pitch-detector` (correct format)
- [x] Commits: 5 commits, each atomic & logical
- [x] No merge conflicts with `main`
- [x] CI (GitHub Actions) passes

---

## How to Review

1. **Read this PR story first** — Understand the "why"
2. **Review `.plans/03-yin-pitch-detector.md`** — Algorithm walkthrough
3. **Review `internal/dsp/yin.go`** — Implementation (check math comments)
4. **Review `internal/dsp/yin_test.go`** — Test coverage
5. **Run locally:**
   ```bash
   git checkout feat/yin-pitch-detector
   task check
   go test -v ./internal/dsp/...
   ```

---

## Known Limitations & Future Work

### Current PR Scope

- ✅ YIN algorithm implemented & tested
- ✅ Note conversion & string detection
- ❌ Audio capture (next PR)
- ❌ UI / Bubbletea integration (next PR)

### Future Improvements

- [ ] Multi-pitch detection (chord recognition)
- [ ] Configurable tolerance (currently hard-coded 3 cents)
- [ ] Performance optimization (SIMD vectorization)
- [ ] Harmonic tracking for extended-range detection

---

## References & Resources

- **YIN Paper:** de Cheveigné, A., & Kawahara, H. (2002). "YIN, a fundamental frequency estimator for speech and music." *The Journal of the Acoustical Society of America, 111*(4), 1917–1930.
- **Guitar Tuning:** [Scientific pitch notation](https://en.wikipedia.org/wiki/Scientific_pitch_notation) - E2 to E4 range
- **DSP Theory:** See `.plans/02-dsp-nuottimuunnin-ja-senttilaskenta.md`

---

## Merge Instructions

Once approved:

```bash
# GitHub CLI (preferred)
gh pr merge feat/yin-pitch-detector --squash --delete-branch

# Or manually
git switch main
git pull origin main
git merge --squash feat/yin-pitch-detector
git commit -m "feat(dsp): implement yin pitch detection algorithm"
git push origin main
git branch -D feat/yin-pitch-detector
```

---

