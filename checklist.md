### Back-end

#### pkg/db

- [ ] sqlite.go - Finalize InitDB() function (db name)
- [ ] sqlite.go - Add function for creating tables if not exist
- [ ] sqlite.go - delete temporary database function

#### pkg/handlers

- [ ] posts.go - on saving new post also save 'ALMOST_PRIVATE' list

#### To-do

- [ ] Create new group route
- [ ] For /follow route - save notification for private profile
- [ ] Deal with pending follow requests after status change to PUBLIC

#### Websockets

- [ ] Plan out notification table (db)
- [ ] Plan chat message handling

