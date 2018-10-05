class Solution:
    def isValidSudoku(self, board):
        """
        :type board: List[List[str]]
        :rtype: bool
        """
        tag = set()
        for i in range(9):
            for q in range(9):
                if '.' != board[i][q]:
                    if (i, board[i][q]) in tag or (board[i][q], q) in tag or (i // 3, q // 3, board[i][q]) in tag:
                        return False

                    tag.add((i, board[i][q]))
                    tag.add((board[i][q], q))
                    tag.add((i // 3, q // 3, board[i][q]))
        return True


if __name__ == '__main__':
    s = Solution()
    print(s.isValidSudoku([
        [".", "9", ".", ".", "4", ".", ".", ".", "."],
        ["1", ".", ".", ".", ".", ".", "6", ".", "."],
        [".", ".", "3", ".", ".", ".", ".", ".", "."],
        [".", ".", ".", ".", ".", ".", ".", ".", "."],
        [".", ".", ".", "7", ".", ".", ".", ".", "."],
        ["3", ".", ".", ".", "5", ".", ".", ".", "."],
        [".", ".", "7", ".", ".", "4", ".", ".", "."],
        [".", ".", ".", ".", ".", ".", ".", ".", "."],
        [".", ".", ".", ".", "7", ".", ".", ".", "."]
    ]))
