import { browser } from '$app/environment';

// Helper to check if logged in
export function isLoggedIn() {
    if (!browser) return false;
    return localStorage.getItem('isLoggedIn') === 'true';
}

// Get current logged-in user profile
export function getCurrentUser() {
    if (!browser) return { name: 'Dr. Admin', username: '12345' };
    const saved = localStorage.getItem('currentUser');
    if (saved) {
        try {
            return JSON.parse(saved);
        } catch (e) {
            // fallback
        }
    }
    return { name: 'Dr. Zaldhy Masud', username: '12345', role: 'Dokter Umum' };
}

/**
 * Update current user profile
 * @param {any} userData
 */
export function updateCurrentUser(userData) {
    if (!browser) return;
    localStorage.setItem('currentUser', JSON.stringify(userData));
}

// Get Faskes Settings
export function getFaskesSettings() {
    const defaultSettings = {
        name: 'RSUD Kebayoran Baru',
        code: '3174012',
        type: 'Rumah Sakit Tipe C',
        address: 'Jl. Gandaria Tengah III No.17, RT.1/RW.3, Kramat Pela, Kebayoran Baru, Jakarta Selatan',
        phone: '(021) 7200234',
        email: 'info@rsudkebayoranbaru.go.id'
    };
    if (!browser) return defaultSettings;
    const saved = localStorage.getItem('faskesSettings');
    if (saved) {
        try {
            return JSON.parse(saved);
        } catch (e) {
            // fallback
        }
    }
    return defaultSettings;
}

/**
 * Update Faskes Settings
 * @param {any} settings
 */
export function updateFaskesSettings(settings) {
    if (!browser) return;
    localStorage.setItem('faskesSettings', JSON.stringify(settings));
}

// Initial dummy data for incoming referrals
const defaultRujukanMasuk = [
    {
        id: 'RM-2026-001',
        noRujukan: '0123R0010526B0001',
        tglRujukan: '2026-05-18',
        pasien: {
            nama: 'Budi Santoso',
            nik: '3174091204850002',
            noKartu: '0001234567890',
            gender: 'Laki-laki',
            usia: 40
        },
        faskesAsal: 'Puskesmas Kebayoran Baru',
        diagnosa: 'Hipertensi Primer (Essential Hypertension)',
        poliTujuan: 'Poli Penyakit Dalam',
        status: 'Pending', // Pending, Diterima, Ditolak
        catatan: 'Hipertensi tidak terkontrol dengan terapi amlodipin 10mg. Perlu pemeriksaan dan evaluasi spesialis Penyakit Dalam.'
    },
    {
        id: 'RM-2026-002',
        noRujukan: '0123R0010526B0002',
        tglRujukan: '2026-05-19',
        pasien: {
            nama: 'Siti Rahmawati',
            nik: '3174095208910003',
            noKartu: '0002345678901',
            gender: 'Perempuan',
            usia: 28
        },
        faskesAsal: 'Klinik Medika Utama',
        diagnosa: 'Appendicitis Akut',
        poliTujuan: 'Poli Bedah Umum',
        status: 'Diterima',
        catatan: 'Nyeri perut kanan bawah sejak 2 hari, McBurney sign (+), Leukositosis 14.000/uL. Rujuk cito untuk rencana appendectomy.'
    },
    {
        id: 'RM-2026-003',
        noRujukan: '0123R0010526B0003',
        tglRujukan: '2026-05-20',
        pasien: {
            nama: 'Wawan Hermawan',
            nik: '3174092211600004',
            noKartu: '0003456789012',
            gender: 'Laki-laki',
            usia: 65
        },
        faskesAsal: 'Puskesmas Pesanggrahan',
        diagnosa: 'Katarak Senilis Matur',
        poliTujuan: 'Poli Mata',
        status: 'Pending',
        catatan: 'Visus OD 1/300, OS 6/60. Lensa keruh total pada mata kanan. Rujuk untuk operasi katarak.'
    },
    {
        id: 'RM-2026-004',
        noRujukan: '0123R0010526B0004',
        tglRujukan: '2026-05-15',
        pasien: {
            nama: 'Ani Wijaya',
            nik: '3174094406980001',
            noKartu: '0004567890123',
            gender: 'Perempuan',
            usia: 25
        },
        faskesAsal: 'Puskesmas Kebayoran Lama',
        diagnosa: 'Hyperemesis Gravidarum Grade II',
        poliTujuan: 'Poli Kebidanan dan Kandungan',
        status: 'Ditolak',
        catatan: 'Kehamilan G1P0A0 10 minggu dengan mual muntah berlebihan, ketonuria (+). Mohon penanganan rawat inap.'
    }
];

