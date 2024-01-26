class Solution {
public:
    int maxNumberOfBalloons(string text) {
        vector<int> counter = {0, 0, 0, 0, 0};
        vector<int> divisors = {1, 1, 2, 2, 1};

        for(auto c:text) {
            if (c == 'b') counter[0]++;
            if (c == 'a') counter[1]++;
            if (c == 'l') counter[2]++;
            if (c == 'o') counter[3]++;
            if (c == 'n') counter[4]++;
        }

        int min = numeric_limits<int>::max();
        for(size_t i = 0; i < counter.size(); i++) {
            if (counter[i] <= 0) {
                return 0;
            }
            if (min > counter[i]/divisors[i]) 
                min = counter[i]/divisors[i];
        }
        return min;
    }
};