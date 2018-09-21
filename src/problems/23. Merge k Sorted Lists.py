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
                l=l.next
        for key in sorted(d.iterkeys()):
            for l in d[key]:
                now.next = l
                now= now.next
        return head.next
        
                

        