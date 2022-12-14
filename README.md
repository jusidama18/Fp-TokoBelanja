# Fp-TokoBelanja


## Pembagian Tugas

| Bagian     | Detail             | Dikerjakan oleh          |
|------------|--------------------|--------------------------|
| Endpoint   | POST /users/register        | Dion Fauzi |
| Endpoint   | POST /users/login        | Dion Fauzi |
| Endpoint   | PATCH /users/topup        | Dion Fauzi |
| Endpoint   | POST /users/admin        | Dion Fauzi |
| Endpoint   | POST /categories        | Dion Fauzi |
| Endpoint   | GET /categories        | Dion Fauzi |
| Endpoint   | PATCH /categories/:id        | Dion Fauzi |
| Endpoint   | DELETE /categories/:id        | Dion Fauzi |
| Endpoint   | GET /products        | Juan Simon Damanik |
| Endpoint   | POST /products       | Juan Simon Damanik |
| Endpoint   | PUT /products/:id    | Juan Simon Damanik |
| Endpoint   | DELETE /products/:id | Juan Simon Damanik |
| Endpoint   | POST /transactions/        | Muhammad Rifqi Al Furqon |
| Endpoint   | GET /transactions/my-transactions       | Muhammad Rifqi Al Furqon |
| Endpoint   | GET /transactions/user-transactions    | Muhammad Rifqi Al Furqon|



## Deployment
Projek dideploy di Railway dengan link berikut [https://fp-tokobelanja-production-b9b7.up.railway.app/](https://fp-tokobelanja-production-b9b7.up.railway.app/)

## How to Run
### Locally
- Clone this repo
```
git clone https://github.com/DionFauzi/fp-Fp-TokoBelanja
```
- Run PostgreSQL Docker script
```
chmod +x ./scripts/run-postgres.sh && ./scripts/run-postgres.sh
```
- Copy .env.example to .env
```
cp .env.example .env
```
- Run go webserver
```
go run ./main.go
```
- Enjoy!
