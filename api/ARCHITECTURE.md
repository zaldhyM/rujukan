# Dokumentasi Arsitektur — API Rujukan

## Stack Teknologi

| Komponen | Teknologi |
|---|---|
| Language | Go 1.24 |
| HTTP Framework | Gin |
| ORM | GORM |
| Database | MySQL (multi-schema) |
| Cache / Session | Redis |
| Auth | JWT HS256 + bcrypt |
| Validasi | go-playground/validator v10 |

---

## Struktur Folder

```
api/
├── cmd/
│   └── app/
│       └── main.go                         # Entry point aplikasi
├── internal/
│   ├── app/
│   │   └── app.go                          # Inisialisasi router, DB, modul, route
│   ├── infrastructure/
│   │   ├── auth/
│   │   │   └── jwt.go                      # Generate & validasi JWT
│   │   ├── cache/
│   │   │   └── redis.go                    # Koneksi Redis, SessionKey()
│   │   └── database/
│   │       └── mysql.go                    # Koneksi GORM, fungsi Switch schema
│   ├── middleware/
│   │   ├── auth.go                         # Auth middleware (Bearer / Cookie)
│   │   └── cors.go                         # CORS middleware (origin dari env CORS_ORIGIN)
│   ├── modules/
│   │   ├── user/                           # Modul user & auth
│   │   ├── faskes/                         # Modul data master faskes
│   │   ├── pasien/                         # Modul data master pasien
│   │   ├── wilayah/                        # Modul data master wilayah
│   │   ├── referensi/                      # Modul data master referensi
│   │   └── rujukan/                        # Modul alur rujukan pasien
│   └── pkg/
│       └── response/
│           └── response.go                 # Helper response standar
├── db/
│   ├── aplikasi/tables/                    # SQL schema user
│   ├── master/tables/                      # SQL schema data master
│   ├── rujukan/                            # SQL schema rujukan
│   └── patches/                            # Patch SQL incremental
├── go.mod
└── go.sum
```

---

## Pola Struktur Modul

Setiap modul mengikuti pola yang sama:

```
modul/
├── domain/
│   └── entity.go        # Struct GORM (representasi tabel database)
├── repository/
│   └── mysql.go         # Interface + implementasi query ke database
├── delivery/
│   └── http/
│       └── handler.go   # Handler Gin (menerima request, memanggil repository)
└── module.go            # Wiring dependency + registrasi route
```

**Alur dependency:**

```
module.go
  └── NewXModule(db) → repository → handler
```

`module.go` adalah satu-satunya titik yang mengenal semua layer. Handler tidak tahu database secara langsung — hanya tahu interface repository.

---

## Alur Request

```
HTTP Request
    │
    ▼
CORSMiddleware (global — semua route)
    │  ├── Preflight OPTIONS → 204 langsung
    │  └── Set header Access-Control-* sesuai origin
    │
    ▼
Gin Router (/v1/...)
    │
    ├── [Public]  → Handler langsung
    │   └── POST /v1/auth/login
    │   └── POST /v1/auth/register
    │
    └── [Protected] → AuthMiddleware
                          │
                          ├── Cek Bearer token / cookie
                          ├── Validasi JWT (signature + expiry)
                          ├── Verifikasi sesi aktif di Redis
                          ├── Cek status & masa aktif user di DB
                          └── Set context: userID, username, role, user
                                │
                                ▼
                            Handler
                              │
                              ├── Bind & validasi input (ShouldBindJSON)
                              ├── Panggil repository
                              └── Kembalikan response via pkg/response
```

---

## Layer Penjelasan

### `cmd/app/main.go`

Entry point. Tugasnya tiga hal:
1. Load `.env` via `godotenv`
2. Defer `database.Close()` dan `cache.Close()`
3. Jalankan `app.NewApp().Start()`

### `internal/app/app.go`

Inisialisasi seluruh aplikasi dengan urutan:

1. Hubungkan DB (`database.Connect()`) dan Redis (`cache.Connect()`)
2. Set mode Gin dari env `GIN_MODE` (`debug` / `release`)
3. Buat router, set trusted proxy ke `127.0.0.1`
4. Pasang `CORSMiddleware` secara global (sebelum route manapun)
5. Inisialisasi semua modul dengan dependency injection manual
6. Daftarkan route ke group `/v1`:
   - **Public** — `RegisterAuthRoutes` (login, register)
   - **Protected** — semua route lain, di-wrap `AuthMiddleware`

### `internal/infrastructure/database/mysql.go`

- Koneksi GORM ke MySQL, singleton via `sync.Once`
- Konfigurasi dari env: `DB_HOST`, `DB_PORT`, `DB_USER`, `DB_PASS`, `DB_NAME`
- Fungsi `Switch(dbName string)` — eksekusi `USE <schema>` untuk ganti schema aktif dalam satu koneksi

### `internal/infrastructure/cache/redis.go`

