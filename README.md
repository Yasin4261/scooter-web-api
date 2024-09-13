# Scooter Rental API

Bu proje, scooter kiralama işlemlerini yönetmek için bir RESTful API'dir. Go dili ve MongoDB kullanılarak geliştirilmiştir.

## Özellikler

- Tüm scooter'ları listeleme
- Yeni scooter oluşturma
- Scooter konumunu güncelleme
- Scooter durumunu güncelleme
- Kiralama işlemleri (başlatma ve tamamlama)
- Kullanıcı kiralamalarını listeleme

## Kurulum

1. **Projeyi Klonlayın:**

   ```bash
   git clone https://github.com/yourusername/scoter-web-api.git
   cd scoter-web-api

2. **Bağımlılıkları Yükleyin:**
```bash
go mod download
```
   
4. **MongoDB'yi Başlatın:** MongoDB'yi kurun ve başlatın.

5. **Uygulamayı Başlatın:**
```bash
go run cmd/api/main.go
```

Uygulama varsayılan olarak `http://localhost:8080` adresinde çalışacaktır.


## API Kullanımı

### Scooter API
- **GET** `/scooters:` Tüm scooter'ları listeleyin.
- **POST** `/scooters:` Yeni scooter ekleyin.
- **PATCH** `/scooters/{id}/location:` Scooter konumunu güncelleyin.
- **PATCH** `/scooters/{id}/status:` Scooter durumunu güncelleyin.

### Rental API
- **POST** `/rentals:` Kiralama işlemini başlatın.
- **GET** `/rentals/user/{user_id}:` Kullanıcının kiralamalarını listeleyin.
- **PATCH** `/rentals/{id}/complete:` Kiralamayı tamamlayın.

### Katkıda Bulunanlar
[Yasin Güneş](https://github.com/Yasin4261)
## Lisans
Bu proje MIT Lisansı altında lisanslanmıştır. Daha fazla bilgi için `LICENSE` dosyasına bakın.

