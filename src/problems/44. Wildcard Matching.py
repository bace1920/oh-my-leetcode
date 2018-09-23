import re
"""
TODO
"""

class Solution(object):
    def isMatch(self, s, p):
        """
        :type s: str
        :type p: str
        :rtype: bool
        """
        pattern = self.parse(p)
        mode, index = self.nextPattern(pattern, -1)  # 0 match as ch, 1 match *, 2 match ?
        for i in range(len(s)):
            if index >= len(pattern):
                return False
            if 0 == mode:
                if s[i:i + len(pattern[index])] == pattern[index]:
                    i += len(pattern[index]) - 1
                    index += 1
                    mode, index = self.nextPattern(pattern, -1)  # 0 match as ch, 1 match *, 2 match ?
                else:
                    i += 1
            if 1 == mode:
                index += 1
                mode, index = self.nextPattern(pattern, -1)  # 0 match as ch, 1 match *, 2 match ?

            if 2 == mode:
                index += 1
                i += 1
                mode, index = self.nextPattern(pattern, -1)  # 0 match as ch, 1 match *, 2 match ?
        return True if index >= len(pattern) else False

    def nextPattern(self, l, index):
        index += 1
        mode = -1
        if (len(l) > index):
            if l[index] == '*':
                mode = 1
            elif l[index] == '?':
                mode = 2
            else:
                mode = 0
        return mode, index

    def parse(self, s):
        result = []
        current = ''
        for ch in s:
            if ch in ['*', '?']:
                if '' != current:
                    result.append(current)
                    current = ''
                result.append(ch)
            else:
                current += ch
        if '' != current:
            result.append(current)
        return result


if __name__ == '__main__':
    s = Solution()
    print(s.isMatch('aa', '*'))
    # print(s.parse('123*456?asd'))
