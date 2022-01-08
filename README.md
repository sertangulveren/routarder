### Routarder

Basic and experimental routing helper.

#### Usage:


```go
package main

import (
    "fmt"
    "github.com/sertangulveren/routarder"
    "net/http"
)

func main() {
    router := routarder.New()
    
    router.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "welcome to home page\n")
    })
    
    router.HandleFunc("/help", HelpHandler)
    
    api = router.Group("/api")
    api.HandleFunc("/posts", PostsHandler) 
    api.HandleFunc("/users", UsersHandler)
    
    routarder.Serve(":3000")
}
```
