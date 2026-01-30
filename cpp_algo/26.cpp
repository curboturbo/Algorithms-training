#include <iostream>
#include <vector>
#include <string>

std::vector<std::string> summaryRanges(std::vector<int>& nums) {
    std::vector<std::string> ans;
    if (nums.empty()) return ans;

    int begin = nums[0];
    int end = nums[0];

    for (int i = 1; i < nums.size(); i++) {
        if (nums[i] == end + 1) {
            end = nums[i];
        } else {
            if (begin == end)
                ans.push_back(std::to_string(begin));
            else
                ans.push_back(std::to_string(begin) + "->" + std::to_string(end));
            begin = end = nums[i];
        }
    }
    if (end==begin){
        ans.push_back(std::to_string(begin));
    }else{
        ans.push_back(std::to_string(begin) + "->" + std::to_string(end));
    }
    return ans;
}


int main(){
    std::vector<int> nums = {0,1,2,4,5,7};
    summaryRanges(nums);
}