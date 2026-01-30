#include <iostream>
#include <vector>

int main() {
    std::vector<int> a = {10, 20, 30, 40, 50};
    std::vector<int> b = {5, 15, 25}; 
    std::vector<int> result;

    size_t i = 0, j = 0; 
    while (i < a.size() && j < b.size()) {
        if (a[i] <= b[j]) {
            result.push_back(a[i]);
            i++;
        } else {
            result.push_back(b[j]);
            j++;
        }
    }

    
    while (i < a.size()) {
        result.push_back(a[i]);
        i++;
    }


    while (j < b.size()) {
        result.push_back(b[j]);
        j++;
    }


    for (int k : result) {
        std::cout << k << " ";
    }
    std::cout << std::endl;

    return 0;
}


#include <iostream>
#include <string>
#include <vector>

class Solution {
public:
    void merge(std::vector<int>& nums1, int m, std::vector<int>& nums2, int n) {
        int midx = m - 1;
        int nidx = n - 1;
        int right = m + n - 1;

        while (nidx >= 0) {
            if (midx >= 0 && nums1[midx] > nums2[nidx]) {
                nums1[right] = nums1[midx];
                midx--;
            } else {
                nums1[right] = nums2[nidx];
                nidx--;
            }
            right--;
        }        
    }
};