# Boomer

## Introduction 

### Problem statement - A small website reporting stolen cars


## Product Requirement:

- Car owners can report stolen cars

- New stolen car cases should be automatically assigned to any free police officer

- A police officer can only handle one stolen car case at a time

- When the Police find a car, the case is marked as resolved and the responsible police officer become available to take a new stolen car case

- The system should be able to assign unassigned stolen car cases automatically when a police officer becomes available


### Other Goals:

- Your task is to provide APIs for a frontend application that satisfies all above requirement and also the frontend for implementing this feature

- Please stick to the Product Requirements. Authentication is not required 


## Deployment
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