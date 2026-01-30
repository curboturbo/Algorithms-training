// Longest common prefix 14
#include <iostream>
#include <vector>
#include <algorithm>

int main(){
    std::vector<int> nums ={5,4,-1,7,8,-120,10,20,30};
    int max_sum = nums[0];
    int current = nums[0];
    if (nums.size() == 1){return nums[0];}
    for (int i = 1;i<nums.size();i++){
        current = std::max(nums[i]+current,nums[i]);
        max_sum = std::max(max_sum, current);
    }
    std::cout<<max_sum<<std::endl;
}
