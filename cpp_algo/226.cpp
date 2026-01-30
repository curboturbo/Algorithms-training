#include <iostream>


struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};



TreeNode* invertTree(TreeNode* root) {
    if (root==nullptr){return root;}
    else{
        auto left_child = invertTree(root->left);
        auto right_child = invertTree(root->right);
        root->left = right_child;
        root->right = left_child;
        return root;
    }
}