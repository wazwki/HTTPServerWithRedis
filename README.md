# How to config redis for in-memory storage or cache in golang HTTP-server


## Redis config: 

### Step 1:
#### Install and import client for connect to Redis:
```bash
go get github.com/go-redis/redis/v9
```
```go
import (
    "github.com/go-redis/redis/v9"
)
```

### Step 2:
#### Create connect object with options:
```go
rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})
```

### Step 3:
#### Using connect object methods:

#### Set: Sets the value for the key.
```go
err := rdb.Set(ctx, "myKey", "Hello, Redis!", 0).Err()
if err != nil {
    panic(err)
}
```
#### Get: Retrieves the value of a key.
```go
val, err := rdb.Get(ctx, "myKey").Result()
if err != nil {
    panic(err)
}
```

#### Del: Deletes a key.
```go
err = rdb.Del(ctx, "myKey").Err()
if err != nil {
    panic(err)
}
```

#### Exists: Checks for the presence of a key.

```go
exists, err := rdb.Exists(ctx, "myKey").Result()
if err != nil {
    panic(err)
}
if exists == 1 {
    fmt.Println("'myKey' exists in Redis")
} else {
    fmt.Println("'myKey' does not exist in Redis")
}
```

#### Expire: Sets the lifetime for the key.

```go
err = rdb.Expire(ctx, "myKey", 10*time.Second).Err()
if err != nil {
    panic(err)
}
```


## Redis for in-memory storage: 

### Step 1:
//
```go
```

### Step 2:
//
```go
```

### Step 3:
//
```go
```

### Step 4:
//
```go
```


## Redis for cache: 

### Step 1:
//
```go
```

### Step 2:
//
```go
```

### Step 3:
//
```go
```

### Step 4:
//
```go
```