# Biến

**Biến và Hằng số trong Ngôn ngữ Go**

Trong ngôn ngữ lập trình Go, biến và hằng số là những yếu tố cơ bản cho việc lưu trữ dữ liệu và quản lý giá trị trong chương trình của bạn.

**Biến (Variables):**

Biến là tên tham chiếu đến vùng bộ nhớ lưu trữ giá trị. Trong Go, bạn khai báo biến bằng cách sử dụng từ khóa `var` và sau đó là tên biến, kiểu dữ liệu và giá trị khởi tạo (nếu có).

Ví dụ:
```go
var age int // Khai báo biến 'age' kiểu int
age = 30    // Gán giá trị 30 cho biến 'age'

var name string = "John" // Khai báo và khởi tạo biến 'name' kiểu string
```

Go cũng cho phép viết ngắn gọn hơn cho khai báo biến sử dụng toán tử `:=`, trong đó kiểu dữ liệu được suy ra từ giá trị khởi tạo:

```go
count := 10 // Khai báo và khởi tạo biến 'count' kiểu int
```

**Hằng số (Constants):**

Hằng số là các giá trị không thay đổi trong suốt quá trình thực thi chương trình. Trong Go, bạn khai báo hằng số bằng cách sử dụng từ khóa `const`.

Ví dụ:
```go
const pi = 3.14159 // Khai báo hằng số 'pi' có giá trị là 3.14159

const (
    daysInWeek = 7
    monthsInYear = 12
)
```

Hằng số trong Go chỉ có thể chứa các giá trị cố định, chẳng hạn như số, chuỗi hoặc boolean. Bạn không thể gán giá trị cho một hằng số sau khi nó đã được khai báo.

**Kiểu Dữ liệu:**

Go có một loạt các kiểu dữ liệu như `int`, `float64`, `string`, `bool`, v.v. Kiểu dữ liệu giúp xác định loại dữ liệu mà biến hoặc hằng số có thể chứa.

Ví dụ:
```go
var temperature float64 = 98.6
var isStudent bool = true
```

Trong ngôn ngữ Go, biến và hằng số chơi một vai trò quan trọng trong việc quản lý dữ liệu và thực hiện tính toán. Sử dụng chúng một cách thông minh sẽ giúp mã của bạn dễ đọc, dễ bảo trì và dễ mở rộng.