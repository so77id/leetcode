class Solution {
public:
    /*
    
    */

    bool canVisitAllRooms(vector<vector<int>>& rooms) {
        stack<int> keys;
        set<int> visited = {0};

        for(auto key:rooms[0]) {
            keys.push(key);
        }

        while(keys.empty() == false) {
            int current_key = keys.top();
            keys.pop();
            if(visited.find(current_key) == visited.end()) {
                visited.insert(current_key);
                for(auto key:rooms[current_key]){
                    keys.push(key);
                }
            }
        }

        return visited.size() == rooms.size(); 
    }
};