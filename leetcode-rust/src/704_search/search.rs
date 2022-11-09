impl Solution {
    pub fn search1(nums: Vec<i32>, target: i32) -> i32 {
        let mut left: usize = 0;
        let mut right: usize = nums.len() - 1;
        while left as i32 <= right as i32 {
            let mid: usize = ((right - left) >> 1) + left;
            if nums[mid] == target {
                return mid as i32;
            }
            if nums[mid] < target {
                left = mid + 1 as usize;
            } else {
                right = mid - 1 as usize;
            }
        }
        return -1;
    }

    pub fn search(nums: Vec<i32>, target: i32) -> i32 {
        let mut i: usize = 0;
        let mut j: usize = nums.len() - 1;
        while i as i32 <= j as i32 {
            if nums[i] == target {
                return i as i32;
            }
            if nums[j] == target {
                return j as i32;
            }
            i += 1;
            j -= 1;
        }
        return -1
    }
}
