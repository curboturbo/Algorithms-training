#include <iostream>
#include <vector>
#include <algorithm>


int main(){
    std::vector<int> num = {1,2,3,4};
    std::vector<int> left = {1};
    std::vector<int> right = {1};

    int cnt = 0;
    for (int i = 1;i<num.size();i++){
        left.push_back(left[i-1]*num[i-1]);    
    }
    for (int i = num.size()-2;i>=0;i--){
        right.push_back(right[cnt]*num[i+1]);
        cnt += 1;
    }

    std::reverse(right.begin(),right.end());
    std::vector<int> ans = {};
    for (int i = 0;i<num.size();i++){
        ans.push_back(left[i]*right[i]);
    }
}

// БОЛЕЕ ОПТИМИЗИРОВАННЫЙ МЕТОД
class Solution {
public:
    std::vector<int> productExceptSelf(std::vector<int>& nums) {
        std::vector<int> output(nums.size(), 1);

        int left = 1;
        for (int i = 0; i < nums.size(); i++) {
            output[i] *= left;
            left *= nums[i];
        }

        int right = 1;
        for (int i = nums.size() - 1; i >= 0; i--) {
            output[i] *= right;
            right *= nums[i];
        }

        return output;        
    }
};