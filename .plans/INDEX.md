# Guitar Tuner — Dokumentaation sisällysluettelo

Luo yhteenveto kaikista `.plans/`-tiedostoista ja mistä löydät mitä.

## 📚 Dokumentaation rakenne

### Workflow & Projektinhallinta

| Tiedosto | Tarkoitus | Lukija |
| --- | --- | --- |
| **WORKFLOW.md** | Täydellinen 5-vaiheen workflow (suunnitelma → PR) | Jokainen kehittäjä — ensimmäisellä kerralla |
| **QUICK_START.md** | Nopea opas (5 minuuttia), käytännölliset esimerkit | Kehittäjät ennen jokaista sessio |
| **PROJECTLOG.md** | Elävä lokikirja, tila, sessiolokit | Agentti päivittää, jokainen lukee |
| **INDEX.md** | Tämä tiedosto — missä löydät mitä | Ensimmäisellä kerralla |

### Suunnitelmat (`.plans/XX-...`)

| Tiedosto | Vaihe | Sisältö | Lukija |
| --- | --- | --- | --- |
| **01-kitaranviritin-arkkitehtuuri-ja-toteutus.md** | 0 | Pääarkkitehtuuri, koko projektin visio | Arkkitehti, lead dev |
| **02-dsp-nuottimuuttaja-ja-senttilaskenta.md** | 1 | DSP matematiikka, frequency-to-note, testaus | DSP-kehittäjä |
| **03-yin-pitch-detector.md** | 2 (tulossa) | YIN-algoritmi 4 vaiheessa, synthetic tests | YIN-kehittäjä |
| **04-audio-capture.md** | 3 (tulossa) | malgo-integraatio, PCM streaming | Audio-kehittäjä |
| **05-tui-bubbletea.md** | 4 (tulossa) | Bubbletea model, Lipgloss styling, gauge | UI-kehittäjä |

### Templateit (`.plans/templates/`)

| Tiedosto | Käyttö | Milloin |
| --- | --- | --- |
| **PLAN_TEMPLATE.md** | Suunnitelman pohja | Kun alustat uuden Vaihe |
| **PR_STORY_COMPACT.template.md** | Pieni PR (~500–800 sana) | Bugfixet, yksittäiset muutokset |
| **PR_STORY_EXTENDED.template.md** | Iso PR (~1500–3000 sana) | Suuret featuret, multi-layer |
| **PR_STORY_VISUAL.template.md** | Diagrammipainottuinen PR (~800–1500 sana) | Data pipelines, UI flows |
| **SESSION_LOG.template.md** | Sessio-merkinnän pohja | Jokaisen session lopussa |
| **TECHNICAL_DESIGN.template.md** | Teknisen designin pohja | Design-ensin vaiheiden varten |

### Teoriaa (`.plans/docs/theory/`)

| Tiedosto | Aihe | Tausta |
| --- | --- | --- |
| **01-music-acoustics-and-dsp.md** (docs/theory/) | Musiikkiteoria, akustiikka, DSP-perusteet | Syvällinen matematiikka |
| **02-musical-intervals-and-tuning-systems.md** (docs/theory/) | 12-TET, senttit, viritysjärjestelmät | Musiikallinen teoria |

---

## 🗺️ Navigointi: Mitä luet milloin?

### 🆕 Ensimmäinen sessio projektissa?

1. Lue: **QUICK_START.md** (5 minuuttia)
2. Lue: **WORKFLOW.md** (20 minuuttia)
3. Lue: **01-kitaranviritin-arkkitehtuuri-ja-toteutus.md** (30 minuuttia)
4. Aloita: Vaihe 1 (suunnitelma)

### 📋 Ennen seuraavaa kehitysvaihetta?

1. Lue: Sopiva `.plans/XX-...` suunnitelma
2. Lue: **QUICK_START.md** (muistutukseksi)
3. Aloita: Vaihe 2 (haara)

