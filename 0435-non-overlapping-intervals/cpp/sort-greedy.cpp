class Solution {
public:
    int eraseOverlapIntervals(vector<vector<int>>& intervals) {
        std::sort(intervals.begin(), intervals.end(), 
            [](const std::vector<int>& a, const std::vector<int>& b) {
                return a[0] != b[0] ? a[0] < b[0] : a[1] < b[1];
            }
        );

        int ans = 0;
        int upper_limit = numeric_limits<int>::min();

        for(auto interval: intervals) {
            if (interval[0] >= upper_limit) {
                upper_limit = interval[1];
                continue;
            }
            if (interval[1] < upper_limit) {
                upper_limit = interval[1];
            }
            ans++;
        }

        return ans;
    }   
};