#include <iostream>
#include <vector>


class Solution {
public:
    bool canJump(std::vector<int>& nums) {
        int goal = nums.size()-1 ;
        int resolve = false;
        if (nums.size()==1){return 1;}
        for (int i = nums.size()-2;i >=0;i--){
            if (i + nums[i] >= goal){
                goal = i;
            }
        }
        return goal == 0;
    }
};
int main(){
    return 1;
}