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

-- TASK
	Viết hàm in ra thay cho Print
	Viết hàm nhập
	Viết hàm thực hiện cho 3 phép tính
*/
func main() {

	fmt.Println("TIỀN SẠCH LÀ TIỀN KO CẦN RỬA")
	var thuNhap, chiPhi, tiLeThue float64
	
	inText("Thu nhập của bạn: ")
	nhapDL(&thuNhap)

	inText("Chi phí của bạn: ")
	nhapDL(&chiPhi)

	inText("Tỉ lệ thuế : ")
	nhapDL(&tiLeThue)

	thuNhapTrcThue := thuNhapTT(thuNhap, chiPhi)
	thuNhapSauThue := thuNhapST(thuNhapTrcThue, tiLeThue)

	// In kết quả đẹp (Dùng Printf)
	fmt.Println("\n--- 📊 KẾT QUẢ ---")
	// %.2f nghĩa là: Số thực, lấy 2 số sau dấu phẩy (VD: 1000.00)
	fmt.Printf("Thu nhập trước thuế : %10.2f\n", thuNhapTrcThue)
	fmt.Printf("Thu nhập sau thuế   : %10.2f\n", thuNhapSauThue)

	// Logic check chia cho 0
	if thuNhap == 0 {
		fmt.Println("Tỉ lệ thu nhập      : Không xác định (Doanh thu = 0)")
	} else {
		tiLeThuNhap := tiLeThuNhap(thuNhapTrcThue, thuNhapSauThue)
		fmt.Printf("Tỉ lệ thu nhập      : %10.2f\n", tiLeThuNhap)
	}

}

func inText(text string) {
	fmt.Print(text)
}

func nhapDL(dl *float64) {
	fmt.Scan(dl)
}

func thuNhapTT(thuNhap, chiPhi float64) float64 {
	loiNhuan := thuNhap - chiPhi
	return loiNhuan
}

func thuNhapST(loiNhuan, tiLeThue float64) float64 {
	thuNhapST := loiNhuan * (1 - tiLeThue/100)
	return thuNhapST
}

func tiLeThuNhap(TNTT, TNST float64) (tiLe float64) {
	tiLe = (TNST / TNTT) * 100
	return
}