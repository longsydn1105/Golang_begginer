package main

// 1. Đo hàm tính toán bình thường
// go test -bench=BenchmarkGetAdultYears -benchmem
// func BenchmarkGetAdultYears(b *testing.B) {
// 	// b.N là số lần lặp, Go tự quyết định (có thể là hàng triệu lần)
// 	for i := 0; i < b.N; i++ {
// 		getAdultYears(32)
// 	}
// }

// // 2. Đo hàm dùng con trỏ (Pointer)
// // go test -bench=BenchmarkPointer -benchmem
// func BenchmarkPointer(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		pointerLogic()
// 	}
// }
