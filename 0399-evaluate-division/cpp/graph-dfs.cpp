#include <vector>
#include <string>
#include <unordered_map>
#include <unordered_set>
using namespace std;

class Solution {
public:
    vector<double> calcEquation(vector<vector<string>>& equations, vector<double>& values, vector<vector<string>>& queries) {
        // Construir el grafo
        unordered_map<string, unordered_map<string, double>> G;
        for (int i = 0; i < equations.size(); ++i) {
            string u = equations[i][0];
            string v = equations[i][1];
            double val = values[i];
            G[u][v] = val;
            G[v][u] = 1.0 / val;
        }

        vector<double> ans;
        for (auto& q : queries) {
            unordered_set<string> visited;
            double result = dfs(G, visited, q[0], q[1]);
            ans.push_back(result);
        }
        return ans;
    }

private:
    double dfs(unordered_map<string, unordered_map<string, double>>& G, 
               unordered_set<string>& visited, 
               const string& start, 
               const string& end) {
        // Si no se encuentra el nodo de inicio o el nodo final, retorna -1.0
        if (G.find(start) == G.end() || G.find(end) == G.end()) return -1.0;

        // Si se encuentra un camino directo, retorna el valor.
        if (G[start].find(end) != G[start].end()) return G[start][end];

        visited.insert(start);

        for (auto& node : G[start]) {
            if (visited.find(node.first) == visited.end()) {
                double temp = dfs(G, visited, node.first, end);
                if (temp != -1.0) {
                    return node.second * temp;
                }
            }
        }

        return -1.0;
    }
};
