class Solution:
    def longestValidParentheses(self, s):
        """
        :type s: str
        :rtype: int
        """
        dp = [0 for i in range(len(s))]
        maximum = 0
        for i in range(len(s)):
            if ')' == s[i] and i - 1 >= 0:
                if '(' == s[i - 1]:
                    dp[i] = (dp[i - 2] if i - 2 >= 0 else 0) + 2
                # elif '(' == dp[i-1] > 0 and s[i - dp[i-1] - 1] == '(':
                #     dp[]
                elif i - dp[i - 1] > 0 and s[i - dp[i - 1] - 1] == '(':
                    # like ()(()) got count 01 as 2 and 34 as 2 than 26 as +2
                    dp[i] = dp[i - 1] + (dp[i - dp[i - 1] - 2] if (i - dp[i - 1]) >= 2 else 0) + 2
                maximum = max(dp[i], maximum)
        return maximum

    # todo more efficiency


if __name__ == '__main__':
    s = Solution()
    # print(s.longestValidParentheses(')('))
    # print(s.longestValidParentheses('(()'))
    # print(s.longestValidParentheses(')()())'))
    # print(s.longestValidParentheses('()(())'))
    # print(s.longestValidParentheses('()(())'))
    # print(s.longestValidParentheses('(()())'))
    print(s.longestValidParentheses('()(())'))
    # print(s.longestValidParentheses('()()()'))
    print('end')
