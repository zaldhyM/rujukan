-- Buat user 'admin' (jika belum otomatis dibuat)
CRCREATE USER 'admin'@'%' IDENTIFIED WITH caching_sha2_password BY 'S!MGos2@kemkes.go.id';

-- Berikan akses penuh ke semua database (termasuk mysql.*, performance_schema.*, sys.*, test, dll)
GRANT ALL PRIVILEGES ON *.* TO 'admin'@'%' WITH GRANT OPTION;

-- Tambahan: kalau test masih ngeyel, grant spesifik juga
GRANT ALL PRIVILEGES ON test.* TO 'admin'@'%';

-- Refresh privilege
FLUSH PRIVILEGES;
