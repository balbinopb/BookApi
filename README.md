#Book API

Book API adalah RESTful API sederhana berbasis **Golang + Gin** untuk mengelola data **Buku** dan **Kategori**.  
Project ini menggunakan PostgreSQL sebagai database, dengan dukungan **migration** dan **Basic Authentication**.  
API sudah **live** di Railway

**Base URL Production:**  
[https://bookapi-production-de36.up.railway.app](https://bookapi-production-de36.up.railway.app)

---

## Daftar Endpoint

### Categories
- `GET    /api/categories` → Ambil semua kategori  
- `POST   /api/categories` → Tambah kategori baru  
- `GET    /api/categories/:id` → Ambil kategori berdasarkan ID  
- `DELETE /api/categories/:id` → Hapus kategori berdasarkan ID  
- `GET    /api/categories/:id/books` → Ambil semua buku pada kategori tertentu  

### Books
- `GET    /api/books` → Ambil semua buku  
- `POST   /api/books` → Tambah buku baru  
- `GET    /api/books/:id` → Ambil buku berdasarkan ID  
- `DELETE /api/books/:id` → Hapus buku berdasarkan ID  
- `PUT    /api/books/:id` → Update data buku
