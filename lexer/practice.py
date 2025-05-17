class Solution:
    def threeSum(self, nums):
        myarr = []
        if len(nums) <= 2:
            return myarr

        for i in range(len(nums) - 2):
            for j in range(i + 1, len(nums) - 1):
                for k in range(j + 1, len(nums)):
                    if nums[i] + nums[j] + nums[k] == 0:
                        temp = [nums[i], nums[j], nums[k]]
                        temp.sort()
                        if not temp in myarr:
                            myarr.append(temp)
        return myarr


sbj = Solution()

print(sbj.threeSum([-1, 0, 1, 2, -1, -4]))
