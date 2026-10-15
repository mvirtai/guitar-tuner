# Quick Start — Workflow käyttöönotto

Tervetuloa Guitar Tuner -kehitysworkflowiin! Tämä dokumentti opastaa sinut nopeasti.

## TL;DR — 5 minuuttia

```bash
# 1. Lue suunnitelma
cat .plans/03-yin-pitch-detector.md

# 2. Luo haara
task git:new-branch FEAT_TYPE=feat BRANCH_NAME=yin-pitch-detector

# 3. Kirjoita koodi (testi ensin, sitten toteutus)
# - Avaa internal/dsp/yin_test.go
# - Kirjoita testi synteettiselle siniaallille
# - Suorita: task check
# - Kirjoita koodi kunnes testit menevät läpi

# 4. Committaa logiikka-yksikkö kerrallaan
task git:commit TYPE=feat SCOPE=dsp MSG="add yin difference function" FILES="internal/dsp/yin.go internal/dsp/yin_test.go"

# 5. Kirjoita PR-story
cp .plans/templates/PR_STORY_EXTENDED.template.md pr_stories/001-feat-dsp-yin-pitch-detector.md
# Täytä pohja merkityksellisillä tiedoilla

# 6. Luo PR
task git:pr FILE=pr_stories/001-feat-dsp-yin-pitch-detector.md
```

---

## Workflow — Vaihe Kerrallaan

### Vaihe 1: Suunnitelma (30 min)

**Mitä tehdään:**
- Luetaan `.plans/XX-nimi.md` suunnitelma
- Ymmärretään tavoite, arkkitehtuuri, testit
- Hyväksytään ennen etenemistä

**Komennot:**
```bash
task log  # Avaa lokikirja, lisää viite suunnitelmaan
```

**Tuottaa:**
- Selvä kuva siitä mitä tehdään ja miksi

---

### Vaihe 2: Haara & Setup (5 min)

**Mitä tehdään:**
- Luodaan uusi git-haara suunnitelmasta
- Varmistetaan että olet oikeassa haarassa

**Komennot:**
```bash
task git:new-branch FEAT_TYPE=feat BRANCH_NAME=yin-pitch-detector
git branch --show-current  # Varmista: feat/yin-pitch-detector
```

**Tuottaa:**
- Uusi, puhdas haara `origin/main` pohjalta

---

### Vaihe 3: Toteutus & Testaus (2–4 tuntia)

**Mitä tehdään:**
- Luetaan suunnitelmasta step-by-step -ohjeet
- Kirjoitetaan testi jokaiselle vaiheelle (TDD)
- Kirjoitetaan koodi kunnes testit menevät läpi
- Testataan säännöllisesti

**Komennot:**
```bash
# Aja testit jokaisen muutoksen jälkeen
go test -v ./internal/dsp/...

# Täydellinen laadunvarmistus
task check  # Tidy, lint, test

# Tarkista status
git status
```

**Tuottaa:**
- Valmis, testattu koodi
- Kaikki testit menevät läpi (`task check` vihreä)

---

### Vaihe 4: Laadunvarmistus & Commitointi (30 min)

**Mitä tehdään:**
- Varmistetaan että `task check` tulee vihreäksi
- Committaan logiikka-yksikkö kerrallaan
- Tarkistetaan git-status

**Komennot:**
```bash
# Varmista että kaikki vihreä
task check

# Committaa yksi logiikka-yksikkö
task git:commit TYPE=feat SCOPE=dsp MSG="add yin difference function" FILES="internal/dsp/yin.go internal/dsp/yin_test.go"

# Tarkista status
git status  # Ei uncommitted changes
```

**Tuottaa:**
- Atomisia, hyviä committeja
- Git-historia on ymmärrettävä

---

### Vaihe 5: PR-Story & GitHub PR (1 tuntia)

**Mitä tehdään:**
- Valitaan sopiva PR-story-template
- Kirjoitetaan PR-story merkityksellisesti
- Luodaan GitHub PR automaattisesti

**Komennot:**
```bash
# Valitse template (COMPACT / EXTENDED / VISUAL)
cp .plans/templates/PR_STORY_EXTENDED.template.md pr_stories/001-feat-dsp-yin-pitch-detector.md

# Avaa editorissa ja täytä:
nano pr_stories/001-feat-dsp-yin-pitch-detector.md

# Luo PR
task git:pr FILE=pr_stories/001-feat-dsp-yin-pitch-detector.md
```

**Tuottaa:**
- GitHub PR luotu
- PR-body sisältää merkityksellisen PR-storyn
- CI/CD testit ajetaan automaattisesti

---

## Template Valinta

### COMPACT

Käytä kun:
- Bugfix tai pieni feature (< 500 rivin muutos)
- 1–3 tiedostoa vaikuttunut
- 1–2 kerrosta projektissa

Esimerkki: NaN-vertailun bugfix

### EXTENDED

Käytä kun:
- Suuri feature tai multi-layer muutos (> 500 rivin muutos)
- > 3 tiedostoa vaikuttunut
- Arkkitehtuurimuutos

Esimerkki: YIN-algoritmin implementointi

### VISUAL

