# **How to Set Up Redis in a Golang HTTP Server**

# **1. Install and Import the Redis Client:**

```bash
go get github.com/redis/go-redis/v9
```

```go
import (
    "github.com/redis/go-redis/v9"
)
```

# **2. Create a Connection Object with Parameters:**

```go
rdb := redis.NewClient(&redis.Options{
    Addr:     "localhost:6379",
    Password: "",
    DB:       0,
})
```

# **3. Using Connection Object Methods:**

## 3.1. Set: Sets the value of a key.

```go
err := rdb.Set(ctx, "myKey", "Hello, Redis!", 0).Err()
if err != nil {
    panic(err)
}
```

## 3.2. Get: Retrieves the value of a key.

```go
val, err := rdb.Get(ctx, "myKey").Result()
if err != nil {
    panic(err)
}
```

## 3.3. Del: Deletes a key.

```go
err = rdb.Del(ctx, "myKey").Err()
if err != nil {
    panic(err)
}
```

## 3.4. Exists: Checks the existence of a key.

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

## 3.5. Expire: Sets the expiration time for a key.

```go
err = rdb.Expire(ctx, "myKey", 10*time.Second).Err()
if err != nil {
    panic(err)
}
```