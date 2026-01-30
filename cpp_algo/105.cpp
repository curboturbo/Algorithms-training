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
int preIndex = 0;

TreeNode* build(std::vector<int>& preorder, int left, int right);

TreeNode* buildTree(std::vector<int>& preorder, std::vector<int>& inorder) {
    for (int i = 0; i < inorder.size(); i++) {
        pos[inorder[i]] = i;
    }
    return build(preorder, 0, inorder.size() - 1);
}

TreeNode* build(std::vector<int>& preorder, int left, int right) {
    if (left > right) return nullptr;

    int rootVal = preorder[preIndex];
    int mid_position = pos[rootVal];
    TreeNode* root = new TreeNode(rootVal);

    preIndex++;
    root->left = build(preorder, left, mid_position - 1);
    root->right = build(preorder, mid_position + 1, right);
    return root;
}

