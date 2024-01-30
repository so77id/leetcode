class Solution {
	public:
		int findCircleNum(vector<vector<int>>& isConnected) {
			unordered_set<int> visited;
			int counter = 0;
			stack<int> dfs;
	
			for(int i=0; i < isConnected.size(); i++){
				if(visited.find(i) != visited.end()) {
					continue;
				}
				counter++;
	
				visited.insert(i);
				dfs.push(i);
	
				while(!dfs.empty()) {
					int actual = dfs.top();
					dfs.pop();
					
					visited.insert(actual);
					for(int j=0; j<isConnected[actual].size(); j++) {
						if(isConnected[actual][j] == 0) continue;
						if(visited.find(j) != visited.end()) continue;
						dfs.push(j);
					}
				}
			}
	
			return counter;
		}
	};