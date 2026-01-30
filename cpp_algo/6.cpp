#include <iostream>
#include <vector>


void swap(std::vector<int>& nums, int left, int right) {
    while (left < right) {
        int cur = nums[right];
        nums[right] = nums[left];
        nums[left] = cur;
        right -= 1;
        left += 1;
    }
}


int main(){
    std::vector<int> nums = {1,2,3,4,5,6,7,8};
    int k=3;
    k = k % nums.size();
    swap(nums,0,nums.size()-1);
    swap(nums,0,k-1);
    swap(nums,k,nums.size()-1);
    for (int i = 0; i < nums.size(); i++) {
        std::cout << nums[i] << std::endl;
    }
}