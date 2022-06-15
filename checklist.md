### Back-end

#### server.go

- [x] Replace test routes

#### pkg/db

- [ ] sqlite.go - Finalize InitDB() function (db name)
- [ ] sqlite.go - Add function for creating tables if not exist
- [x] users.go - Delete Test function
- [ ] sqlite.go - delete temporary database function

#### pkg/handlers

- [x] Middleware - check functionality after real request from front-end
- [x] Middleware - Handle errors and responses
- [x] register.go - delete test handlers
- [x] register.go - handle errors and responses
- [x] user.go - delete temp function - Add(User)error

#### pkg/utils

- [x] session.go - check if cookie gets sent to front end
- [x] session.go - check if cookie gets destroyed when conn to front end

### Front-end

### Register.vue

- [x] Added "credentials" to allow sending cookie
- [x] Refactored form sending
- [x] Deleted headers - https://muffinman.io/blog/uploading-files-using-fetch-multipart-form-data/
