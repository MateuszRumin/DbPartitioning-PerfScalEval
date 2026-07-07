# Mini Forum — Streamlit + MySQL

Mała aplikacja forum korzystająca bezpośrednio z istniejącego modelu Stack Exchange:

- `users` — autorzy,
- `posts` z `post_type_id = 1` — pytania,
- `posts` z `post_type_id = 2` i `parent_id` — odpowiedzi,
- `comments` — komentarze.

Aplikacja działa z wariantem niepartycjonowanym `testdb` i partycjonowanym `testdbp`.

## Funkcje

- lista najnowszych pytań,
- podgląd pytania, odpowiedzi i komentarzy,
- publikowanie nowych pytań,
- publikowanie odpowiedzi,
- publikowanie komentarzy,
- wybór aktywnego użytkownika na podstawie `users.id`.

Nie ma logowania ani autoryzacji. Jest to lokalny interfejs demonstracyjny, a nie kompletna aplikacja produkcyjna.

## 1. Utworzenie środowiska

```bash
cd streamlit_forum

python3 -m venv .venv
source .venv/bin/activate

python -m pip install --upgrade pip
python -m pip install -r requirements.txt
```

## 2. Konfiguracja połączenia

```bash
cp .streamlit/secrets.toml.example .streamlit/secrets.toml
nano .streamlit/secrets.toml
```

Przykład:

```toml
[mysql]
host = "127.0.0.1"
port = 3306
user = "forum"
password = "twoje_haslo"
database = "testdbp"
pool_size = 5
```

Możesz utworzyć osobne konto MySQL:

```sql
CREATE USER 'forum'@'localhost' IDENTIFIED BY 'haslo';
GRANT SELECT, INSERT, UPDATE ON testdbp.* TO 'forum'@'localhost';
FLUSH PRIVILEGES;
```

Dla bazy niepartycjonowanej zmień `testdbp.*` na `testdb.*` i ustaw `database = "testdb"`.

## 3. Indeksy

Przy dużej bazie wykonaj indeksy z pliku `indexes.sql`. Bez indeksu po `parent_id` pobieranie odpowiedzi może wymagać pełnego skanu tabeli `posts`, a bez indeksu po `post_id` analogiczny problem wystąpi dla `comments`.

```bash
mysql -u root -p testdbp < indexes.sql
```

Nie uruchamiaj skryptu ponownie, jeżeli indeksy już istnieją.

## 4. Uruchomienie

```bash
streamlit run app.py
```

Domyślnie aplikacja będzie dostępna pod adresem wyświetlonym przez Streamlit, zwykle `http://localhost:8501`.

## Uwagi do modelu

W partycjonowanym wariancie tabela `posts` ma klucz główny `(id, creation_date)`. Aplikacja korzysta z `AUTO_INCREMENT` dla `id` i zawsze zapisuje `creation_date = NOW()`, dlatego nowe rekordy trafiają do właściwej partycji zakresowej.

Odpowiedź jest tworzona jako rekord `posts` z:

```text
post_type_id = 2
parent_id = id pytania
```

Pytanie ma:

```text
post_type_id = 1
parent_id = NULL
```
