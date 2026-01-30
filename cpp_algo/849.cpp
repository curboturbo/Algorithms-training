#include <iostream>
#include <vector>
#include <algorithm>

class Solution {
public:
    int maxDistToClosest(std::vector<int>& seats) {
        int n = seats.size();
        int answer = 0;
        int last = -1;
        for (int i = 0; i < n; i++) {
            if (seats[i] == 1) {
                if (last==-1) {
                    answer=i;
                } else {
                    answer = std::max(answer, (i - last) / 2);
                }
                last = i; 
            }
        }
        if (last != n - 1) {
            answer = std::max(answer, (n - 1) - last);
        }

        return answer;
    }
};