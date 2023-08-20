# Stdin trong go

```go
package main

import (
	"fmt"
)

func main() {
	// In ra màn hình dòng thông báo yêu cầu người dùng nhập tên.
	fmt.Println("Please enter your name.")
	
	// Khai báo biến name kiểu string để lưu trữ tên người dùng nhập.
	var name string
	
	// Sử dụng hàm fmt.Scanln để đọc dữ liệu từ người dùng và lưu vào biến name.
	// Đọc dữ liệu từ bàn phím và gán vào biến name, cho phép người dùng nhập tên của họ.
	fmt.Scanln(&name)
	
	// In ra màn hình dòng thông báo chào mừng người dùng với tên vừa nhập.
	fmt.Printf("Hi, %s! I'm Go!", name)
}
```

**Giải thích:**

1. `package main`: Dòng đầu tiên trong một tệp Go xác định rằng bạn đang làm việc trong gói (package) có tên là "main". Gói "main" đặc biệt trong Go vì nó là gói mà chương trình chính (hàm `main`) sẽ chạy.

2. `import ("fmt")`: Dòng này import gói "fmt" (format) từ thư viện chuẩn của Go. Gói "fmt" cung cấp các chức năng để định dạng và hiển thị dữ liệu.

3. `func main() { ... }`: Đây là hàm chính của chương trình. Tất cả mã lệnh trong chương trình Go bắt đầu từ hàm `main`. Mọi thứ bên trong cặp dấu ngoặc nhọn `{}` là phần thân của hàm `main`.

4. `fmt.Println("Please enter your name.")`: Dòng này sử dụng hàm `fmt.Println` để in ra màn hình thông báo yêu cầu người dùng nhập tên.

5. `var name string`: Dòng này khai báo biến `name` có kiểu dữ liệu là `string` để lưu trữ tên người dùng sắp nhập.

6. `fmt.Scanln(&name)`: Dòng này sử dụng hàm `fmt.Scanln` để đọc dữ liệu mà người dùng nhập từ bàn phím và lưu vào biến `name`.

7. `fmt.Printf("Hi, %s! I'm Go!", name)`: Dòng này sử dụng hàm `fmt.Printf` để in ra màn hình thông báo chào mừng người dùng với tên vừa nhập. Chuỗi `%s` sẽ được thay thế bằng giá trị của biến `name`.

Như vậy, đoạn code trên thực hiện việc yêu cầu người dùng nhập tên, sau đó chào mừng người dùng bằng tên họ vừa nhập, và sử dụng gói "fmt" để thực hiện các tác vụ này.