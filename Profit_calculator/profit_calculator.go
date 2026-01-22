package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Constants nên đặt ở đầu, rõ ràng
const ResultFileName = "financial_report.txt"

// FinancialReport: Dùng Struct để gom nhóm dữ liệu (Data Container) [cite: 5, 753]
// Giúp truyền dữ liệu giữa các hàm dễ dàng hơn là truyền 3-4 tham số rời rạc.
type FinancialReport struct {
	Revenue     float64 // Doanh thu (Thu nhập)
	Expenses    float64 // Chi phí
	TaxRate     float64 // Tỉ lệ thuế (%)
	EBT         float64 // Earnings Before Tax (Thu nhập trước thuế)
	NetIncome   float64 // Thu nhập sau thuế
	ProfitRatio float64 // Tỉ lệ lợi nhuận
}

func main() {
	fmt.Println("💰 TIỀN SẠCH LÀ TIỀN KO CẦN RỬA 💰")

	// 1. INPUT PHASE
	revenue, err := getValidInput("Nhập doanh thu (Revenue): ", 0)
	if err != nil {
		fmt.Printf("❌ Lỗi: %v\n", err)
		return
	}

	expenses, err := getValidInput("Nhập chi phí (Expenses): ", 0)
	if err != nil {
		fmt.Printf("❌ Lỗi: %v\n", err)
		return
	}

	// Thuế có thể là 0%, nhưng không được âm.
	taxRate, err := getValidInput("Nhập tỉ lệ thuế (Tax Rate %): ", 0)
	if err != nil {
		fmt.Printf("❌ Lỗi: %v\n", err)
		return
	}

	// 2. PROCESSING PHASE
	// Tạo object report và tính toán
	report := FinancialReport{
		Revenue:  revenue,
		Expenses: expenses,
		TaxRate:  taxRate,
	}

	// Calculate logic tách biệt
	report.calculateMetrics()

	// 3. OUTPUT PHASE
	printReport(report)

	// 4. STORAGE PHASE
	if err := saveReportToFile(report); err != nil {
		fmt.Printf("⚠️ Không thể lưu file: %v\n", err)
	} else {
		fmt.Printf("✅ Đã lưu báo cáo vào file '%s'!\n", ResultFileName)
	}
}

// getValidInput: Hàm nhập liệu chuẩn chỉ.
// - minVal: Giá trị nhỏ nhất chấp nhận được (Validation)
// - Trả về error thay vì panic hay exit (Go Idiomatic) [cite: 663, 688]
func getValidInput(prompt string, minVal float64) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		inputStr, err := reader.ReadString('\n')
		if err != nil {
			return 0, fmt.Errorf("lỗi đọc dữ liệu")
		}

		inputStr = strings.TrimSpace(inputStr)
		value, err := strconv.ParseFloat(inputStr, 64)

		if err != nil {
			fmt.Println("⚠️  Vui lòng nhập một con số hợp lệ!")
			continue // Cho nhập lại
		}

		if value <= minVal {
			return 0, fmt.Errorf("giá trị phải lớn hơn %.0f", minVal)
		}

		return value, nil
	}
}

// calculateMetrics: Logic tính toán nằm gọn trong method của struct.
// High Cohesion: Dữ liệu và hành vi xử lý dữ liệu đó nằm cùng nhau. [cite: 877]
func (f *FinancialReport) calculateMetrics() {
	f.EBT = f.Revenue - f.Expenses
	f.NetIncome = f.EBT * (1 - f.TaxRate/100)

	// Xử lý chia cho 0
	if f.EBT == 0 {
		f.ProfitRatio = 0
	} else {
		f.ProfitRatio = (f.NetIncome / f.EBT) * 100
	}
}

// printReport: Chỉ làm nhiệm vụ hiển thị (Presentation Layer)
func printReport(f FinancialReport) {
	fmt.Println("\n--- 📊 BÁO CÁO TÀI CHÍNH ---")
	fmt.Printf("Thu nhập trước thuế (EBT) : %15.2f\n", f.EBT)
	fmt.Printf("Thu nhập sau thuế (Net)   : %15.2f\n", f.NetIncome)

	// Logic hiển thị đặc biệt nên nằm ở tầng hiển thị
	if f.EBT == 0 {
		fmt.Println("Tỉ lệ thu nhập            : Không xác định (EBT = 0)")
	} else {
		fmt.Printf("Tỉ lệ thu nhập            : %14.2f%%\n", f.ProfitRatio)
	}
	fmt.Println("-----------------------------")
}

// saveReportToFile: Chỉ làm nhiệm vụ IO (Ghi file)
func saveReportToFile(f FinancialReport) error {
	dataText := fmt.Sprintf(
		"Timestamp: %s\nRevenue: %.2f | Expenses: %.2f | Tax: %.2f%%\nEBT: %.2f\nNet Income: %.2f\nRatio: %.2f%%\n-------------------\n",
		"NOW", // Thực tế nên dùng time.Now().Format(...)
		f.Revenue, f.Expenses, f.TaxRate, f.EBT, f.NetIncome, f.ProfitRatio,
	)

	// Dùng os.OpenFile cờ chuẩn để Append
	file, err := os.OpenFile(ResultFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close() // Đảm bảo file luôn được đóng [cite: 494]

	if _, err := file.WriteString(dataText); err != nil {
		return err
	}
	return nil
}
