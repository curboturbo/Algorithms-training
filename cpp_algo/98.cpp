
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
    bool isValidBST(TreeNode* root) {

        
    }
private:
    void in_order(TreeNode* root){
        if (root==nullptr){return;}
        in_order(root->left);
        if (previous){
            if (previous->val < root->val){
                previous = root;
            }
            else{
                flag = false;
                return;
            }
        }
        else{
            previous = root;
        }
        in_order(root->right);
    }
    bool flag = false;
    TreeNode* previous = nullptr;
};
