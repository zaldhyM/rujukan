# Catatan Pengembangan — api (Backend)

## Stack
- **Language:** Go
- **Framework:** Gin
- **ORM:** GORM
- **Database:** MySQL (multi-schema)
- **Auth:** JWT HS256 + bcrypt

---

## Struktur Folder

```
api/
├── cmd/app/main.go                  # Entry point
├── internal/
│   ├── app/app.go                   # Inisialisasi router, DB, modul
│   ├── infrastructure/
│   │   ├── auth/jwt.go              # Generate & validasi JWT
│   │   └── database/mysql.go        # Koneksi GORM, fungsi Switch schema
│   ├── middleware/auth.go           # Auth middleware (Bearer/Cookie)
│   └── modules/
│       ├── user/                    # Login, Register, Logout, Me, CRUD user
│       ├── faskes/                  # Data master faskes
│       ├── pasien/                  # Data master pasien
│       ├── wilayah/                 # Data master wilayah
│       ├── referensi/               # Data master referensi
│       └── rujukan/                 # ← BARU: modul rujukan
```

Setiap modul mengikuti pola:
```
modul/
├── domain/entity.go          # Struct GORM
├── repository/mysql.go       # Interface + implementasi query
├── delivery/http/handler.go  # Handler Gin
└── module.go                 # Inisialisasi & registrasi route
```

---

## Database Schemas

| Schema | Isi |
|---|---|
| `aplikasi` | `user`, `rujukan` |
| `master` | `faskes`, `pasien`, `wilayah`, `jenis_wilayah`, `referensi`, `jenis_referensi` |

SQL skema ada di `db/`:
```
db/
├── aplikasi/tables/user.sql
├── master/tables/faskes.sql, pasien.sql, wilayah.sql, ...
└── rujukan/
    ├── create_db.sql
    └── tables/rujukan.sql
```

---

## Endpoints

### Auth (Public)
| Method | URL | Keterangan |
|---|---|---|
| POST | `/v1/auth/login` | Login, dapat JWT token |
| POST | `/v1/auth/register` | Daftar user baru |

### Protected (perlu Bearer token / cookie)
| Method | URL | Keterangan |
|---|---|---|
| GET | `/v1/aplikasi/user` | List semua user |
| GET | `/v1/auth/me` | Profil user login |
| POST | `/v1/auth/logout` | Logout |
| GET | `/v1/faskes` | List faskes |
| GET | `/v1/pasien` | List pasien |
| GET | `/v1/wilayah` | List wilayah |
| GET | `/v1/referensi` | List referensi |
| GET | `/v1/rujukan` | List rujukan (filter: search, status, limit, offset) |
| GET | `/v1/rujukan/:id` | Detail rujukan |
| POST | `/v1/rujukan` | Buat rujukan baru |
| PUT | `/v1/rujukan/:id/status` | Update status rujukan |

### Status nilai `rujukan`:
`menunggu` → `diterima` / `ditolak` → `selesai`

---

## Modul Rujukan — Detail

**Entity penting (`domain/entity.go`):**
- `ID_FASKES_ASAL` dan `ID_FASKES_TUJUAN` → `char(10)` (kode faskes, bukan ID integer)
- `ID_USER` → otomatis diisi dari JWT context saat `POST /v1/rujukan`
- `NO_RUJUKAN` → auto-generate format `RJK-YYYYMMDD-XXXXXXXX`
- `TableName()` → `rujukan.rujukan`

---

## Status Pengembangan

| Fase | Keterangan | Status |
|---|---|---|
| Fase 1 | Middleware & Security (JWT, CORS, Logger) | ✅ JWT & Auth Middleware selesai |
| Fase 2 | Auth lengkap (Login, Register, Logout, Me) | ✅ Selesai |
| Fase 3 | Modul data master + Modul Rujukan | ✅ Selesai |
| Fase 4 | Validasi input & standarisasi error response | ⏳ Belum |
| Fase 5 | Swagger/OpenAPI + migrasi DB otomatis | ⏳ Belum |

---

## Rencana Pengembangan Detail (Roadmap)

### 🚀 Fase 1: Middleware & Security ✅
- **JWT Authentication Middleware** — token-based auth untuk endpoint sensitif.
- **CORS Middleware** — agar `rujukan-app` dapat mengakses API dengan aman.
- **Logger & Recovery Middleware** — mencatat log request dan mencegah server crash saat panic.

### 📦 Fase 2: Peningkatan Fitur Modul User ✅
- `POST /v1/auth/login` — verifikasi bcrypt password + generate JWT.
- `POST /v1/auth/register` — pendaftaran user baru dengan enkripsi password.
- `GET /v1/auth/me` — profil user dari context JWT.
- `POST /v1/auth/logout` — hapus token dari DB + clear cookie.
- CRUD lengkap: `GET /v1/aplikasi/user/:id`, `POST`, `PUT`, `DELETE`.

### 🏥 Fase 3: Modul Data Master + Rujukan ✅
- **Faskes** — data rumah sakit, puskesmas, klinik.
- **Pasien** — manajemen data medis dasar pasien.
- **Wilayah** — provinsi, kota/kabupaten, kecamatan, kelurahan.
- **Referensi** — kelompok faskes dan jenis kelamin.
- **Rujukan** ✅ — alur rujukan pasien antar faskes:
  - `POST /v1/rujukan` — buat rujukan baru.
  - `GET /v1/rujukan/:id` — cek status rujukan.
  - `GET /v1/rujukan` — list dengan filter.
  - `PUT /v1/rujukan/:id/status` — update status.

### 📝 Fase 4: Validasi & Standarisasi Error ⏳
- **Request Validation:** Integrasi `github.com/go-playground/validator` pada handler.
- **Standard Response Helper:** Format response konsisten untuk sukses dan error di semua modul.

### 📖 Fase 5: Dokumentasi API & Deployment ⏳
- **Swagger/OpenAPI:** Integrasi `swaggo/swag` untuk generate dokumentasi otomatis.
- **Database Migration:** Patch SQL terurut di `db/patches/`, integrasi `golang-migrate`.

---

## Pengelolaan Database Patch

```
db/
├── init.sql
├── aplikasi/           # Schema aplikasi (user, rujukan)
├── master/             # Schema master (faskes, pasien, dll)
├── rujukan/            # Schema rujukan (tabel rujukan)
│   ├── create_db.sql
│   └── tables/rujukan.sql
└── patches/            # Perubahan schema incremental
    └── README.md
```

Setiap perubahan skema wajib dicatat sebagai file patch baru di `db/patches/` dengan format `XXXX_deskripsi.sql`.

---

## Yang Perlu Dilakukan Selanjutnya

- [ ] Jalankan `db/rujukan/create_db.sql` lalu `db/rujukan/tables/rujukan.sql` di database
- [ ] Fase 4: Integrasi `go-playground/validator` di semua handler
- [ ] Fase 4: Buat helper `response.go` untuk format JSON standar (success/error)
- [ ] Fase 4: CORS middleware untuk akses dari `rujukan-app`
- [ ] Fase 5: Integrasi `swaggo/swag` untuk dokumentasi otomatis
- [ ] Fase 5: Setup `golang-migrate` untuk eksekusi patch otomatis
