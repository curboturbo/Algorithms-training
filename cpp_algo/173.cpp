#include <iostream>
#include <stack>
#include <memory>

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right) : val(x), left(left), right(right) {}
};


class BSTIterator {
public:
    BSTIterator(TreeNode* root): root(root){
        left_bottom(root);
    }
    
    int next() {
        auto head = stack.top();
        stack.pop();
        left_bottom(head->right);
        return head->val;
    }
    
    bool hasNext() {
        return !stack.empty();
    }

private:
    std::stack<TreeNode*> stack;
    TreeNode* root;
    void left_bottom(TreeNode* root){
        while (root!=nullptr){
            stack.push(root);
            root=root->left;
        }
    }

};


