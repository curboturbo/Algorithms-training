#include <iostream>
#include <vector>


class Solution {
public:
    int canCompleteCircuit(std::vector<int>& gas, std::vector<int>& cost) {
        int total_gas = 0;   
        int current_gas = 0; 
        int start_point = 0;  
    
        for (int i = 0; i < gas.size(); i++) {
            int diff = gas[i] - cost[i];
            total_gas += diff;
            current_gas += diff;
    
            if (current_gas < 0) {
                start_point = i + 1;  
                current_gas = 0;     
            }
        }
    
        if (total_gas < 0 || start_point >= gas.size()) {
            return -1;
        }
    
        return start_point;
    }
};