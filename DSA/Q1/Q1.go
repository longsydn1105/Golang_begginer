package main

import "fmt"

func main() {
	input := ")ebc#da@f("
	fmt.Println("Input: ", input)
	fmt.Println("Output:", reverseLettersAndSpecials(input))
}

func reverseLettersAndSpecials(s string) string {
	n := len(s)
	// 1. Tạo 2 cái kho (Slice) để chứa riêng Chữ và Dấu
	// make([]byte, 0, n) giúp tối ưu bộ nhớ, tránh cấp phát lại nhiều lần
	letters := make([]byte, 0, n)
	specials := make([]byte, 0, n)

	// Duyệt lần 1: Phân loại "rác" và "hoa"
	for i := range n {
		char := s[i]
		// Check xem là chữ thường (a-z)
		if char >= 'a' && char <= 'z' {
			letters = append(letters, char)
		} else {
			specials = append(specials, char)
		}
	}

	// 2. Chuẩn bị ghép lại
	// Dùng mảng byte để lắp ráp cho nhanh
	result := make([]byte, n)

	// Mẹo giả lập Stack: Dùng index trỏ vào phần tử CUỐI CÙNG của slice
	lIdx := len(letters) - 1
	sIdx := len(specials) - 1

	// Duyệt lại chuỗi gốc để biết vị trí nào điền cái gì
	for i := 0; i < n; i++ {
		char := s[i]

		if char >= 'a' && char <= 'z' {
			// Chỗ này vốn là chữ -> Lấy chữ từ cuối kho letters đắp vào
			result[i] = letters[lIdx]
			lIdx-- // Lùi index xuống (giống pop)
		} else {
			// Chỗ này vốn là dấu -> Lấy dấu từ cuối kho specials đắp vào
			result[i] = specials[sIdx]
			sIdx-- // Lùi index xuống
		}
	}

	// Convert mảng byte về string để trả hàng
	return string(result)
}
