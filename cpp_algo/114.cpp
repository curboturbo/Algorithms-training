#include <vector>
#include <iostream>
#include <unordered_map>



struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

std::vector<int> list = {};
void flatten(TreeNode* root) {
    rebuild(root);
}

TreeNode* rebuild(TreeNode* root){
    if (root == nullptr){return nullptr;}
    auto left_side = rebuild(root->left);
    auto right_side = rebuild(root->right);
    if (left_side){
        left_side->right = root->right;
        root->right = root->left;
        root->left = nullptr;
    }
    if (right_side){return right_side;}
    if (left_side){return left_side;}
    return root;
}