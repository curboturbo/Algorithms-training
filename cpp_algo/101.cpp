#include <iostream>
#include <queue>
/// проверка на симметричность дерева

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};

bool isSymmetric(TreeNode* root) {
    std::queue<std::pair<TreeNode*, TreeNode*>> q;
    q.push({root->left,root->right});
    while (!q.empty()){
        auto [t1,t2] = q.front();
        q.pop();
        if (!t1 && !t2) continue; 
        if (!t1 || !t2) return false;
        if (t1->val != t2->val) return false;
        q.push({t1->left,t2->right});
        q.push({t1->right,t2->left});
    }
    return true;
}
