func removeDuplicates(nums []int) int {
    if len(nums)==0{return 0}

    for i:=len(nums)-2; i>=0; i--{
        if nums[i]==nums[i+1]{
            copy (nums[i:], nums[i+1:])
            nums=nums[:len(nums)-1]

        }
   }

    return len(nums)
}
