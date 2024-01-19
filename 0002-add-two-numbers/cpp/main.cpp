/**
 * Definition for singly-linked list.
 * struct ListNode {
 *     int val;
 *     ListNode *next;
 *     ListNode() : val(0), next(nullptr) {}
 *     ListNode(int x) : val(x), next(nullptr) {}
 *     ListNode(int x, ListNode *next) : val(x), next(next) {}
 * };
 */
class Solution {
public:
    ListNode* addTwoNumbers(ListNode* l1, ListNode* l2, int res=0) {
        // ListNode *ret, *itl1, *itl2;
        int sum;
        if(l1 && l2) {
            sum = l1->val + l2->val + res;
        } else if(l1) {
            sum = l1->val + res;
        } else if(l2) {
            sum = l2->val + res;
        } else if(res){
            return new ListNode(res);
        } else {
            return NULL;
        }
        if(sum >= 10) {
            res = sum/10;
            sum = sum%10;
        } else {
            res = 0;
        }
        
        ListNode *next = addTwoNumbers(l1?l1->next:NULL, l2?l2->next:NULL, res);
        
        ListNode *ret = new ListNode(sum, next);
        
        return ret;
    }
};