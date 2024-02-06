class Solution {
public:
    int minCostClimbingStairs(vector<int>& cost) {
        int n = cost.size();

        vector<int> memo(n+1, 0);
        memo[n-1] = cost[n-1];
        
        for(int i =n-2; i >= 0; i--) {
            memo[i] = cost[i] + min(memo[i+1], memo[i+2]);
        }

        return min(memo[0], memo[1]);
    }
};