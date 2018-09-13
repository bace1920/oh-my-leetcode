class Solution:
    def myAtoi(self, str):
        """
        :type str: str
        :rtype: int
        """
        start = False
        read = ''
        signed = 1
        for c in str:
            if not start and c == '+':
                start = True
                signed = 1
                read += c
            elif not start and c == '-':
                start = True
                signed = - 1
            elif ord('0') <= ord(c) <= ord('9'):
                start = True
                read += c
            elif start and not ord('0') <= ord(c) <= ord('9'):
                break
            elif not start and c != ' ':
                break

        if not read or read == '+' or read == '-':
            return 0

        read = int(read) * signed
        if read > 2 ** 31 - 1:
            return 2 ** 31 - 1
        elif read < -2 ** 31:
            return -2 ** 31
        else:
            return read

if __name__ == '__main__':
    s = Solution()
    print(s.myAtoi('ad 123'))
