class Solution {
public:
    unordered_map<int, vector<string>> keyboard = { {2, {"a", "b", "c"}}, {3, {"d", "e", "f"}}, {4, {"g", "h", "i"}}, {5, {"j", "k", "l"}}, {6, {"m", "n", "o"}}, {7, {"p", "q", "r", "s"}}, {8, {"t", "u", "v"}}, {9, {"w", "x", "y", "z"}} };
    
    vector<string> letterCombinations(string digits, int ini=0) {
        if(digits.length() == 0) return {};
        if(ini+1 == digits.length()) {
            return keyboard[digits[ini]-'0'];
        }
        vector<string> next = letterCombinations(digits, ini+1);
        
        vector<string> sol;
        for(const auto &c:keyboard[digits[ini] - '0']){
            for(const auto &n:next) {    
                sol.push_back(c+n);      
            }
        }
        
        return sol;
    }
};