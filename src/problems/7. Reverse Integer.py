class Solution:
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """
        r = int(('-' if x < 0 else '') + str(abs(x))[::-1])
        return r if -2147483648 <= r <= 2147483647 else 0


if __name__ == '__main__':
    s = Solution()
    print(s.reverse(-321))