- Koneksi Redis dari env `REDIS_ADDR` (default: `127.0.0.1:6379`)
- `SessionKey(userID uint) string` → menghasilkan key `session:<id>` untuk menyimpan token aktif

### `internal/infrastructure/auth/jwt.go`

- `GenerateToken(...)` — buat JWT HS256 dengan custom claims: `user_id`, `username`, `role`, `id_faskes`, `nama`
- `ValidateToken(tokenString)` — parse dan validasi JWT, kembalikan `*Claims`
- Secret key dari env `JWT_SECRET`

### `internal/middleware/auth.go`

Middleware autentikasi yang dijalankan sebelum setiap protected handler:

1. Ambil token dari header `Authorization: Bearer <token>` atau cookie `session_token`
2. Validasi JWT (signature + expiry)
3. Cocokkan token dengan Redis key `session:<userID>` — jika tidak cocok, sesi dianggap tidak valid
4. Ambil data user dari DB, cek `STATUS = 1` dan `MASA_AKTIF`
5. Set ke context Gin: `userID`, `username`, `role`, `user`

### `internal/middleware/cors.go`

Middleware CORS yang dipasang secara global di router (sebelum semua route):

- **Origin** yang diizinkan dikonfigurasi via env `CORS_ORIGIN` (koma-separated)
- Default development: `http://localhost:3000,http://localhost:5173`
- Request preflight `OPTIONS` dijawab langsung dengan status `204` tanpa masuk ke handler
- Header yang di-set: `Access-Control-Allow-Origin`, `Allow-Methods`, `Allow-Headers`, `Allow-Credentials`, `Max-Age`
- `Vary: Origin` ditambahkan agar cache proxy tidak salah menyajikan response

### `internal/pkg/response/response.go`

Helper untuk response JSON yang konsisten di semua handler:

| Fungsi | Status | Kegunaan |
|---|---|---|
| `OK(c, data)` | 200 | Response sukses dengan data |
| `OKList(c, data, total, limit, offset)` | 200 | Response list dengan pagination |
| `OKMessage(c, message)` | 200 | Response sukses dengan pesan teks |
| `Created(c, data)` | 201 | Response setelah create data baru |
| `BadRequest(c, message)` | 400 | Input tidak valid |
| `Unauthorized(c, message)` | 401 | Tidak terautentikasi |
| `Conflict(c, message)` | 409 | Konflik data (misal: username duplikat) |
| `NotFound(c, message)` | 404 | Data tidak ditemukan |
| `InternalError(c, message)` | 500 | Error server |
| `ValidationError(c, err)` | 400 | Error validasi dengan detail per field |

**Format error validasi:**

```json
{
  "success": false,
  "message": "Validasi input gagal",
  "errors": {
    "password": "minimal 4 karakter",
    "role": "harus salah satu dari: admin, faskes, dinkeskab, dinkesprov"
  }
}
```

---

## Modul User (`internal/modules/user/`)

Modul ini menangani dua kelompok route:

**Auth (public):**
- `POST /v1/auth/login` — verifikasi username + password (bcrypt), generate JWT, simpan ke Redis, set cookie
- `POST /v1/auth/register` — buat user baru, hash password bcrypt

**Protected:**
- `GET /v1/auth/me` — profil user dari context middleware
- `POST /v1/auth/logout` — hapus sesi dari Redis, clear cookie
- `GET /v1/aplikasi/user` — list semua user (kolom PASSWORD tidak ikut diambil dari DB)

**Fitur khusus login:**
- Mendukung fallback password plaintext (untuk migrasi data lama) — jika cocok, password otomatis di-upgrade ke bcrypt
- Single active session: login baru menimpa sesi lama di Redis
- Durasi token dari env `TOKEN_EXPIRED_HOURS` (default: 24 jam)

### Entity User

Tabel: `aplikasi.user`

| Field | Tipe | Keterangan |
|---|---|---|
| `ID` | uint | Primary key |
| `USERNAME` | string | Unik |
| `PASSWORD` | string | bcrypt hash — **tidak pernah dikembalikan ke client** |
| `NAMA` | string | Nama lengkap |
| `NIK` | string | Nomor Induk Kependudukan |
| `ROLE` | string | `admin` / `faskes` / `dinkeskab` / `dinkesprov` |
| `ID_FASKES` | uint | FK ke faskes |
| `STATUS` | int8 | `1` = aktif, `0` = nonaktif |
| `MASA_AKTIF` | string | Tanggal kedaluarsa akun (`YYYY-MM-DD`) |
| `TOKEN` | string | — |

---

## Modul Rujukan (`internal/modules/rujukan/`)

Mengelola alur rujukan pasien antar fasilitas kesehatan.

### Endpoint

| Method | URL | Keterangan |
|---|---|---|
| `GET` | `/v1/rujukan` | List rujukan (filter: `search`, `status`, `limit`, `offset`) |
| `GET` | `/v1/rujukan/:id` | Detail rujukan |
| `POST` | `/v1/rujukan` | Buat rujukan baru |
| `PUT` | `/v1/rujukan/:id/status` | Update status rujukan |

