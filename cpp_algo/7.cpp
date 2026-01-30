#include <iostream>
#include <vector>
#include <algorithm>

class Solution {
public:
    int maxProfit(std::vector<int>& prices) {
        int m = 1010*10;
        int delta = -1;
        if (prices.size()<=1){return 0;}
        for (int i = 1;i<prices.size();i++){
            m = std::min(m,prices[i-1]);
            delta = std::max(prices[i]-m,delta);
        }
        if (delta < 0){return 0;}
        return delta;

    }
};