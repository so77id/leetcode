class Solution {
public:
    long long totalCost(vector<int>& costs, int k, int candidates) {
        auto compare = [](const pair<int, int>& a, const pair<int, int>& b) {
            if (a.first == b.first) return a.second > b.second; // Desempate por el segundo elemento
            return a.first > b.first; // Ordena primero por el primer elemento
        };
        priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(compare)> pq(compare);

        // set<int> index_not_included;
        int next_head, next_tail;
        
        if(candidates * 2 >= costs.size()) {
            for(int i = 0; i < costs.size(); i++) {
                pq.push(make_pair(costs[i], i));
                next_head = costs.size();
                next_tail = -1;
            }
        } else {
            int i;
            for(i = 0; i < candidates; i++) {
                pq.push(make_pair(costs[i], i));
                pq.push(make_pair(costs[costs.size()-1-i], costs.size()-1-i));
            }
            next_head = i;
            next_tail = costs.size()-1-i;

            // for(int j = i; j < costs.size()-i; j++) {
            //     index_not_included.insert(j);
            // }
        }

        long long ans = 0;
        for(int i = 0; i < k; i++) {
            auto [cost, index] = pq.top();
            pq.pop();
            ans += cost;
            if(next_tail >= next_head) {
                int next;
                if(index > next_head) {
                    next = next_tail--;
                } else {
                    next = next_head++;
                }
                pq.push(make_pair(costs[next], next));
            }
        }

        return ans;
    }
};