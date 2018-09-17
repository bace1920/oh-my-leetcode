class Solution:
    def longestPalindrome(self, s):
        """
        :type s: str
        :rtype: str
        """
        length = len(s)
        start = end = 0
        for i in range(0, length - 1):
            s1, e1 = self.tryGetPalindrome(s, i, i)
            s2, e2 = self.tryGetPalindrome(s, i, i + 1)
            longer = max(e1 - s1 + 1, e2 - s2 + 1)
            if longer > end - start:
                (start, end) = (s1, e1) if longer % 2 else (s2, e2)

        return s[int(start): int(end) + 1]

    def tryGetPalindrome(self, s, smid1, smid2):
        length = len(s)
        mid1, mid2 = smid1, smid2
        while mid1 >= 0 and mid2 < length and s[mid1] == s[mid2]:
            mid1 -= 1
            mid2 += 1
        return mid1 + 1, mid2 - 1


if __name__ == '__main__':
    s = Solution()
    print(s.longestPalindrome('bb'))
