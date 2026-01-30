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

int maxLevelSum(TreeNode* root) {
    std::queue<TreeNode*> q;
    q.push(root);
    int global_level = 1;
    int max_level = 1;
    int max_sum = root->val;
    while (!q.empty()) {
        int nodes = q.size();
        int sum = 0;
        for (int i = 0; i < nodes; i++) {
            TreeNode* a = q.front();
            q.pop();
            sum += a->val;
            if (a->left) q.push(a->left);
            if (a->right) q.push(a->right);
        }
        if (sum > max_sum) {
            max_sum = sum;
            max_level = global_level;
        }
        global_level++;
    }
    return max_level;
}


int main() {
    TreeNode* root = new TreeNode(1);
    root->left = new TreeNode(7);
    root->right = new TreeNode(0);
    root->left->left = new TreeNode(7);
    root->left->right = new TreeNode(-8);
    //Solution sol;
    //int result = sol.maxLevelSum(root);

   // std::cout << "Level with maximum sum: " << result << std::endl;

    return 0;
}