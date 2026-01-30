#include <iostream>
#include <queue>
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
    int getMinimumDifference(TreeNode* root) {
        inorder(root);
        std::cout<<answer<<std::endl;
        return answer;
    }
private:
    int answer = 1012312;
    int last = 0;
    void inorder(TreeNode* root) {
    if (!root) return;
    inorder(root->left);
    std::cout << root->val << std::endl;
    answer = std::min(answer, root->val - last);
    last = root->val;
    inorder(root->right);
}
};
int main() {
    // создаём узлы
    TreeNode* root = new TreeNode(236);
    TreeNode* n104 = new TreeNode(104);
    TreeNode* n701 = new TreeNode(701);
    TreeNode* n227 = new TreeNode(227);
    TreeNode* n911 = new TreeNode(911);

    // связываем их как в массиве
    root->left = n104;
    root->right = n701;

    n104->right = n227;
    n701->right = n911;
    auto cl = Solution();
    // для проверки выведем inorder
    cl.getMinimumDifference(root); // 104 227 236 701 911
    return 0;
}