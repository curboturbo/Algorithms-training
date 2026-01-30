#include <iostream>
#include <vector>

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
    TreeNode* sortedArrayToBST(std::vector<int>& nums) {
        int left = 0;
        int right = nums.size()-1;
        return partition(left,right,nums);
    }
private:
    TreeNode* partition(int left,int right,std::vector<int>& nums){
        if (left>right){return nullptr;}
        int mid = left + (right - left) / 2;
        TreeNode* el = new(TreeNode);
        el->val = nums[mid];
        el->left = partition(left,mid-1,nums);
        el->right = partition(mid+1,right,nums);
        
        return el;
    }
};