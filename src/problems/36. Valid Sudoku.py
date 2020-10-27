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
    import time

    for i in range(100000):
        if False:
            pass
        a=1

    start_time = time.time()
    for i in range(10000):
        if False:
            pass
        a=1
    print("--- %s ms ---" % ((time.time() - start_time) * 1000))

    start_time = time.time()
    for i in range(10000):
        if False:
            pass
        else:
            a=1
    print("--- %s ms ---" % ((time.time() - start_time) * 1000))
