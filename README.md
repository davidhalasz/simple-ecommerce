## A simple ecommerce web application

Practicing to build a secure ecommerce application with golang. This app based on Trevor Sawler's project from [Udemy course](https://www.udemy.com/course/building-web-applications-with-go-intermediate-level/?couponCode=KEEPLEARNING)

#### GOLANG version
- go 1.21.4

#### USAGE
1. install make 
2. install [Buffalo](https://gobuffalo.io/documentation/getting_started/installation/)
3. run:
```
soda migrate
```
4. Copy your Stripe keys, smtp data and edit ports if you needed in Makefile.example (remove .example extension)
5. Run te app:
```
make start
```

#### Todos

- [x] Build frontend using GO
- [x] Build backend API using GO
- [x] Stripe credit card processing implementation
- [x] Connecting to MySQL
- [x] Build database with Go Buffalo & FIZZ
- [x] Authentication 
- [x] Protecting Routes
- [ ] Build Dashboard 
- [ ] Admin Pages
- [ ] Sending Mails
- [ ] build Microservices (generating PDF)