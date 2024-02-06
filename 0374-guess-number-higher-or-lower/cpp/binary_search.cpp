/** 
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * int guess(int num);
 */

class Solution {
public:
    int guessNumber(int n) {
        int lower = 0;
        int high = n;
        int mid;

        while(lower <= high) {
            mid = ((lower - high)/ 2) + high;
            int ret = guess(mid);
            if(ret == 0) return mid;
            if(ret == 1) {
                lower = mid + 1;
            } else {
                high = mid - 1;
            }
        }
        return -1;
    }
};