# MaidNovel - AI-Powered Novel Refiner

MaidNovel adalah aplikasi web berbasis Go (Fiber) dan HTMX yang dirancang untuk mengambil konten novel dari web scraper dan menyempurnakan kualitas teks Bahasa Indonesia menggunakan AI (OpenRouter/Gemini). Fokus utama aplikasi ini adalah mengubah terjemahan kaku menjadi prosa yang luwes, natural, dan berkualitas literasi profesional.

## ✨ Fitur Utama

- **AI Prose Refinement**: Menggunakan `Gemini 2.0 Flash` via OpenRouter untuk memoles teks agar lebih mengalir dan natural.
- **Automated Scraping**: Sinkronisasi daftar chapter dan konten secara otomatis dari sumber novel (Meionovels).
- **Premium UI/UX**: Tampilan minimalis dengan animasi skeleton loader saat memuat data.
- **Dark/Light Mode**: Dukungan tema gelap dan terang dengan persistensi pilihan tema.
- **Modular Architecture**: Struktur kode yang rapi dan mudah dikelola (Standard Go Project Layout).
- **Fast & Lightweight**: Dibangun dengan Go Fiber dan HTMX untuk performa maksimal tanpa overload JavaScript.

## 🛠️ Tech Stack

- **Backend**: Go (Golang)
- **Web Framework**: [Fiber](https://gofiber.io/)
- **Frontend**: HTMX, Vanilla CSS, Vanilla JS
- **Database**: SQLite3
- **AI Engine**: OpenRouter (Gemini 2.0 Flash)

## 🚀 Persiapan & Instalasi

### Prasyarat
- [Go](https://go.dev/doc/install) (versi 1.25 atau terbaru direkomendasikan)
- [OpenRouter API Key](https://openrouter.ai/)

### Langkah Instalasi

1. **Clone Repository**
   ```bash
   git clone https://github.com/myusuf27/novel_translate_indonesia.git
   cd novel_translate_indonesia
   ```

2. **Setup Environment Variables**
   Buat file `.env` di root directory:
   ```bash
   touch .env
   ```
   Isi dengan API Key Anda:
   ```env
   OPENROUTER_API_KEY=sk-or-v1-xxxx...
   ```

3. **Install Dependencies**
   ```bash
   go mod tidy
   ```

4. **Jalankan Aplikasi**
   ```bash
   go run ./cmd/server/main.go
   ```
   Aplikasi akan berjalan di `http://localhost:3000`.

## 📂 Struktur Proyek

```text
.
├── cmd/
│   └── server/          # Entry point aplikasi
├── internal/
│   ├── database/        # Inisialisasi database
│   ├── handler/         # Controller/Route handler (Fiber)
│   ├── models/          # Definisi struktur data
│   ├── repository/      # Akses data (SQL Queries)
│   └── service/         # Logika bisnis (Scraper & AI)
├── static/
│   ├── css/             # Styling (Modern Vanilla CSS)
│   └── js/              # Client-side scripts (Theme toggle, etc)
├── views/               # HTML Templates (Fiber HTML Engine)
└── .env                 # Konfigurasi kunci API (tidak masuk git)
```

## 📝 Lisensi
Proyek ini dibuat untuk tujuan edukasi dan penggunaan pribadi.

---
Dibuat dengan ❤️ untuk komunitas pembaca novel Indonesia.
