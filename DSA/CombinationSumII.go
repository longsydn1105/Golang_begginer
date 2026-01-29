package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	var results [][]int

	// BƯỚC 1: Sort trước để dễ dàng xử lý trùng lặp và cắt nhánh
	sort.Ints(candidates)

	// Hàm đệ quy (closure) để tiện truy cập biến
	var backtrack func(start int, currentTarget int, path []int)
	backtrack = func(start int, currentTarget int, path []int) {
		// Base Case 1: Tìm thấy đích!
		if currentTarget == 0 {
			// Phải copy path ra slice mới, vì slice trong Go là tham chiếu
			// Nếu không copy, các bước sau sẽ sửa đè lên kết quả cũ
			temp := make([]int, len(path))
			copy(temp, path)
			results = append(results, temp)
			return
		}

		// Duyệt qua các ứng viên bắt đầu từ vị trí 'start'
		for i := start; i < len(candidates); i++ {
			// BƯỚC 2: Cắt nhánh thông minh (Pruning)
			// Vì đã sort, nếu số hiện tại > target thì mấy số sau chắc chắn cũng tạch -> Break luôn
			if candidates[i] > currentTarget {
				break
			}

			// BƯỚC 3: Xử lý trùng lặp (Logic quan trọng nhất bài)
			// Nếu số này giống số liền trước VÀ đây không phải là vòng lặp đầu tiên của đệ quy hiện tại
			// (i > start nghĩa là ta đã xét thằng candidates[i-1] ở vị trí này rồi, giờ xét tiếp thằng giống hệt thì sẽ ra kết quả y chang)
			if i > start && candidates[i] == candidates[i-1] {
				continue
			}

			// Thử chọn candidates[i]
			path = append(path, candidates[i])

			// Đệ quy tiếp với target nhỏ hơn.
			// Lưu ý: truyền 'i+1' vì đề bài cấm dùng lại chính phần tử đó (each number used once)
			backtrack(i+1, currentTarget-candidates[i], path)

			// Backtrack: Trả lại trạng thái cũ để thử số tiếp theo (bỏ số vừa thêm vào ra)
			path = path[:len(path)-1]
		}
	}

	backtrack(0, target, []int{})
	return results
}

func main1() {
	candidates := []int{10, 1, 2, 7, 6, 1, 5}
	target := 8
	fmt.Println(combinationSum2(candidates, target))
}