// Initial dummy data for outgoing referrals
const defaultRujukanKeluar = [
    {
        id: 'RK-2026-001',
        noRujukan: '3174R0020526K0001',
        tglRujukan: '2026-05-17',
        pasien: {
            nama: 'Joko Widodo',
            nik: '3174091510700001',
            noKartu: '0005678901234',
            gender: 'Laki-laki',
            usia: 56
        },
        faskesTujuan: 'RSCM (Rumah Sakit Cipto Mangunkusumo)',
        diagnosa: 'Penyakit Jantung Koroner (CAD)',
        poliTujuan: 'Poli Jantung & Pembuluh Darah',
        status: 'Selesai', // Dikirim, Diproses, Selesai
        catatan: 'Angina pectoris khas, hasil EKG menunjukkan ST depresi di V1-V4. Dirujuk untuk tindakan kateterisasi jantung (coronary angiography).'
    },
    {
        id: 'RK-2026-002',
        noRujukan: '3174R0020526K0002',
        tglRujukan: '2026-05-19',
        pasien: {
            nama: 'Dewi Lestari',
            nik: '3174094101850005',
            noKartu: '0006789012345',
            gender: 'Perempuan',
            usia: 41
        },
        faskesTujuan: 'RS Kanker Dharmais',
        diagnosa: 'Karsinoma Mammae Sinistra Suspek Metastase',
        poliTujuan: 'Poli Bedah Onkologi',
        status: 'Diproses',
        catatan: 'Massa payudara kiri ukuran 5x4 cm berbenjol-benjol, mobile terbatas. Pembesaran KGB aksila (+). Rujuk untuk biopsi dan staging.'
    }
];

// Get Rujukan Masuk from localStorage
export function getRujukanMasuk() {
    if (!browser) return defaultRujukanMasuk;
    const saved = localStorage.getItem('rujukanMasuk');
    if (saved) {
        try {
            return JSON.parse(saved);
        } catch (e) {
            // fallback
        }
    }
    // Initialize if empty
    localStorage.setItem('rujukanMasuk', JSON.stringify(defaultRujukanMasuk));
    return defaultRujukanMasuk;
}

/**
 * Save Rujukan Masuk to localStorage
 * @param {any[]} data
 */
export function saveRujukanMasuk(data) {
    if (!browser) return;
    localStorage.setItem('rujukanMasuk', JSON.stringify(data));
}

// Get Rujukan Keluar from localStorage
export function getRujukanKeluar() {
    if (!browser) return defaultRujukanKeluar;
    const saved = localStorage.getItem('rujukanKeluar');
    if (saved) {
        try {
            return JSON.parse(saved);
        } catch (e) {
            // fallback
        }
    }
    // Initialize if empty
    localStorage.setItem('rujukanKeluar', JSON.stringify(defaultRujukanKeluar));
    return defaultRujukanKeluar;
}

/**
 * Save Rujukan Keluar to localStorage
 * @param {any[]} data
 */
export function saveRujukanKeluar(data) {
    if (!browser) return;
    localStorage.setItem('rujukanKeluar', JSON.stringify(data));
}

/**
 * Authenticate user
 * @param {string} username
 * @param {string} password
 */
export function login(username, password) {
    if (!browser) return { success: false, message: 'Server error' };
    
    // Check default credentials
    if (username === '12345' && password === '12345') {
        localStorage.setItem('isLoggedIn', 'true');
        const defaultProfile = { name: 'Dr. Zaldhy Masud', username: '12345', role: 'Dokter Umum' };
        localStorage.setItem('currentUser', JSON.stringify(defaultProfile));
        return { success: true, user: defaultProfile };
    }
    
    // Check registered users
    const registered = localStorage.getItem('registeredUsers');
    if (registered) {
        try {
            /** @type {any[]} */
            const users = JSON.parse(registered);
            const matched = users.find((/** @param {any} u */ u) => u.username === username && u.password === password);
            if (matched) {
                localStorage.setItem('isLoggedIn', 'true');
                const profile = { name: matched.name, username: matched.username, role: 'Pengguna Terdaftar' };
                localStorage.setItem('currentUser', JSON.stringify(profile));
                return { success: true, user: profile };
            }
        } catch (e) {
            // ignore
        }
    }
    
    return { success: false, message: 'Username atau password salah!' };
}

/**
 * Register user
 * @param {string} name
 * @param {string} username
 * @param {string} password
 */
export function register(name, username, password) {
    if (!browser) return { success: false, message: 'Server error' };
    
    if (username === '12345') {
        return { success: false, message: 'Username "12345" sudah digunakan sebagai akun bawaan.' };
    }
    
    /** @type {any[]} */
    let users = [];
    const registered = localStorage.getItem('registeredUsers');
    if (registered) {
        try {
            users = JSON.parse(registered);
        } catch (e) {
            users = [];
        }
    }
    
    // Check if username already exists
    if (users.find((/** @param {any} u */ u) => u.username === username)) {
        return { success: false, message: 'Username sudah terdaftar!' };
    }
    
    users.push({ name, username, password });
    localStorage.setItem('registeredUsers', JSON.stringify(users));
    return { success: true };
}

// Logout user
export function logout() {
    if (!browser) return;
    localStorage.setItem('isLoggedIn', 'false');
    localStorage.removeItem('currentUser');
}