Käytä kun:
- Paljon kaavioita ja diagrammeja tarvitaan
- Data pipeline tai dataflow muutos
- UI/workflow uudelleenjärjestely

Esimerkki: Audio pipeline arkkitehtuuri

---

## Komennot Yhteenveto

```bash
# 📋 Dokumentaatio
task log              # Avaa lokikirja editorissa
task log:status       # Näytä projektin nykytila

# 🌳 Git & Haara
task git:new-branch FEAT_TYPE=feat BRANCH_NAME=name
task git:commit TYPE=feat SCOPE=dsp MSG="message" FILES="file1.go file2.go"
task git:diff         # Näytä erot main vastaan
task git:status       # GitHub PR status

# ✅ Laadunvarmistus
task check            # Tidy + Lint + Test (KAIKKI)
task tidy             # go mod tidy
task lint             # go vet + golangci-lint
task test             # go test -v ./...
task test:cov         # go test + coverage

# 📝 PR & GitHub
task git:pr FILE=pr_stories/001-...md
task git:pr-body FILE=pr_stories/001-...md PR=1
task git:merge        # Squash & merge PR

# 🔖 Versio
task version          # Näytä versio
task version:bump PART=patch|minor|major

```

---

## Esimerkkisessio

Oletetaan että toteutat YIN-algoritmin:

```bash
# 1️⃣ Lue suunnitelma
cat .plans/03-yin-pitch-detector.md
# ✅ Ymmärrät: 4 vaihe, synteettinen testaus, math

# 2️⃣ Luon haaran
task git:new-branch FEAT_TYPE=feat BRANCH_NAME=yin-pitch-detector
# ✅ Nyt olet: feat/yin-pitch-detector

# 3️⃣ Implementoin Vaihe 1: Difference Function
# Avaan internal/dsp/yin_test.go
# Kirjoitan: func TestYIN_DifferenceFunc_PureSine_440Hz(t *testing.T) { ... }
# Kirjoitan: func DifferenceFunc(buffer []float32, maxLag int) []float32 { ... }
# Suorita: go test -v ./internal/dsp/...
# ✅ Testi menee läpi

# 4️⃣ Committain
task git:commit TYPE=feat SCOPE=dsp MSG="add yin difference function" FILES="internal/dsp/yin.go internal/dsp/yin_test.go"
# ✅ Commit muodossa: feat(dsp): add yin difference function

# 5️⃣ Implementoin Vaihe 2: Cumulative Mean Normalization
# ... (samat askeleet)
task git:commit TYPE=feat SCOPE=dsp MSG="add yin cumulative mean normalization" FILES="internal/dsp/yin.go internal/dsp/yin_test.go"

# 6️⃣ Jatkan Vaihe 3–4 samalla tavalla

# 7️⃣ Kaikki valmis, laadunvarmistus
task check
# ✅ Kaikki vihreä

# 8️⃣ Kirjoita PR-story
cp .plans/templates/PR_STORY_EXTENDED.template.md pr_stories/001-feat-dsp-yin-pitch-detector.md
nano pr_stories/001-feat-dsp-yin-pitch-detector.md
# Täytä: Executive Summary, Problem, Architecture, Files Changed, Testing, Checklist

# 9️⃣ Luo GitHub PR
task git:pr FILE=pr_stories/001-feat-dsp-yin-pitch-detector.md
# ✅ PR luotu, CI/CD ajaa testit automaattisesti

# 🔟 Odota reviewia, korjaa jos tarvii, merger kun hyväksytty
task git:merge
# ✅ PR merged, haara poistettu
```

---

## Troubleshooting

### ❓ "task check" epäonnistuu

1. Lue virheviesti tarkkaan
2. Korjaa koodi
3. Aja uudelleen: `task check`
4. Vasta sitten commitoi

### ❓ Commitin viesti väärin

Huono commit:
```
WIP: lol fixed stuff
```

Hyvä commit:
```
feat(dsp): add yin difference function
```

Käytä: `task git:commit TYPE=feat SCOPE=dsp MSG="add yin difference function" FILES="..."`

### ❓ PR-story on liian lyhyt / pitkä

- **COMPACT:** 500–800 sanaa
- **EXTENDED:** 1500–3000 sanaa
- **VISUAL:** 800–1500 sanaa (+ diagrammit)

Valitse oikea template alusta alkaen!

### ❓ Haara ei ole origin/main:issa

```bash
task git:post-merge-branch  # Resetoi main ja poista haara
```

---

## Best Practices

1. **Committaa usein** — jokainen logiikka-yksikkö = commit
2. **Testaa ensin (TDD)** — testi ennen koodia
3. **Pyydä apua lokiin** — `task log` päivitykset
4. **Lue suunnitelma huolella** — sääli tuhlata aikaa väärään suuntaan
5. **PR-story on tärkeä** — vie aikaa sen kirjoittamiseen
6. **Vältä suuria PRjä** — max 500–1000 rivin PRt parempia

---

## Seuraavaksi

Kun olet valmis ensimmäiseen PRhun:

```bash
task log  # Päivitä projektiloki
# Merkitse: "Sessio 3 — YIN-algoritmi valmis, PR #1 luotu"

# Siirry seuraavaan vaiheeseen
# Vaihe 3: internal/audio capture
```

---

