# Catatan Perubahan

## 2026-06-27 — Pindah Tombol Collapse Sidebar ke Topbar

**File:** `src/routes/dashboard/+layout.svelte`

### Apa yang diubah
- Tombol "Kecilkan Menu" **dihapus** dari `sidebar-footer`
- Tombol collapse **dipindahkan** ke topbar, di sebelah kiri search bar
- Tombol hanya menampilkan **ikon** (tanpa teks): panel sidebar (kotak dengan garis vertikal di kiri)

### Detail Teknis
- Class baru: `.collapse-toggle-btn`
- Fungsi yang dipanggil: `toggleCollapse()`
- Disembunyikan di mobile (≤1024px) — hamburger menu tetap dipakai di mobile
- Style lama `.collapse-btn` dihapus
