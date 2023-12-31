# gofiber-mongodb-rest-api

GoMongoDB API, Go dilinde geliştirilmiş ve MongoDB veritabanını kullanan basit bir API projesidir. 

Fiber framework'ü ile oluşturulmuş bu proje, veritabanı işlemleri için MongoDB'yi kullanır ve kullanıcılara ürünler oluşturma, görüntüleme ve silme gibi temel CRUD (Create, Read, Update, Delete) operasyonlarını gerçekleştirebilecekleri bir RESTful API sunar.

## Kurulum

1. **MongoDB Kurulumu:** Öncelikle, projeyi çalıştırmadan önce MongoDB veritabanını kurmanız ve çalışır durumda olduğundan emin olmanız gerekmektedir.

2. **Proje Kurulumu:**

```bash
   git clone https://github.com/mustafadikyar/gofiber-mongodb-rest-api.git
   cd gofiber-mongodb-rest-api
   go run main.go
```
Proje bu adımlarla yerel bir sunucuda çalıştırılacaktır.

## API Endpoints
  
* *Ürün Oluşturma:*

```http
POST

  /api/product
```

* *Tüm Ürünleri Görüntüleme:*

```http
GET

  /api/products
```

* *Ürün Silme:*

```http
DELETE

  /api/product/{id}
```
