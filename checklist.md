### Back-end

#### server.go

- [ ] Replace test routes

#### pkg/db

- [ ] sqlite.go - Finalize InitDB() function (db name)
- [ ] sqlite.go - Add function for creating tables if not exist
- [ ] users.go - Delete Test function

#### pkg/handlers

- [ ] Middleware - check functionality after real request from front-end
- [ ] Middleware - Handle errors and responses
- [ ] register.go - delete test handlers
- [ ] register.go - handle errors and responses
- [ ] user.go - delete temp function - Add(User)error

#### pkg/utils

- [ ] session.go - check if cookie gets sent to front end
- [ ] session.go - check if cookie gets destroyed when conn to front end
