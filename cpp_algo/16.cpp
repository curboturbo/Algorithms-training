#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>

int main(){
    std::vector<int> height = {4,2,0,3,2,5};
    std::vector<int> left = {height[0]};
    std::vector<int> right = {};
    int n = height.size();
    int V = 0;
    for (int i =0;i<height.size();i++){
        right.push_back(0);
    }

    int mx_l = left[0];
    

    for (int i = 1;i <height.size();i++){
        left[i] = mx_l;
        mx_l = std::max(mx_l,height[i]);
    }

    right[right.size()-1] = height[right.size()-1];
    int mx_r = height[right.size()-1];

    for (int i = right.size()-2;i>=0;i--){
        right[i] = mx_r;
        mx_r = std::max(mx_r, height[i]);
    }



    for (int i =0 ;i<n;i++){
        if (std::min(right[i],left[i]) - height[i] > 0){
            V += std::min(right[i],left[i]) - height[i];
        }
    }
    return V;
    std::cout<<V<<std::endl;
}