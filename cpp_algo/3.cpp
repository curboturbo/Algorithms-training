#include <iostream>
#include <vector>


class Solution {
public:
    int removeDuplicates(std::vector<int>& nums) {
        int install = 1;
        int last = nums[0];
        int count = 1;
        int cnt = 0;
        for (int i = 1;i < nums.size(); i++){
            if (nums[i] != nums[i-1]){
                nums[install] = nums[i];
                install += 1;
            }
        }
        return install; 

    }
};