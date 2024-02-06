class Solution {
public:
    vector<int> successfulPairs(vector<int>& spells, vector<int>& potions, long long success) {
        // n log n sort
        sort(potions.begin(), potions.end());
        vector<int> ans;

        for(auto spell:spells) {
            int min = success / spell;
            auto it = lower_bound(potions.begin(), potions.end(), min);
            if (it == potions.end()) {
                ans.push_back(0);
                continue;
            }
            
            long long mult;
            while(it != potions.end()) {
                mult = static_cast<long long>(*it) * static_cast<long long>(spell);
                if(mult >= static_cast<long long>(success)) break;
                it++;
            };
            ans.push_back(potions.end()-it);
        }

        return ans;
    }
};