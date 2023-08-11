# Struct

Struct trong Golang là một cấu trúc dữ liệu phức tạp được sử dụng để định nghĩa một tập hợp các trường dữ liệu có liên quan. Struct có thể được sử dụng để tạo các kiểu do người dùng xác định. Struct là một kiểu kết hợp có nghĩa là nó có thể có các thuộc tính khác nhau và mỗi thuộc tính có thể có kiểu và giá trị riêng của chúng²⁴.

Để định nghĩa một struct, ta dùng từ khóa **type** sau đó là tên của struct và danh sách các trường trong cặp ngoặc nhọn. Ví dụ:

```go
type Person struct {
    FirstName string
    LastName string
    Age int
}
```

Để khởi tạo một biến struct, ta có thể sử dụng cú pháp như sau:

```go
var p Person // khai báo p có kiểu dữ liệu là Person, các trường của biến p sẽ có giá trị mặc định
p1 := Person{"A", "Nguyễn Văn", 26} // khai báo p1 có kiểu dữ liệu là Person, các trường có giá trị theo thứ tự khai báo
p2 := Person{FirstName: "John", LastName: "Snow", Age: 45} // khai báo p2 có kiểu dữ liệu là Person, các trường có giá trị được truyền theo key
p3 := Person{FirstName: "Robert"} // khai báo p3 có kiểu dữ liệu là Person, các trường FirstName có giá trị "Robert" các trường khác mang giá trị mặc định
```

Để truy cập vào các trường trong struct, ta dùng dấu chấm **.** sau tên biến. Ví dụ:

```go
fmt.Println("LastName :", p.LastName) // in ra LastName của biến p
fmt.Println("FirstName :", p.FirstName) // in ra FirstName của biến p
```

Struct cũng có thể được định nghĩa bên trong struct khác để thể hiện tính kế thừa. Ví dụ:

```go
type Student struct {
    Person // Student kế thừa từ Person
    Class string
    Point float
}
```

Đây là một số thông tin cơ bản về struct trong Golang. Bạn có thể tìm hiểu thêm về cách sử dụng phương thức, struct tag, interface với struct qua các nguồn tham khảo sau:

¹: [Struct trong GOLang - w3seo Cấu trúc trong ngôn ngữ Go](https://bing.com/search?q=struct+trong+golang)
²: [[Go Lang] lập trình golang căn bản - Struct trong golang - Viblo](https://viblo.asia/p/go-lang-lap-trinh-golang-can-ban-struct-trong-golang-aWj536JGl6m)
³: [Defining Structs in Go | DigitalOcean](https://www.digitalocean.com/community/tutorials/defining-structs-in-go)
⁴: [Struct trong GOLang - w3seo Cấu trúc trong ngôn ngữ Go](https://websitehcm.com/struct-trong-golang/)

Source: Conversation with Bing, 6/29/2023
(1) [Go Lang] lập trình golang căn bản - Struct trong golang - Viblo. https://viblo.asia/p/go-lang-lap-trinh-golang-can-ban-struct-trong-golang-aWj536JGl6m.
(2) Struct trong GOLang - w3seo Cấu trúc trong ngôn ngữ Go. https://websitehcm.com/struct-trong-golang/.
(3) Struct trong GOLang - w3seo Cấu trúc trong ngôn ngữ Go. https://bing.com/search?q=struct+trong+golang.
(4) Defining Structs in Go | DigitalOcean. https://www.digitalocean.com/community/tutorials/defining-structs-in-go.
