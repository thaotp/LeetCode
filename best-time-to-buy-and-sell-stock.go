# https://leetcode.com/problems/best-time-to-buy-and-sell-stock
# Key: Tìm khác số hiệu lớn nhất giữa 2 phần tử trong mảng theo thứ tự tăng dần.
func maxProfit(prices []int) int {
    if prices == nil {
        return 0
    }
    
    n := len(prices)
    if n <= 1 {
        return 0
    }
    
    max, min := 0, prices[0]
    for ind := range prices {
        if prices[ind] - min > max {
            max = prices[ind] - min
        }

        if min > prices[ind] {
            min = prices[ind]
        }
    }
    
    return max
}
