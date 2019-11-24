package array2

func matchUp(nums1 []int, nums2 []int) int {
	n := 0
	for i := 0; i < len(nums1); i++ {
		if nums1[i] == nums2[i] {
			continue
		}
		if nums1[i] > nums2[i] {
			if nums1[i]-nums2[i] <= 2 {
				n++
			}
		} else {
			if nums2[i]-nums1[i] <= 2 {
				n++
			}
		}
	}
	return n
}
