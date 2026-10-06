# Custom Agent Rules & Guidelines (AGENTS.md)

This document provides a comprehensive, unified collection of all operational conventions, communication protocols, architectural standards, and development workflows for the `guitar_tuner` codebase.

> **MANDATORY — READ, UNDERSTAND, AND ACKNOWLEDGE THE LANGUAGE POLICY BEFORE EVERY TASK**
>
> Every agent working in this repository, including the main agent and any delegated or custom agent, MUST read and follow the language policy in Section 1 before inspecting or changing files. At the start of each task, the agent MUST explicitly acknowledge in its first user-visible progress message: **"Luin ja ymmärsin kieliasetukset: suunnitelmat, niiden pohjat, sisäiset dokumentaatiot ja muistiot suomeksi; PR storyt englanniksi; koodi ja kommentit englanniksi; git-haarat ja muu git-/projektin metatieto englanniksi."**

---

## 1. Communication & Document Language Policy

* **Finnish (Suomi)**:
  * Used for all direct conversational chat interactions between the AI agent and the developer.
  * Used for all plans and plan templates, internal documentation, memos, and instruction documents (`.plans/`).
  * Used for step-by-step mentoring guides, tutorial walkthroughs, and architectural explanations directed to the developer.

* **English**:
  * Used for all source code (Go).
  * Used for all code comments, commit messages, branch names, and other Git or project metadata.
  * Used for all Pull Request stories located in `pr_stories/`.
  * All English content must maintain senior-level software engineering terminology without superficial embellishments.

* **Markdown Formatting Quality (MD032 & Lint Compliance)**:
  * All markdown documents (`.plans/`, `pr_stories/`, rules, and reviews) must strictly adhere to standard markdownlint rules (`.markdownlint.json`).
  * Always place proper blank lines before and after lists (MD032), headings, code blocks, and blockquotes.
  * Ensure all markdown links use valid paths without broken syntax.

* **Internal Documentation vs. Commits**:
  * Files in `.plans/` and local notes are developer reference materials.
  * The documentation files committed to git and tracked in PRs are Pull Request stories (`pr_stories/`).

---

## 2. Pair Programming & Mentoring Model

* **Developer Writes Code & Code Ownership (Feature Implementations)**:
  * By default, the developer writes the code directly from the design plan (`.plans/`) to learn the architecture, DSP math, and audio stream programming.
  * The agent does **NOT** write, autocomplete, or generate new feature source code files ahead of the developer unless explicitly requested.

* **Agent Scoped Bugfixing & Debugging ("Explain and fix with AI")**:
  * When the developer asks for help or shares error logs/compiler errors, the agent fixes the relevant errors, typos, missing imports, and syntax issues.
  * **WIP Boundary Rule**: The agent MUST preserve developer work-in-progress and never jump ahead by writing unrequested future sections.

* **Session Progress & Status Briefing Protocol**:
  * Whenever the agent completes an edit, debugging step, or verification run, it must provide a structured 4-point briefing:
    1. **What Changed / Resolved**: Specific files modified, key diffs, and exact line references.
    2. **Quality Gates Status**: Results of executed checks (`task check`, `task test`, `task lint`).
    3. **Version Bump Reminder**: Proactively assess if milestone changes warrant a semantic version bump (`task version:bump PART=patch|minor`).
    4. **Next Recommended Action**: The immediate next command or coding step for the developer.

* **Educational Approach**:
  * Emphasize *why* specific patterns, interfaces, mathematical algorithms (YIN, autocorrelation, parabolic interpolation) or idioms are chosen and how they fit into the overall application architecture.

---

## 3. Git Workflow, Branching & Taskfile Automation

* **Branch Naming Conventions**:
  * Topic branches must always be named strictly in lowercase: `<type>/<kebab-case-description>`.
  * Allowed types: `feat`, `fix`, `ci`, `docs`, `refactor`, `chore`, `test`, `perf`.
  * *Examples:* `feat/yin-pitch-detector`, `feat/audio-stream-capture`, `feat/bubbletea-ui`.

* **PR Titles & Commit Messages (Conventional Commits)**:
  * Follow Conventional Commits format: `<type>(<scope>): <lowercase description in English>`.
  * *Examples:* `feat(dsp): implement yin pitch detection algorithm`, `test(dsp): add synthetic sine wave tests`.

* **Taskfile Automation Commands**:
  * `task check`: Runs full project verification (`tidy`, `lint`, and `test`).
  * `task test`: Runs all Go unit tests.
  * `task test:cov`: Runs tests with coverage profile.
  * `task version`: Prints current project version from `VERSION`.
  * `task version:bump PART=patch|minor|major`: Bumps semantic version.
  * `task git:new-branch FEAT_TYPE=... BRANCH_NAME=...`: Creates topic branch.
  * `task git:commit TYPE=... SCOPE=... MSG="..." FILES="..."`: Formatted commit.
  * `task git:pr FILE=<file.md> TITLE="..."`: Runs quality gates and opens GitHub PR via `gh pr create`.
  * `task git:merge`: Developer-driven squash & merge via GitHub CLI.
  * `task git:post-merge-branch`: Cleans local branch, resets `main` against `origin/main`.

---

## 4. Pull Request Stories & Fresh-Eyes Audit (`pr_stories/`)

* **File Naming & Location**:
  * Saved under `pr_stories/` using sequential numeric prefixes: `pr_stories/<seq>-<type>-<kebab-case-description>.md` (e.g., `pr_stories/001-feat-dsp-yin-pitch-detection.md`).

* **Templates**:
  * `pr_stories/templates/PR_STORY_COMPACT.template.md` (for lean operational changes, single bugfixes).
  * `pr_stories/templates/PR_STORY_EXTENDED.template.md` (for major features and multi-layer milestones).

* **The 4 Pillars of PR Story Audit**:
  1. **Strict Factual Accuracy (Zero Fake Claims)**: Claims must match `git diff main...HEAD`.
  2. **Purposeful Visualizations (~0–3 Diagrams)**: Mermaid diagrams clarifying architecture, audio dataflow, or state transitions.
  3. **Professional Software Engineering Rigor**: Clear rationale for design decisions, exact Files Changed tables.
  4. **Markdown & Link Quality Gates**: Strict markdownlint compliance (MD032: blank lines around lists, headings, and code fences).

---

## 5. Go Audio & DSP Architecture Guidelines

* **Domain Separation**:
  * `internal/dsp`: Pure Go mathematical algorithms (YIN, note conversion, cents calculation). **Zero CGo or hardware dependencies**, enabling deterministic and lightning-fast unit tests with synthetic sine waves.
  * `internal/audio`: Hardware abstraction and PCM audio streaming via `malgo` (miniaudio Go wrapper). Buffers audio samples and feeds them downstream to DSP via Go channels.
  * `internal/ui`: Bubbletea Model-Update-View architecture and Lipgloss styling. Renders real-time tuner gauge, cents offset, frequency, and string detection.
  * `cmd/tuner`: Main application entry point orchestrating audio stream initialization, DSP processing, and Bubbletea program execution.

* **Error Handling & Zero Panic**:
  * Gracefully handle missing or inaccessible microphone input devices with clear user error messages.
  * Zero-tolerance for unhandled errors.
