---

## 🧠 **1. HTTP Bersifat Stateless**
- HTTP tidak menyimpan *state* antar request.
- Karena itu, server tidak otomatis “ingat” siapa kamu → butuh sistem seperti cookie/session.
- Keuntungan: mudah di-*scale* karena server tidak menyimpan memori user.

---

## 🍪 **2. Cookie di Golang**
- Cookie adalah cara menyimpan data kecil di browser agar dikirim kembali ke server pada setiap request.
- Golang menggunakan `http.Cookie` untuk membuat cookie, dan `http.SetCookie()` untuk mengirimnya ke browser.
- Contoh:
  ```go
  cookie := http.Cookie{Name: "user", Value: "nabiel", Path: "/", MaxAge: 3600}
  http.SetCookie(w, &cookie)
  ```

---

## 🎯 **3. Pointer dan Struct**
- Pointer (`*`) menunjuk ke alamat memori suatu data.
- `new(http.Cookie)` membuat pointer ke struct kosong.
- `cookie := http.Cookie{}` langsung membuat struct (bukan pointer).
- `http.SetCookie()` butuh `*http.Cookie` → makanya sering pakai pointer.
- Tapi tiap request handler dijalankan dalam Goroutine terpisah, jadi tidak saling mengganggu antar user.

---

## 👤 **4. Isolasi Antar User**
- Setiap request HTTP diproses secara **terpisah** (isolated), jadi pointer/objek dalam handler hanya milik request itu.
- Cookie disimpan di **browser user**, bukan di server.
- User A dan B punya cookie masing-masing, server hanya membaca dan merespon cookie sesuai user yang mengirim request.

---

## 🧰 **5. Middleware**
- Middleware adalah fungsi perantara sebelum handler utama.
- Bisa dipakai untuk:
  - Autentikasi (cek login)
  - Logging
  - Validasi role (otorisasi)
- Contoh middleware auth:
  ```go
  func AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
      return func(w http.ResponseWriter, r *http.Request) {
          // Cek login, kalau nggak redirect
          next(w, r)
      }
  }
  ```

---

## 🔐 **6. Authentication vs Authorization**
| Konsep         | Penjelasan                                       |
|----------------|--------------------------------------------------|
| Authentication | Verifikasi identitas user (siapa kamu)           |
| Authorization  | Cek hak akses user (boleh ngapain)               |
| Login/Register/Logout | Bagian dari sistem Authentication         |

---

## 🧩 **7. Fitur Login, Daftar, Logout**
- Disebut **Authentication System** atau **Auth Feature**
- Meliputi:
  - Pendaftaran (register)
  - Login (dengan validasi)
  - Logout (hapus session/cookie)
  - Middleware untuk melindungi route

---

Kalau kamu ingin membuat **auth system sederhana di Golang** (register, login, logout + middleware), aku bisa bantu buatkan struktur kodenya. Mau lanjut ke situ?
