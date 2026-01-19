package main

import (
	"fmt"
)

/**
--Input
	Doanh thu
	Chi phi
	Ti Le thue
-- Tinh va in ra
	Thu nhap truoc thue
	Thu nhap sau thue
	Ti le thu nhap

*/
func main() {

	fmt.Println("TIỀN SẠCH LÀ TIỀN KO CẦN RỬA")
	var thuNhap, chiPhi, tiLeThue float64
	
	fmt.Print("Thu nhập của bạn: ")
	fmt.Scan(&thuNhap)

	fmt.Print("Chi phí của bạn: ")
	fmt.Scan(&chiPhi)

	fmt.Print("Tỉ lệ thuế : ")
	fmt.Scan(&tiLeThue)

	thuNhapTrcThue := thuNhap - chiPhi
	thuNhapSauThue := thuNhapTrcThue * (1 - tiLeThue/100)

	// In kết quả đẹp (Dùng Printf)
	fmt.Println("\n--- 📊 KẾT QUẢ ---")
	// %.2f nghĩa là: Số thực, lấy 2 số sau dấu phẩy (VD: 1000.00)
	fmt.Printf("Thu nhập trước thuế : %10.2f\n", thuNhapTrcThue)
	fmt.Printf("Thu nhập sau thuế   : %10.2f\n", thuNhapSauThue)

	// Logic check chia cho 0
	if thuNhap == 0 {
		fmt.Println("Tỉ lệ thu nhập      : Không xác định (Doanh thu = 0)")
	} else {
		tiLeThuNhap := thuNhapSauThue / thuNhap
		fmt.Printf("Tỉ lệ thu nhập      : %10.2f\n", tiLeThuNhap)
	}

}