# [Type(Scope)]: [Feature Title with Visual Focus]

> **File:** `pr_stories/[NNN]-[type]-[kebab-case-name].md`
>
> **Use this template for:** Data pipelines, UI flows, architectural refactors
>
> **Target length:** 800–1500 words  
> **Diagrams:** 3–5 (Mermaid, ASCII art)

---

## Overview

One sentence summary of what this PR does visually.

---

## Data Flow Visualization

### Before

```
[Source] ─→ [Dead End]
```

### After

```
[Source] ─→ [Step 1] ─→ [Step 2] ─→ [Result]
```

---

## Architecture Diagram

```mermaid
graph TD
    A["Input"]
    B["Process"]
    C["Output"]
    A --> B --> C
```

---

## Component Interaction

```mermaid
graph TB
    subgraph Layer1["Layer 1"]
        A["Component A"]
    end
    subgraph Layer2["Layer 2"]
        B["Component B"]
    end
    A --> B
```

---

## Test Coverage Matrix

| Module | Coverage | Status |
|--------|----------|--------|
| Module A | 95% | ✅ Complete |
| Module B | 85% | 🟡 Partial |

---

## Files Changed

| File | Type | Purpose |
|------|------|---------|
| file1.go | Feature | Main implementation |
| file2_test.go | Test | Test suite |

---

## Quality Checklist

- [x] `task check` passes
- [x] All tests passing
- [x] Coverage > 90%

---

