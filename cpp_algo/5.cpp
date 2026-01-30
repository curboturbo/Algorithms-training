#include <iostream>
#include <unordered_map>
#include <vector>

class Solution {
public:
    int majorityElement(std::vector<int>& nums) {
        std::unordered_map<int, int> mapa = {};
        for (int i = 0; i<nums.size();i++){
            if (mapa.count(nums[i])){
                mapa[nums[i]] +=1;
            }
            else{
                mapa[nums[i]] = 1;
            }
        }

        int count = 0;
        int repeat = 0; 
        for (const auto& pair : mapa){
            if (pair.second > repeat){ 
                count = pair.first;
                repeat = pair.second;
            }
        }
    
        return count;
    }
};