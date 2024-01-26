/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    class Pair {
        public int level;
        public TreeNode node;
        public Pair(int level, TreeNode node) {
            this.level = level;
            this.node = node;
        }
    }
    
    public int maxLevelSum(TreeNode root) {
        Queue<Pair> bfs = new LinkedList<Pair>();
        
        int max = Integer.MIN_VALUE;
        int levelMax = -1;
        
        int currentSum = 0;
        int currentLevel = 1;
        
        bfs.add(new Pair(1, root));
        
        while(!bfs.isEmpty()) {
            Pair tmp = bfs.remove();
            
            if(currentLevel != tmp.level) {
                if(currentSum > max) {
                    max = currentSum;
                    levelMax =  currentLevel;
                }
                currentSum = 0;
                currentLevel = tmp.level;
            }
            
            currentSum += tmp.node.val;
            
            if(tmp.node.left != null) {
                bfs.add(new Pair(tmp.level + 1, tmp.node.left));
            }
            if(tmp.node.right != null) {
                bfs.add(new Pair(tmp.level + 1, tmp.node.right));
            }
        }
        
        if(currentSum > max) {
            levelMax = currentLevel;
            max = currentSum;
        }
        
        return levelMax;
    }
}