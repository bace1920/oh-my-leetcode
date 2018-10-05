class Solution:
    def maxArea(self, height):
        """
        :type height: List[int]
        :rtype: int
        """
        index = [0 for h in height]
        maximum = height[-1]
        for i in range(len(height) - 2, -1, -1):
            if height[i] > maximum:
                maximum = height[i]
            index[i] = maximum
        return index

if __name__ == '__main__':
    s = Solution()
    res = s.maxArea([1, 8, 6, 2, 5, 4, 8, 3, 7])
    # print(dir(res))
    print(res)
