#include <iostream>
#include <vector>
#include <algorithm>


int main(){
        std::vector<int> citations = {100}; //0 1 3 6 5
        int n = citations.size();
        sort(citations.begin(), citations.end());

        for (int i = 0; i < n; i++) {
            if (citations[i] >= n - i) {
                std::cout<<n-i<<std::endl;
                return n - i;
            }
        }
        return 0;
    }

