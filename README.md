# Go Simple API

This is a simple short REST API wrote with Golang using features as use-cases tests, JWT, smtp.SendMail and Postgres / SQLite (or any SQL database surported by GORM).  
The following packages were used to build this application:

**Gorm:** An ORM library for Golang to deal with SQL databases  
**sqlite:** Drive used to deal with sqlite3 databases  
**Crypto:** Collects common cryptographic constants, commonly used on auth module to deal with passwords  
**Govalidator:** A package of validators and sanitizers for strings, structs and collections  
**Testify:** A toolkit with common assertions and mocks that plays nicely with the standard library  
**satori/go.uuid** A package to generate UUID strings to be used on ids  
**Pq:** Postgres driver to deal with Postgres databases  
**Mux:** A HTTP router and URL matcher for building Go web servers  
**jwt-go:** A friendily package to deal with JSON Web Tokens  
**faker:** A package to generate fake data, used on tests to generate fake users  
