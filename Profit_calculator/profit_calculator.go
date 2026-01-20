package main

import (
	"fmt"
	"os"
)

/*
*
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

-- RULE

	Show error mess and exit if invalid input is provided
	not negative nums
	not 0
	* Store result calculator in to file
*/
const resultFile = "result.txt"

func main() {

	fmt.Println("TIỀN SẠCH LÀ TIỀN KO CẦN RỬA")
	var thuNhap, chiPhi, tiLeThue float64

	inText("Thu nhập của bạn: ")
	nhapDL(&thuNhap)

	if thuNhap <= 0 {
		fmt.Println("Thu nhap ko hop le")
		return
	}
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
		writeToFile(thuNhapTrcThue, thuNhapSauThue, tiLeThuNhap)
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

func writeToFile(ebt, profit, ratio float64) {
	dataText := fmt.Sprintf("EBT (Trước thuế): %.2f\nProfit(Sau Thuế): %.2f\nRatio: %.2f\n", ebt, profit, ratio)
	file, err := os.OpenFile(resultFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		fmt.Println("Cant not write file: ", err)
		panic(err)
	}

	defer file.Close()

	if _, err := file.WriteString(dataText); err != nil {
		fmt.Println("Error when write file: ", err)
	} else {
		fmt.Println("✅ Đã lưu báo cáo vào file 'resutl.txt'!")
	}
}
