#include <iostream>
#include <vector>

class Solution {
public:
    int removeDuplicates(std::vector<int>& nums) {
        int cnt = 2;
        int position;
        int last = -1;
        if (nums.size() == 2){return 2;}
        for (int i = 2; i< nums.size(); i++){
            if (nums[i] != nums[cnt-2]){
                nums[cnt] = nums[i];
                cnt += 1;
            }
        }
    }
};