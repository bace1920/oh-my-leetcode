class Solution:
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        length = len(s)
        re = s[::-1]
        for i in range (0,length):
            

if __name__ == '__main__':
    s = Solution()
    print(s.longestPalindrome('cbbd'));