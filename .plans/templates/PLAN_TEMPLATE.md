# Suunnitelmapohja: [Ominaisuuden nimi suomeksi]

**Tunniste:** `.plans/[NN]-[kebab-case-nimi].md`  
**Päivämäärä:** YYYY-MM-DD  
**Tila:** LUONNOS / HYVÄKSYTTY / TOTEUTUKSESSA / VALMIS  
**Tekijä:** Tekoäly-arkkitehtimentori & Kehittäjä  

---

## 1. Tavoite ja käyttäjäkokemus (UX & Liiketoimintakonteksti)

### Nykytila ja ongelma

* Kuvaa lyhyesti nykyinen toiminta tai puuttuva ominaisuus.
* Mitä ongelmia tai rajoitteita käyttäjä kohtaa? (Esim. 401 Unauthorized vierastilassa, puuttuva normalisointi haussa).

### Tavoitetila ja arvolupaus

* Miten ratkaisu toimii loppukäyttäjälle tai kehittäjälle?
* Konkreettiset UX-parannukset (esim. ilmoitusbannerit, viiveettömät siirtymät, automaattinen tallennus).

---

## 2. Arkkitehtuuri ja komponenttirajat

### Kerrosvastuut (Layer Boundaries)

* **API Layer (`internal/api/`)**: HTTP-reititys (`Go 1.22+ ServeMux`), request decoding, vastausten serialisointi, $O(1)$ virtaus.
* **Service Layer (`internal/services/`)**: Liiketoimintalogiikka, orkestrointi, eräkäsittely (500 tietueen erissä).
* **Repository Layer (`internal/db/`)**: Tietokantakyselyt (`*sql.DB`), parametrisoidut kyselyt (`$1, $2`), `context.Context`-peruutus.
* **Frontend Layer (`frontend/src/`)**: React 19.2 -komponentit, TailwindCSS v4, tyypitykset, kaksikielisyys (`i18n.ts`).

---

## 3. Tietokantamigraatiot (tarvittaessa)

> [!NOTE]
> Kaikkien tietokantamuutosten tulee olla 100 % yhteensopivia Neon PostgreSQL -tuotantokannan kanssa ja toimia rinnakkain in-memory SQLite (`:memory:`) -yksikkötesteissä.

### Tiedosto: `backend/migrations/[NNN]_[nimi].sql`

```sql
-- Up migration
CREATE TABLE IF NOT EXISTS feature_items (
    id VARCHAR(64) PRIMARY KEY,
    user_id VARCHAR(64) REFERENCES users(id) ON DELETE CASCADE,
    title TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX IF NOT EXISTS idx_feature_items_user_id ON feature_items(user_id);
```

---

## 4. Backend-toteutus (Go 1.22+)

### 4.1 Reititys ja API (`internal/api/`)

* **Endpoint:** `METHOD /api/...`
* **Pääsynhallinta:** `middleware.RequireAuth` tai `middleware.OptionalAuth`
* **Kontekstin purku:** `ctxkeys.GetUserID(r.Context())`

### 4.2 Tietokantakerros (`internal/db/`)

* Kyselyiden parametrisointi ja peruutettavuus (`QueryContext`, `ExecContext`).

### 4.3 Palvelukerros (`internal/services/`)

* Virheenkäsittely ja resurssienhallinta (ei hiljaista virheiden nielentää).

---

## 5. Frontend-toteutus (React 19.2 & TypeScript)

### 5.1 React 19.2 & React Compiler Pre-Flight Audit

> [!IMPORTANT]
> Jokaisessa frontend-suunnitelmassa tulee nimenomaisesti arvioida ja varmistaa seuraavat React 19.2 & Compiler -konventiot:
>
> 1. **Zero `useEffect` for State Sync:** Älä synkronoi tilaa tai kuuntele ulkoisia selaimen/ikkunan tapahtumia `useEffect` + `useState` -parilla. Käytä **`useSyncExternalStore`**-hookia.
> 2. **`useActionState` & Form Actions:** Korvaa asynkronisten pyyntöjen vanhanaikaiset `useState`-liput (`loading`, `saving`, `error`) React 19.2:n `useActionState`- ja `<form action={...}>` -malleilla.
> 3. **Puhdas johdettu tila (Derived State):** Vältä redundanttia tilaa. Laske arvot render-aikana.
> 4. **Ei prop-muutos-efektejä:** Älä aseta tai resetoi tilaa `useEffect`:illä propin muuttuessa. Käytä `key`-attribuuttia tai render-aikaista säätöä.
> 5. **Kaksikielisyys (`i18n.ts`):** Kaikki tekstit suomeksi ja englanniksi.

### 5.2 Tyypitykset (`frontend/src/types/...`)

```typescript
export interface FeatureModel {
  id: string;
  userId?: string;
  title: string;
  createdAt: string;
}
```

### 5.3 Komponentit ja näkymät

* Standardi funktionaalinen komponenttisyntaksi (`export function ComponentName(props: Props): JSX.Element`).
* TailwindCSS v4 -luokat ja responsiivisuus.

### 5.4 Kaksikielisyys (`frontend/src/i18n.ts`)

* Kaikki tekstit, napit ja virheilmoitukset lisätään sekä suomeksi (`fi`) että englanniksi (`en`).

---

## 6. Varmistus- ja testaussuunnitelma

### 6.1 Automaatiotestit

* **Backend:** `task backend:check` (mod tidy, linter, yksikkötestit ja kattavuus).
* **Frontend:** `task frontend:check` (tyyppitarkistus, ESLint, Vitest).
* **Koko projekti:** `task check`.

### 6.2 Manuaalinen tarkistuslista selaimessa

* [ ] Kirjautuneen käyttäjän näkymä ja toiminnot toimivat.

* [ ] Vieras- / anonyymitila toimii ilman 401-virheitä tai rikkinäistä tilaa.
* [ ] Kielenvaihto fi <-> en toimii virheettömästi.
* [ ] Tumma ja vaalea teema (Dark/Light mode) renderöityy oikein.

---

## 7. Vaiheittainen toteutusopas kehittäjälle (Step-by-Step)

> [!TIP]
> Tämä osio opastaa kehittäjää koodaamaan ominaisuuden vaihe kerrallaan. Agentti antaa tarvittavat koodirakenteet ja selittää valitut ratkaisut.

### Vaihe 1: Tietokanta ja mallit

1. Lisää migraatiotiedosto `backend/migrations/...`
2. Määritä tietomallit tiedostoon `backend/internal/models/...`

### Vaihe 2: Backend-palvelut ja reititys

1. Kirjoita repository-metodit tiedostoon `backend/internal/db/...`
2. Kirjoita palvelulogiikka tiedostoon `backend/internal/services/...`
3. Rekisteröi API-reitti tiedostoon `backend/internal/api/...`

### Vaihe 3: Frontend-integraatio

1. Lisää API-kutsu tiedostoon `frontend/src/api/...`
2. Luo UI-komponentti tiedostoon `frontend/src/components/...`
3. Päivitä `frontend/src/i18n.ts`

### Vaihe 4: Laadunvarmistus ja julkaisuvalmistelu

1. Aja `task check` ja korjaa mahdolliset virheet.
2. Suorita oikeasuhtainen semanttinen versiobump (`task version:bump PART=patch|minor`).
3. Laadi PR-tarina `pr_stories/`-kansioon.
