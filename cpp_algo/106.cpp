#include <iostream>
#include <vector>
#include <unordered_map>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};


std::unordered_map<int, int> pos;
int postIndex = 0;

TreeNode* buildTree(std::vector<int>& inorder, std::vector<int>& postorder){
    postIndex = inorder.size()-1;
    for (int i = 0; i < inorder.size(); i++) {
        pos[inorder[i]] = i;
    }
    return build(postorder, 0, inorder.size() - 1);
}

TreeNode* build(std::vector<int>& postorder, int left, int right) {
    if (left>right){return nullptr;}

    int rootVal = postorder[postIndex--];
    TreeNode* node = new TreeNode(rootVal);
    int inorderIndex = pos[rootVal];
    node->right = build(postorder,inorderIndex+1,right);
    node->left = build(postorder,left,inorderIndex-1);
    return node;
}