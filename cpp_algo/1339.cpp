#include <iostream>
#include <vector>

#include <algorithm>

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
    int maxProduct(TreeNode* root){
        int sum = calculate(root);
        dfs(root,sum);
        return Max % MOD;
    }
private:
    int calculate(TreeNode* root){
        if (root==nullptr){return 0;}
        root->val += calculate(root->left) + calculate(root->right);
        return root->val;
    }
    void dfs(TreeNode* root,int total){
        if (root==nullptr){return;}
        long long subtree1 = root->val;
        long long subtree2 = (total - subtree1)*subtree1;
        Max = std::max(Max, subtree2);
        dfs(root->left, total);
        dfs(root->right,total);
    }


    long long Max = 0;
    static constexpr long long MOD = 1e9 + 7;
};
