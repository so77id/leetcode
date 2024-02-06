class Solution {
public:
    int rob(vector<int>& nums) {
        if(nums.size() == 1) return nums[0];
        if(nums.size() == 2) return max(nums[0], nums[1]);

        vector<int> memo(nums.size(), 0);
        for(int i = 0; i < nums.size(); i++) {
            if(i-2 >= 0) {
                memo[i] = nums[i] + *max_element(memo.begin(), memo.begin()+i-1);
            } else {
                memo[i] = nums[i];
            }
            
        }

        return max(memo[nums.size()-1], memo[nums.size()-2]);
    }
};