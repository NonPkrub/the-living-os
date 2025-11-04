# URL Shortener API
ได้เลยครับ ผมช่วยปรับ README.md ภาษาไทย ให้ **รวมข้อสังเกต / ข้อติดขัด** ที่คุณเจอด้วย ✅
จะช่วย reviewer เข้าใจปัญหาและ workaround

````markdown
# URL Shortener API

โปรเจกต์ URL Shortener API พัฒนาโดย **Go Fiber** ใช้ **PostgreSQL** สำหรับเก็บข้อมูล พร้อม Docker & Docker Compose support  

---

## 🛠 คุณสมบัติหลัก

- สร้าง URL สั้น (`POST /shorten`)  
- Redirect ไปยัง URL ต้นฉบับ (`GET /:shortCode`)  
- แสดงสถิติการคลิก (`GET /stats/:shortCode`)  
- AutoMigrate database tables  
- Docker & Docker Compose support  
- Bonus (option): Custom short code, Input validation, Error handling  
````
---

## 📦 การติดตั้งและรัน

### 1. Clone โปรเจกต์
```bash
git clone <your-repo-url>
cd url-shortener


### 2. รันด้วย Docker Compose

```bash
docker-compose up --build
```

* `db` → PostgreSQL
* `app` → Go Fiber API (port 3000)



### 3. ข้อสังเกต / ข้อติดขัด

* **Volumes:**

  ```yaml
  volumes:
    # - ./data:/var/lib/postgresql/data
  ```

  การแมป volume แบบ local (`./data`) อาจทำให้ DB container มี permission issues บางครั้ง แนะนำ comment ไว้หรือใช้ Docker volume ปกติ (`pg_data`)

* **App ต้อง re-run:**
  ถ้าแก้โค้ด Go ต้อง rebuild / restart container app เพื่อให้การเปลี่ยนแปลงมีผล

* **Route redirect กับ Swagger:**
  เส้นทาง redirect (`GET /:shortCode`) จะ **ติด CORS** ใน Swagger UI

  * แนะนำทดสอบ redirect ด้วย browser หรือ curl แทน Swagger

---

## 🔧 ตัวแปร Environment

**ถ้าไม่ใช้ Docker:**

```env
DB_HOST=localhost
DB_USER: postgres
DB_PASSWORD: "12345"
DB_NAME: short-url
DB_PORT: 5432
```

**ใน Docker Compose:**

```yaml
environment:
    DB_HOST: db
    DB_USER: postgres
    DB_PASSWORD: "12345"
    DB_NAME: short-url
    DB_PORT: 5432
```

> แนะนำใช้ `TimeZone=UTC` เพื่อหลีกเลี่ยง error `unknown time zone`

---

## 🚀 API Endpoints

### 1. สร้าง URL แบบสั้น

```http
POST /shorten
Content-Type: application/json

{
  "originalUrl": "https://example.com",
  "customCode": "abc123" // default
}
```

**Response:**

```json
{
  "originalUrl": "https://example.com",
  "shortCode": "abc123",
  "shortUrl": "http://localhost:3000/abc123",
}
```

### 2. Redirect

```http
GET /:shortCode
```

* Redirect (301) ไป URL ต้นฉบับ
* เพิ่มจำนวนคลิกในฐานข้อมูล
* ⚠️ **ทดสอบผ่าน Swagger อาจติด CORS** → ใช้ browser / curl แทน

### 3. ดูสถิติ URL

```http
GET /stats/:shortCode
```

**Response:**

```json
{
  "shortCode": "abc123",
  "originalUrl": "https://example.com",
  "clicks": 5,
  "createdAt": "2025-11-05T12:34:56Z",
}
```

---

## 📝 ตัวอย่างใช้งาน (curl)

**สร้าง URL แบบสั้น**

```bash
curl -X POST http://localhost:3000/shorten \
-H "Content-Type: application/json" \
-d '{"originalUrl":"https://example.com","customCode":"abc123"}'
```

**Redirect**

```bash
curl -L http://localhost:3000/abc123
```

**ดูสถิติ**

```bash
curl http://localhost:3000/stats/abc123
```

---

## 💡 สมมติฐานและ trade-offs

* TimeZone ใช้ `UTC` เพื่อหลีกเลี่ยงปัญหา timezone ใน container
* ไม่ทำ unit test แบบเต็ม เพราะเวลาจำกัด (สามารถเพิ่มเป็น bonus)
* Docker Compose ใช้เพื่อให้ reviewer run ง่าย และ database auto-migrate
* Custom short code รองรับ แต่ยังไม่มีระบบตรวจสอบ duplicate ซับซ้อน
* Route redirect ติด CORS ใน Swagger → แนะนำทดสอบด้วย curl หรือ browser
* Production deploy: Assignment focus on local/dev environment → ไม่ deploy จริง
* ⚠️ จึง **ไม่มีการสร้าง `.env` หรือ `.env.example`**  
  Credentials ถูกใส่ตรง ๆ ใน `docker-compose.yml` เพื่อให้ reviewer run ได้ทันที (dev/test only)
* การ set-up ที่ผิดพลาดยังส่งผลต่อกับเวลาในการแก้ไขข้อผิดพลาด

## 📚 References / อ้างอิง

- Clean Architecture: [https://docs.mikelopster.dev/c/goapi-essential/chapter-7/clean-code]
