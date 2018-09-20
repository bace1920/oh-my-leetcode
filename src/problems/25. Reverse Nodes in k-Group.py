# Definition for singly-linked list.
class ListNode(object):
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution(object):
    def reverseKGroupV2(self, head, k):
        before = last_last = ListNode(0)
        before.next = left = right = head
        while True:
            count = 0
            while right is not None and count < k:
                count += 1
                right = right.next
            if count == k:
                prev, now = left, right
                for _ in range(k):
                    nex = prev.next
                    prev.next = now
                    now = prev
                    prev = nex
                last_last.next = now
                last_last = left
                left = right
            else:
                return before.next

    def reverseKGroup(self, head, k):
        """
        :type head: ListNode
        :type k: int
        :rtype: ListNode
        """
        if not head or k == 1:
            return head
        new_head = head
        if self.isEnough(new_head, k):
            now_head, now_last, next_head = self.reverseK(head, k)
            new_head = now_head
            now_last.next = next_head
            last_last = now_last
            while self.isEnough(last_last.next, k):
                now_head, now_last, next_head = self.reverseK(next_head, k)
                last_last.next = now_head
                now_last.next = next_head
                last_last = now_last
        return new_head

    def isEnough(self, now_head, k):
        tmp = now_head
        for i in range(k):
            if not tmp:
                return False
            tmp = tmp.next
        return True

    def reverseK(self, head, k):
        now = head.next
        prev = head
        nex = now.next
        count = 0
        while count < k - 1:
            now.next = prev
            prev = now
            now = nex
            nex = now.next if now else None
            count += 1
        return prev, head, now


if __name__ == '__main__':
    now = head = ListNode(1)
    for i in range(2, 6):
        now.next = ListNode(i)
        now = now.next

    solu = Solution()
    head = solu.reverseKGroupV2(head, 2)
    while head is not None:
        print(head.val)
        head = head.next
