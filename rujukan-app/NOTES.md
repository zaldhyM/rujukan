# Catatan Pengembangan — rujukan-app (Frontend)

## Stack
- **Framework:** Svelte 5 (file-based routing)
- **Build tool:** Vite
- **Language:** TypeScript
- **Dev server:** `npm run dev` → default port `5174`

---

## Struktur Halaman (`src/routes/`)

| Route | File | Keterangan |
|---|---|---|
| `/` | `+page.svelte` | Landing page |
| `/login` | `login/+page.svelte` | Halaman login |
| `/register` | `register/+page.svelte` | Halaman daftar faskes |
| `/dashboard` | `dashboard/+page.svelte` | Ringkasan/KPI |
| `/dashboard/rujukan-masuk` | `dashboard/rujukan-masuk/+page.svelte` | Kelola rujukan masuk |
| `/dashboard/rujukan-keluar` | `dashboard/rujukan-keluar/+page.svelte` | Buat rujukan keluar |
| `/dashboard/pengaturan` | `dashboard/pengaturan/+page.svelte` | Pengaturan akun |

Layout dashboard (sidebar + topbar + footer): `dashboard/+layout.svelte`  
CSS global + variabel tema: `src/routes/global.css`

---

## Tema — Opsi 3 Minimalis (aktif)

Redesign selesai 2026-06-26. Seluruh tampilan menggunakan palet minimalis.

```css
--bg-primary:          #F9FAFB   /* latar halaman */
--bg-secondary:        #FFFFFF   /* kartu/panel */
--color-primary:       #14B8A6   /* teal utama */
--color-primary-hover: #0F9488
--color-secondary:     #A3E635   /* lime aksen */
--text-primary:        #374151
--text-muted:          #6B7280
--border-glass:        #E5E7EB
```

**Aturan desain:**
- Sidebar dashboard: putih `#FFFFFF`, border kanan `#E5E7EB`
- Nav item aktif: background `#14B8A6`, teks putih
- Nav item hover: teks teal, background teal transparan
- Topbar: putih, tinggi 64px, ada search bar pill-shaped
- Tombol flat (tanpa gradien): primary = teal solid, outline = border `#D1D5DB`
- Badge status: Selesai `#DCFCE7/#16A34A`, Menunggu `#FEF3C7/#D97706`, Ditolak `#FEE2E2/#DC2626`
- Modal header: `#14B8A6`, backdrop: `rgba(0,0,0,0.4)`
- Toast: bg putih `#FFFFFF`, border `#BBF7D0`

**Halaman Login/Register:**
- Background `#F0F4F8`, kartu terpusat
- Wave SVG dekorasi di sudut kiri-atas dan kanan-bawah
- Logo Kemenkes + eye-toggle password + checkbox "Ingat saya"
- SSO button (Google + SSO Kemenkes), security badge di bawah kartu

**Landing Page (`/`):**
- Navbar sticky: logo Kemenkes + nav links + tombol "Masuk"
- Hero 2-kolom: kiri (heading "Sehat Untuk **Indonesia**" + CTA) + kanan (ilustrasi SVG 3 tenaga medis)
- Blob dekorasi: hijau organik + lingkaran teal di belakang ilustrasi
- 3 floating badge cards animasi
- Features strip 4-kolom di bawah hero

---

## Status Halaman

| Halaman | Status | Catatan |
|---|---|---|
| Landing page | ✅ Selesai | Hero 2-kolom, navbar, ilustrasi SVG, floating cards |
| Login | ✅ Selesai | Kartu terpusat, wave SVG, SSO button, eye-toggle |
| Register | ✅ Selesai | Gaya sama login, form 2-kolom, 2 eye-toggle |
| Dashboard | ✅ Selesai | Mock data, sidebar putih, topbar + search bar |
| Rujukan Masuk | ⏳ Partial | Tampilan selesai, belum connect ke API |
| Rujukan Keluar | ⏳ Partial | Tampilan selesai, belum connect ke API |
| Pengaturan | ⏳ Partial | Tampilan selesai, belum connect ke API |

---

## Yang Perlu Dilakukan Selanjutnya

- [ ] Integrasi API login ke `POST /v1/auth/login`
- [ ] Simpan JWT token ke localStorage/cookie setelah login
- [ ] Koneksi halaman Rujukan Masuk ke `GET /v1/rujukan?status=menunggu`
- [ ] Koneksi halaman Rujukan Keluar ke `POST /v1/rujukan`
- [ ] Update status rujukan via `PUT /v1/rujukan/:id/status`
- [ ] Halaman pengaturan: ganti password, update profil
