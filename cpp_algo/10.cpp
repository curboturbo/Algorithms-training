#include <iostream>
#include <algorithm>
#include <vector>

#include <iostream>
#include <vector>

int main() {
    std::vector<int> nums = {2, 3, 1, 1, 4};
    int goal = nums.size() - 1; 
    int end = nums.size() - 2;  
    int point = 0;              
    int m = goal;               

    while (goal != 0) {
        bool found = false; 
        for (int i = end; i >= 0; i--) {
            if (nums[i] + i >= goal) {
                m = i; 
                found = true;
            }  
        }
        if (!found) {
            return 1;
        }

        end = m - 1; 
        goal = m;    
        point += 1;  
    }
    return point;
}