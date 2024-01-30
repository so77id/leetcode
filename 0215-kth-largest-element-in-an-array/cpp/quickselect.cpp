/*
    Complexity Analysis

Given nnn as the length of nums,

Time complexity: O(n)O(n)O(n) on average, O(n2)O(n^2)O(n 
2
 ) in the worst case

Each call we make to quickSelect will cost O(n)O(n)O(n) since we need to iterate over nums to create left, mid, and right. The number of times we call quickSelect is dependent on how the pivots are chosen. The worst pivots to choose are the extreme (greatest/smallest) ones because they reduce our search space by the least amount. Because we are randomly generating pivots, we may end up calling quickSelect O(n)O(n)O(n) times, leading to a time complexity of O(n2)O(n^2)O(n 
2
 ).

However, the algorithm mathematically almost surely has a linear runtime. For any decent size of nums, the probability of the pivots being chosen in a way that we need to call quickSelect O(n)O(n)O(n) times is so low that we can ignore it.

On average, the size of nums will decrease by a factor of ~2 on each call. You may think: that means we call quickSelect O(log⁡n)O(\log n)O(logn) times, wouldn't that give us a time complexity of O(n⋅log⁡n)O(n \cdot \log n)O(n⋅logn)? Well, each successive call to quickSelect would also be on a nums that is a factor of ~2 smaller. This recurrence can be analyzed using the master theorem with a = 1, b = 2, k = 1:

T(n)=T(n2)+O(n)=O(n)\Large{T(n) = T(\frac{n}{2}) + O(n)} = O(n)T(n)=T( 
2
n
​
 )+O(n)=O(n)

Space complexity: O(n)O(n)O(n)

We need O(n)O(n)O(n) space to create left, mid, and right. Other implementations of Quickselect can avoid creating these three in memory, but in the worst-case scenario, those implementations would still require O(n)O(n)O(n) space for the recursion call stack.
*/

class Solution {
public:
    int findKthLargest(vector<int>& nums, int k) {
        return quickSelect(nums, k);
    }
    
    int quickSelect(vector<int>& nums, int k) {
        int pivot = nums[rand() % nums.size()];
        
        vector<int> left;
        vector<int> mid;
        vector<int> right;
        
        for (int num: nums) {
            if (num > pivot) {
                left.push_back(num);
            } else if (num < pivot) {
                right.push_back(num);
            } else {
                mid.push_back(num);
            }
        }
        
        if (k <= left.size()) {
            return quickSelect(left, k);
        }
        
        if (left.size() + mid.size() < k) {
            return quickSelect(right, k - left.size() - mid.size());
        }
        
        return pivot;
        
    }
};