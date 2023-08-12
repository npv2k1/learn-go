# Vòng lặp

**Vòng lặp trong Ngôn ngữ Go**

Trong ngôn ngữ lập trình Go, vòng lặp là một cấu trúc cho phép bạn thực hiện một khối mã lặp đi lặp lại cho đến khi một điều kiện nào đó không còn đúng nữa. Go cung cấp hai loại vòng lặp chính: `for` và `range`.

**Vòng lặp `for`:**

Vòng lặp `for` trong Go có thể được sử dụng để thực hiện một khối mã lặp lại cho đến khi một điều kiện được thỏa mãn hoặc cho một số lần lặp cụ thể.

Ví dụ về vòng lặp `for`:
```go
// In ra các số từ 1 đến 5
for i := 1; i <= 5; i++ {
    fmt.Println(i)
}

// Lặp vô hạn sử dụng vòng lặp vô điều kiện
for {
    // Thực hiện mã lệnh
}
```

**Vòng lặp `range`:**

Vòng lặp `range` được sử dụng để lặp qua các phần tử của một tập hợp (mảng, slice, map, channel) và trả về cặp giá trị (chỉ số và giá trị).

Ví dụ về vòng lặp `range`:
```go
numbers := []int{1, 2, 3, 4, 5}
for index, value := range numbers {
    fmt.Printf("Index: %d, Value: %d\n", index, value)
}

// Lặp qua các cặp key-value trong map
ages := map[string]int{
    "Alice": 25,
    "Bob":   30,
}
for name, age := range ages {
    fmt.Printf("%s is %d years old\n", name, age)
}
```

Trong cả hai loại vòng lặp, bạn có thể sử dụng câu lệnh `break` để thoát khỏi vòng lặp ngay lập tức hoặc `continue` để bỏ qua phần còn lại của lần lặp hiện tại và chuyển sang lần lặp tiếp theo.

Vòng lặp trong Go giúp bạn thực hiện các tác vụ lặp lại một cách dễ dàng và tiết kiệm mã nguồn. Tùy thuộc vào tình huống cụ thể, bạn có thể chọn sử dụng vòng lặp `for` hoặc `range` để đáp ứng yêu cầu của mình.