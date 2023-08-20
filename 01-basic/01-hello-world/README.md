# Hello world

`main.go`

```go
package main

import "fmt"

func main() {
	fmt.Print("Hello, World!")
}
```

Để chạy code go ta dùng lệnh `go run`:

```bash
$ go run main.go
Hello, World!
```

Trong đoạn code trên ta sử dụng thư viên `fmt` và gọi tới phương thức `Print` để in ra màn hình chuỗi `Hello, World!`.