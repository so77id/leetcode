class Solution {
public:
    int minEatingSpeed(vector<int>& piles, int h) {
        int max = *max_element(piles.begin(), piles.end()); // O(n)
        int min = 1;
        
        while(min < max) {
            int mid = ((max - min) / 2) + min;
            int num = 0;
            for(auto pile:piles) {
                num += ceil(pile*1.0/mid);
            }
            
            if(num <= h) {
                max = mid;
            } else {
                min = mid+1;
            }
        }

        return max;
    }
};