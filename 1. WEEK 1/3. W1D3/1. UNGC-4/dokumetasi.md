Tentu, berikut adalah contoh Markdown lengkap untuk dokumentasi API Anda, bersama dengan penjelasan setiap bagian:

```markdown
# API Documentation

Ini adalah dokumentasi API untuk [nama proyek Anda]. Dokumen ini memberikan panduan untuk menggunakan API ini, termasuk rute-rute yang tersedia, metode HTTP yang digunakan, parameter yang dibutuhkan, dan contoh permintaan serta respons.

## Rute API

### Create a Criminal Report

**URL:** `/criminal-reports`

**Method:** `POST`

Membuat laporan kejadian kriminal baru. Anda perlu mengirimkan data dalam format JSON.

**Request:**
```json
{
  "hero_id": 2,
  "villain_id": 3,
  "description": "Deskripsi kejadian kriminal",
  "event_time": "2023-10-29 11:30:00"
}
```

**Response:**
```json
{
  "id": 1,
  "hero_id": 2,
  "villain_id": 3,
  "description": "Deskripsi kejadian kriminal",
  "event_time": "2023-10-29 11:30:00"
}
```

### Get a Criminal Report

**URL:** `/criminal-reports/{id}`

**Method:** `GET`

Mengambil informasi tentang laporan kejadian kriminal berdasarkan ID yang diberikan.

**Response:**
```json
{
  "id": 1,
  "hero_id": 2,
  "villain_id": 3,
  "description": "Deskripsi kejadian kriminal",
  "event_time": "2023-10-29 11:30:00"
}
```

### Update a Criminal Report

**URL:** `/criminal-reports/{id}`

**Method:** `PUT`

Memperbarui laporan kejadian kriminal berdasarkan ID yang diberikan. Anda perlu mengirimkan data dalam format JSON.

**Request:**
```json
{
  "hero_id": 4,
  "villain_id": 5,
  "description": "Deskripsi kejadian kriminal yang diperbarui",
  "event_time": "2023-11-01 15:45:00"
}
```

**Response:**
```
Criminal Report berhasil diperbarui
```

### Delete a Criminal Report

**URL:** `/criminal-reports/{id}`

**Method:** `DELETE`

Menghapus laporan kejadian kriminal berdasarkan ID yang diberikan.

**Response:**
```
Criminal Report berhasil dihapus
```

## Panduan Pengguna

Untuk menggunakan API ini, Anda dapat mengirimkan permintaan HTTP sesuai dengan rute dan metode yang tersedia. Pastikan untuk menyertakan data yang diperlukan dalam format JSON. Anda akan menerima respons yang sesuai setelah setiap operasi.

## Kesalahan

- Jika laporan kejadian kriminal tidak ditemukan, Anda akan menerima respons `404 Not Found`.
- Jika terjadi kesalahan saat memproses permintaan, Anda akan menerima respons `500 Internal Server Error`.

## Contoh Penggunaan

### Contoh Permintaan Create:

```bash
curl -X POST -H "Content-Type: application/json" -d '{
  "hero_id": 2,
  "villain_id": 3,
  "description": "Deskripsi kejadian kriminal",
  "event_time": "2023-10-29 11:30:00"
}' http://example.com/criminal-reports
```

### Contoh Permintaan Update:

```bash
curl -X PUT -H "Content-Type: application/json" -d '{
  "hero_id": 4,
  "villain_id": 5,
  "description": "Deskripsi kejadian kriminal yang diperbarui",
  "event_time": "2023-11-01 15:45:00"
}' http://example.com/criminal-reports/1
```

## Catatan

Pastikan untuk menggunakan tautan, header, dan data yang sesuai saat mengirim permintaan ke API ini. Ikuti panduan ini dengan cermat untuk berinteraksi dengan API dengan benar.

Terima kasih telah menggunakan API [Avenger Criminal Report]!
```

