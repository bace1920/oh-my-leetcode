# Definition for singly-linked list.
# class ListNode(object):
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution(object):
    def mergeKLists(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        head = now = ListNode(0)

        d = {}
        for l in lists:
            while l:
                if not d.get(l.val, None):
                    d[l.val] = [l]
                else:
                    d[l.val].append(l)
                l = l.next
        for key in sorted(d.iterkeys()):
            for l in d[key]:
                now.next = l
                now = now.next
        return head.next

    def mergeKListsV2(self, lists):
        """
        :type lists: List[ListNode]
        :rtype: ListNode
        """
        count = len(lists)
        if not count:
            return lists
        p = 1
        while p < count:
            for i in range(0, count - p, p * 2):
                lists[i] = self.merge2list(lists[i], lists[i + p])
            p *= 2
        return lists[0]

    def merge2list(self, l1, l2):
        head = now = ListNode(0)
        while l1 and l2:
            if l1.val < l2.val:
                now.next = l1
                l1 = l1.next
            else:
                now.next = l2
                l2 = l2.next
            now = now.next
        now.next = l1 if l1 else l2
        return head.next
