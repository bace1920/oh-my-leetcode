# -*- coding: utf-8 -*-


class Solution:
    def convert(self, s, numRows):
        """
        we can treat the ZipZag of like follow matrix
        1   7    14
        2 4 8 13 15
        3 5 9 12
          6   11
        :type s: str
        :type numRows: int
        :rtype: str
        """
        length: int = len(s)
        if (numRows == 1) or length <= numRows:
            return s
        groupSize: int = numRows - 1
        groupNum: int = length // groupSize if length // groupSize == 0 else length // groupSize + 1
        result = ''
        i = 0
        while i < (groupSize + 1) * groupNum:
            y = i // groupNum
            x = i - y * groupNum
            if (x % 2 == 0 and y == numRows - 1) or (x % 2 == 1 and y == 0):
                # these coordinate on the matrix is empty
                pass
            else:
                index = x * groupSize + ((numRows - y - 1) if x % 2 else y)
                if index < length:
                    result += s[index]
            i += 1
        return result

    def convertV2(self, s, numRows):
        if numRows == 1:
            return s
        else:
            result = [''] * numRows
            currentRow = 0
            for ch in s:
                if currentRow == 0:
                    step = 1
                elif currentRow == numRows - 1:
                    step = -1
                result[currentRow] += ch
                currentRow += step
            return ''.join(result)


if __name__ == '__main__':
    s = Solution()
    print(s.convertV2('PAYPALISHIRING', 3))
