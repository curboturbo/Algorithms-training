#include <iostream>
#include <queue>

 struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

class Solution {
public:
    int kthSmallest(TreeNode* root, int k) {
        int travel_slut = 0;
        inorder_seacrh(root,k);
        return answer;
    }

private:
    void inorder_seacrh(TreeNode* root,int k){
        if (root==nullptr){return;}
        inorder_seacrh(root->left,k);
        cnt ++;
        if (cnt == k){
            answer = root->val;
            return;
        }
        std::cout<<root->val<<std::endl;
        inorder_seacrh(root->right,k);
    }
    int answer;
    int cnt = 0;
};
// 1 000 0 000 1
