class Solution {
public:
    int minReorder(int n, vector<vector<int>>& connections) {
        unordered_map<int, vector<pair<int, bool>>> adj;
        for(auto connection: connections) {
            if(adj.find(connection[0]) == adj.end()) {
                adj[connection[0]] = vector<pair<int, bool>>{make_pair(connection[1], true)};
            } else {
                adj[connection[0]].push_back(make_pair(connection[1], true));
            }

            if(adj.find(connection[1]) == adj.end()) {
                adj[connection[1]] = vector<pair<int, bool>>{make_pair(connection[0], false)};
            } else {
                adj[connection[1]].push_back(make_pair(connection[0], false));
            }
        }

        unordered_set<int> visited;
        stack<int> dfs;
        dfs.push(0);
        int counter = 0;

        while(!dfs.empty()) {
            int current = dfs.top();
            dfs.pop();

            cout << current << endl;
            visited.insert(current);

            for(auto node_pair: adj[current]) {
                if(visited.find(node_pair.first) != visited.end()) continue;
                if(node_pair.second == true) {
                    cout << current << " " << node_pair.first << endl;
                    counter++;
                }
                
                dfs.push(node_pair.first);
            }
        }

        return counter;
    }
};