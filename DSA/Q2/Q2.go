package main 

func minimumK(nums []int) int {
    maxVal := 0
    for _, v := range nums {
        if v > maxVal {
            maxVal = v
        }
    }

    n := len(nums)
    // FIX: Biên phải phải bao quát cả trường hợp n > maxVal^2
    // Nếu maxVal quá nhỏ, đáp án có thể rơi vào khoảng sqrt(n)
    right := maxVal
    if n > right {
        right = n
    }

    left := 1
    ans := right 

    for left <= right {
        mid := left + (right-left)/2
        
        // --- Check ---
        ops := 0
        limit := mid * mid
        possible := true
        
        for _, x := range nums {
            // Phép chia trần: (x + k - 1) / k
            ops += (x + mid - 1) / mid
            
            // Optimization: Cắt sớm nếu lố
            if ops > limit {
                possible = false
                break
            }
        }
        
        if possible {
            ans = mid
            right = mid - 1
        } else {
            left = mid + 1
        }
    }
    return ans
}