### ✅ Ennen PR:n tekemistä?

1. Katso: **QUICK_START.md** Vaihe 5 (PR-story & PR)
2. Valitse: Sopiva template (`COMPACT` / `EXTENDED` / `VISUAL`)
3. Kirjoita: PR-story

### 🐛 Kun törmää ongelmaan?

1. Tarkista: **PROJECTLOG.md** (samojen ongelmien aikaisemmat ratkaisut?)
2. Tarkista: **AGENTS.md** (konventiot, säännöt)
3. Tarkista: Sopivan suunnitelman ohjeet

---

## 🎯 Workflow Vaihe-by-Vaihe

```
Vaihe 1: SUUNNITELMA
└─ Lue: Sopiva .plans/XX-...md
└─ Hyväksy: Tavoite, arkkitehtuuri, testit

Vaihe 2: HAARA & SETUP
└─ Komento: task git:new-branch FEAT_TYPE=feat BRANCH_NAME=...

Vaihe 3: TOTEUTUS & TESTING
└─ Seuraa: .plans/XX-...md "Vaiheittainen toteutusopas"
└─ Komento: task check (jokaisen muutoksen jälkeen)

Vaihe 4: COMMITOINTI
└─ Komento: task git:commit TYPE=... SCOPE=... MSG=... FILES=...

Vaihe 5: PR-STORY & PR
└─ Valitse: COMPACT / EXTENDED / VISUAL template
└─ Kirjoita: pr_stories/NNN-...md
└─ Komento: task git:pr FILE=pr_stories/NNN-...md
```

---

## 🔍 Nopea haku: Etsi vastaus

| Kysymys | Vastaus löytyy |
| --- | --- |
| "Miten aloitan uuden feature?" | QUICK_START.md, Vaihe 1–2 |
| "Mitä komentoja käytän?" | Taskfile.yml, task workflow |
| "Kuinka kirjoitan PR:n?" | WORKFLOW.md Vaihe 5, PR_STORY_*.template.md |
| "Miksi projekti on järjestetty näin?" | 01-kitaranviritin-arkkitehtuuri-ja-toteutus.md |
| "Miten DSP-matematiikka toimii?" | 02-dsp-nuottimuuttaja-ja-senttilaskenta.md |
| "Mitkä konventiot pätee?" | .agents/AGENTS.md |
| "Missä vaiheessa ollaan?" | .plans/PROJECTLOG.md (Projektin tila -taulukko) |
| "Mitä minun pitää tehdä seuraavaksi?" | task log:status, .plans/PROJECTLOG.md |

---

## 🚀 Oikea tie edessä

Kehitysvaiheiden järjestys (Roadmap):

```
✅ Vaihe 0: Infrastruktuuri (Taskfile, CI, AGENTS, plans)
✅ Vaihe 1: DSP Nuottilaskenta (notes.go, tests)
❌ Vaihe 2: DSP YIN-algoritmi (yin.go, tests)
   → Dokumentaatio: 03-yin-pitch-detector.md (kirjoitettava)
   → Template: PR_STORY_EXTENDED.template.md
   
❌ Vaihe 3: Audio Capture (malgo integration)
   → Dokumentaatio: 04-audio-capture.md (kirjoitettava)
   → Template: PR_STORY_EXTENDED.template.md
   
❌ Vaihe 4: TUI (Bubbletea + Lipgloss)
   → Dokumentaatio: 05-tui-bubbletea.md (kirjoitettava)
   → Template: PR_STORY_EXTENDED.template.md
   
❌ Vaihe 5: Integrointi & Shipping
   → cmd/tuner/main.go: Full program
   → Template: PR_STORY_EXTENDED.template.md
```

---

## 📞 Apua tarvitaan?

```bash
task workflow        # Näytä quick reference
task log            # Avaa projektiloki
cat QUICK_START.md  # Lue käytännöllinen opas
```

---

