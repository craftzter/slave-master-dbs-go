## Directory structure

```
/cmd      --> entry point
/internal --> internal workflow
  /entity --> Bussines domain model
  /usecase --> Interface and main bussines logic
  /repository --> Interface or abstraction of data layer
  /infrastructure --> Implement of data layer interface
  /delivery --> handler of usecase
/pkg      --> helper and utililty
/migrations --> database migration
```

## API Endpoints

### Authentication
- `POST /register` - Register user (JSON: username, email, password)
- `POST /login` - Login user (JSON: email, password) -> returns JWT token

### Profile
- `POST /profile` - Create profile (Auth required, JSON: bio, avatar_url)
- `GET /profile/:userID` - Get profile by user ID
- `PUT /profile` - Update profile (Auth required, JSON: bio, avatar_url)

### Posts
- `POST /posts` - Create post (Auth required, Multipart: title, content, image[file])
- `GET /posts` - Get all posts (Public)
- `GET /posts/:id` - Get post by ID (Auth required, ownership check)
- `PUT /posts/:id` - Update post (Auth required, ownership check, Multipart: title, content, image[file])
- `DELETE /posts/:id` - Delete post (Auth required, ownership check)

### Static Files
- `GET /uploads/*filepath` - Serve uploaded images

## Setup
1. Start Docker DB: `make docker-up` (starts PostgreSQL master-slave)
2. Run migrations: `make migrate-up-primary`
3. Generate sqlc: `make sqlc-gen`
4. Build and run: `make build && make run`

## Features
- Clean Architecture
- PostgreSQL Master-Slave Replication
- JWT Authentication
- File Upload for Posts
- Ownership-based Access Control