### Alur Status

```
menunggu → diterima → selesai
         ↘ ditolak
```

### Entity Rujukan

Tabel: `rujukan.rujukan`

| Field | Tipe | Keterangan |
|---|---|---|
| `ID` | int | Primary key, auto increment |
| `NO_RUJUKAN` | string | Auto-generate: `RJK-YYYYMMDD-XXXXXXXX` |
| `ID_PASIEN` | string | ID pasien |
| `ID_FASKES_ASAL` | string | Kode faskes asal (char 10) |
| `ID_FASKES_TUJUAN` | string | Kode faskes tujuan (char 10) |
| `KODE_ICD` | *string | Kode diagnosa ICD (opsional) |
| `DIAGNOSA` | *string | Deskripsi diagnosa (opsional) |
| `CATATAN` | *string | Catatan tambahan (opsional) |
| `STATUS` | string | `menunggu` / `diterima` / `ditolak` / `selesai` |
| `ID_USER` | uint | Otomatis dari JWT context |
| `TANGGAL` | *time.Time | Auto set saat create |

---

## Modul Data Master

Modul-modul berikut hanya menyediakan data referensi (read + CRUD sederhana):

| Modul | Tabel | Route prefix |
|---|---|---|
| `faskes` | `master.faskes` | `/v1/faskes` |
| `pasien` | `master.pasien` | `/v1/pasien` |
| `wilayah` | `master.wilayah`, `master.jenis_wilayah` | `/v1/wilayah` |
| `referensi` | `master.referensi`, `master.jenis_referensi` | `/v1/referensi` |

Semua modul master mendukung query parameter: `search`, `limit`, `offset`.

---

## Database — Skema MySQL

Aplikasi menggunakan tiga schema MySQL yang diakses melalui satu koneksi dengan `USE <schema>`:

| Schema | Tabel |
|---|---|
| `aplikasi` | `user` |
| `master` | `faskes`, `pasien`, `wilayah`, `jenis_wilayah`, `referensi`, `jenis_referensi` |
| `rujukan` | `rujukan` |

Perpindahan schema dilakukan via `database.Switch("nama_schema")` — fungsi ini menjalankan `USE <schema>` pada koneksi GORM yang aktif.

---

## Keamanan

### Yang Sudah Diimplementasikan

| Aspek | Implementasi |
|---|---|
| SQL Injection | Semua query GORM menggunakan parameterized (`WHERE "field = ?"`) |
| Password | bcrypt `DefaultCost`, migrasi otomatis dari plaintext |
| JWT | HS256, signature + expiry divalidasi setiap request |
| Session | Single active session via Redis — token lama invalid saat login baru |
| Password exposure | `QueryAll` user menggunakan `SELECT` eksplisit, kolom `PASSWORD` tidak diambil dari DB |
| CORS | `CORSMiddleware` global, origin dikontrol via env `CORS_ORIGIN` |
| HttpOnly cookie | Cookie `session_token` tidak bisa diakses JavaScript |

### Backlog (Belum Diimplementasikan)

Lihat `NOTES.md` bagian **Backlog Keamanan** untuk detail dan solusi yang direncanakan:

- Rate limiting pada `POST /v1/auth/login` (brute force)
- Role-based authorization (RBAC) per endpoint
- Tidak meneruskan error internal ke client
- Cookie `Secure: true` untuk environment production
- Validasi wajib `JWT_SECRET` saat startup (tolak fallback)

---

## Environment Variables

File `.env` di root `api/`:

```env
# Database
DB_HOST=127.0.0.1
DB_PORT=3306
DB_USER=root
DB_PASS=password
DB_NAME=aplikasi

# Redis
REDIS_ADDR=127.0.0.1:6379

# JWT
JWT_SECRET=isi_dengan_secret_acak_yang_kuat

# CORS — origin frontend yang diizinkan (koma-separated)
CORS_ORIGIN=http://localhost:3000,http://localhost:5173

# Server
PORT=8081
GIN_MODE=debug          # debug | release
TOKEN_EXPIRED_HOURS=24
```

---

## Menjalankan Aplikasi

```bash
# Dari direktori api/
go run cmd/app/main.go

# Atau build terlebih dahulu
go build -o server cmd/app/main.go
./server
```

Server akan berjalan di `http://localhost:8081` (atau sesuai `PORT`).

---

## Konvensi Response API

**Sukses (single data):**
```json
{ "success": true, "data": { ... } }
```

**Sukses (list dengan pagination):**
```json
{ "success": true, "total": 100, "limit": 10, "offset": 0, "data": [ ... ] }
```

**Sukses (pesan):**
```json
{ "success": true, "message": "..." }
```

**Error:**
```json
{ "success": false, "message": "..." }
```

**Error validasi:**
```json
{ "success": false, "message": "Validasi input gagal", "errors": { "field": "pesan" } }
```
