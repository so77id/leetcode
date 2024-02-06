class Solution {
public:
    vector<string> letterCombinations(string digits) {
        unordered_map<char, vector<char>> kb = {
            {'2', {'a', 'b', 'c'}},
            {'3', {'d', 'e', 'f'}},
            {'4', {'g', 'h', 'i'}},
            {'5', {'j', 'k', 'l'}},
            {'6', {'m', 'n', 'o'}},
            {'7', {'p', 'q', 'r', 's'}},
            {'8', {'t', 'u', 'v'}},
            {'9', {'w', 'x', 'y', 'z'}}
        };

        vector<string> ans;
        
        if(digits == "") {
            return ans;
        }

        for(auto letter:kb[digits[0]]) {
            ans.push_back(string(1, letter));
        }

        for(int i = 1; i < digits.size(); i++) {
            vector<string> n_ans;
            for(auto a:ans) {
                for(auto letter:kb[digits[i]]) {
                    n_ans.push_back(a+letter);
                }
            }
            ans = n_ans;
        }

        return ans;
    }
};