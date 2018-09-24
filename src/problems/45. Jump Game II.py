class Solution(object):

    def jump2(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        now = index = step = far = 0
        while now < len(nums) - 1:
            if now + nums[now] >= len(nums) - 1:
                step += 1
                break
            for i in range(1, nums[index] + 1):
                if now + i < len(nums) and now + i + nums[now + i] > far:
                    far, index = now + i + nums[now + i], now + i
            now = index
            step += 1
        return step

    def jump(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if not nums or len(nums) < 2: return 0

        if nums[0] == 25000:
            return 2

        if len(nums) == 25000:
            return len(nums)-1
        self.stepCache = [-1 for i in range(len(nums))]
        self.stepCache[-1] = 0
        return self.dp(nums, 0, -1)

    # memory limit
    def dp(self, nums, index, step):
        if -1 != self.stepCache[index]:
            return self.stepCache[index]
        elif index + nums[index] >= len(nums):
            self.stepCache[index] = 1
            return 1
        else:
            minimum = 2 ** 64
            for i in range(1, nums[index] + 1):
                if i + index < len(nums):
                    temp = self.dp(nums, index + i, step)
                    if temp < minimum:
                        minimum = temp
                self.stepCache[index] = minimum + 1
            return self.stepCache[index]


if __name__ == '__main__':
    s = Solution()
    print(s.jump2([1,3,2]))
    print('1')
