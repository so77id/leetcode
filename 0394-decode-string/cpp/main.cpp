#include <string>

class Solution {
public:
    
    int i=0;
    
    string repeat(string s, int n) 
    { 
        // Copying given string to temparory string. 
        string s1 = s; 

        for (int i=1; i<n;i++) 
            s += s1; // Concatinating strings 

        return s; 
    }
    
    string decodeString(string s) {
        int num = 0;
        int par =0;
        string res = "";
        while(i < s.length()) {
            if(s[i] == '[' ) {
                i++;
                par++;
                string tmp = decodeString(s);
                res += repeat(tmp, num);
                num = 0;
            } else if(s[i] == ']') {
                if(par == 0) {
                    return res;
                } else par--, i++;
            } else if (s[i] >= 48 && s[i] <= 57) {
                // numeric
                num = 10 * num + s[i]-'0';
                i++;
            } else{
                // string
                res += s[i];
                i++;
            }
        }
        return res;
        
    }
};