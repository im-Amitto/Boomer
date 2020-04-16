## Boomer

### Prerequisite

- Npm
- Go
- MongoDB

### Configuration

- The port to be used can be customized from `/server/main.go`
- The mongodb connection url is to be updated in `/server/middleware/middleware.go:21`

### How to run?

```
  npm run build
  cd server
  go main.go
```

You can see your application at `htttp://localhost:{port-number}/` by default it's `3200`.

### Postman collection

Postman collection can be found at `postman/Boomer-API.postman_collection.json`

### App demonstration
<video width="320" height="240" controls>
  <source src="https://www.loom.com/share/08bf6c62a6a84554ab7c19840f6c7285" type="video/mp4">
</video>