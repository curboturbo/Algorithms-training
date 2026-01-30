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



int maxDepth(TreeNode* root) {
    std::queue<TreeNode*> q = {};
    int answer = 1;
    q.push(root);
    while (!q.empty()) {
        int level_size = q.size();
        for (int i = 0; i < level_size; ++i) {
            TreeNode* node = q.front();
            q.pop();
            if (node->left) q.push(node->left);
            if (node->right) q.push(node->right);
        }
        answer++;
    }
    return answer;
}