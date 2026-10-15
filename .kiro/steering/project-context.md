---
inclusion: always
---

# Guitar Tuner — Projektin konteksti & Workflow

Tämä steering-tiedosto latautuu automaattisesti jokaiseen Kiro-sessioon.

## 🚀 Workflow — 5 vaihetta

```
Vaihe 1: SUUNNITELMA          (.plans/XX-nimi.md)
   ↓
Vaihe 2: HAARA & SETUP        (task git:new-branch)
   ↓
Vaihe 3: TOTEUTUS & TESTING   (TDD, task check)
   ↓
Vaihe 4: COMMITOINTI          (task git:commit)
   ↓
Vaihe 5: PR-STORY & PR        (task git:pr)
```

**Quick Start:**
```bash
task workflow       # Näytä workflow quick reference
task log            # Avaa projektiloki
cat .plans/QUICK_START.md  # Lue ohjeet
```

---

## 📂 Tärkeimmät tiedostot

| Tiedosto | Tarkoitus |
| --- | --- |
| `.plans/WORKFLOW.md` | Täydellinen workflow-dokumentaatio (5 vaihe) |
| `.plans/QUICK_START.md` | 5 minuuttia, käytännöllinen opas |
| `.plans/PROJECTLOG.md` | Elävä sessioloki ja projektitila |
| `.plans/01-...-arkkitehtuuri.md` | Pääsuunnitelma |
| `.plans/02-...-dsp-nuottimuunnin.md` | DSP Vaihe 1 opetusdokumentti |
| `Taskfile.yml` | Kaikki automaatiot (workflow, git, build) |

---

## 🎯 Nykytila (2026-10-15)

- **Versio:** 0.1.0
- **Aktiivinen haara:** `feat/dsp-note-converter`
- **Projektitila:**
  - ✅ Vaihe 0: Infrastruktuuri valmis
  - ✅ Vaihe 1: DSP Nuottilaskenta valmis
  - ❌ Vaihe 2: YIN-algoritmi edessä
  - ❌ Vaihe 3: Audio capture edessä
  - ❌ Vaihe 4: TUI edessä
  - ❌ Vaihe 5: Integrointi & shipping

---

## 📋 PR-Story Templatet

Valitse sopiva, kun kirjoitat PR:n:

1. **COMPACT** — Pienet bugfixet, yksittäiset muutokset
   - Pituus: 500–800 sanaa
   - Diagrammit: 0–1
   - Tiedosto: `.plans/templates/PR_STORY_COMPACT.template.md`

2. **EXTENDED** — Suuret featuret, multi-layer arkkitehtuuri
   - Pituus: 1500–3000 sanaa
   - Diagrammit: 1–3
   - Tiedosto: `.plans/templates/PR_STORY_EXTENDED.template.md`

3. **VISUAL** — Data pipelines, UI flows, architectural changes
   - Pituus: 800–1500 sanaa
   - Diagrammit: 3–5
   - Tiedosto: `.plans/templates/PR_STORY_VISUAL.template.md`

---

## 🔧 Workflow-komennot

```bash
# Dokumentaatio
task workflow       # Näytä quick reference
task log            # Avaa lokikirja
task log:status     # Näytä projektin tila

# Git & Haara
task git:new-branch FEAT_TYPE=feat BRANCH_NAME=name
task git:commit TYPE=feat SCOPE=dsp MSG="message" FILES="file.go"
task git:pr FILE=pr_stories/001-...md
task git:merge      # Squash & merge PR

# Laadunvarmistus
task check          # Tidy + Lint + Test (KAIKKI)
task test           # go test -v ./...

# PR-templatet
task pr:template    # Näytä template vaihtoehdot
```

---

## 🌳 Haaran nimeäminen

Muoto: `<type>/<kebab-case-description>`

**Esimerkit:**
- `feat/yin-pitch-detector`
- `fix/nan-comparison-edge-case`
- `docs/dsp-algorithm-explanation`
- `test/add-synthetic-sine-tests`

**Sallitut TYPE:t:**
- `feat` — uusi ominaisuus
- `fix` — bugfix
- `test` — testien lisäys
- `docs` — dokumentaatio
- `refactor` — koodi-uudelleenjärjestely
- `ci` — CI/CD muutokset

---

## 📝 Conventional Commits

Commitit noudattavat muotoa: `type(scope): message`

**Esimerkit:**
```
feat(dsp): implement yin pitch detection algorithm
fix(dsp): handle nan frequency comparison correctly
test(dsp): add synthetic sine wave tests
docs(workflow): add pr-story templates
```

Komento automatisoi tämän:
```bash
task git:commit TYPE=feat SCOPE=dsp MSG="implement yin algorithm" FILES="yin.go yin_test.go"
```

---

## 💡 Lokitukseen liittyvät komennot

Jokaisen session lopussa:

```bash
task log  # Avaa lokikirja editorissa, lisää merkintä:
# ### YYYY-MM-DD — Sessio N: Lyhyt kuvaus
# **Tehtävä:** ...
# **Mitä tehtiin:** ...
# **Löydökset:** ...
# **Päätökset:** ...
# **Versio:** x.y.z
```

Pohja: `.plans/templates/SESSION_LOG.template.md`

---

## ✅ Tarkistuslista ennen PR:n tekemistä

1. [ ] `task check` tulee vihreäksi
2. [ ] Kaikki testit menevät läpi
3. [ ] Coverage > 90% DSP-moduuleille
4. [ ] Git-status on puhdas (ei uncommitted changes)
5. [ ] Commitit noudattavat Conventional Commits -muotoa
6. [ ] Suunnitelma (.plans/XX-...) päivitetty (jos tarvii)
7. [ ] PR-story kirjoitettu oikealla templatella
8. [ ] Branch on ajantasalla `origin/main` kanssa

---

## 🎓 Oppimateriaalit

- `.plans/02-dsp-nuottimuunnin-ja-senttilaskenta.md` — DSP matematiikka
- `docs/theory/` — Syvällinen teoriatieto
- `.agents/AGENTS.md` — Projektisäännöt, konventiot

---

## 📞 Workflow-apua

```bash
task workflow           # Quick reference
cat .plans/WORKFLOW.md  # Täydellinen dokumentaatio
cat .plans/QUICK_START.md  # Käytännöllinen opas
task log                # Avaa lokikirja
```

---

