#include <iostream>
#include <vector>
#include <map>

std::vector<int> twoSum(std::vector<int>& nums, int target) {
    std::map<int, int> seen;
    for (int i = 0; i < nums.size(); i++) {
        int complement = target - nums[i];
        if (seen.find(complement) != seen.end()) {
            return {seen[complement], i};
        }
        seen[nums[i]] = i;
    }
    return {};
}

int main() {
    std::vector<int> name = {2, 7, 11, 15};
    std::vector<int> result = twoSum(name, 9);
    
    if (!result.empty()) {
        std::cout << "[" << result[0] << ", " << result[1] << "]" << std::endl;
    } else {
        std::cout << "No solution found" << std::endl;
    }
    
    return 0;
}