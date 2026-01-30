#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>


int main(){
    std::vector<int> ratings = {1,2,2};
    std::vector<int> candyland = {};
    for (int i = 0;i<ratings.size();i++){
        candyland.push_back(1);
    }
    for (int i=1;i<ratings.size();i++){
        if (ratings[i] > ratings[i-1]){
            candyland[i] = candyland[i-1] + 1;
        }
    }
    for (int i =ratings.size()-2;i>=0;i--){
        if (ratings[i] > ratings[i+1]){
            candyland[i] = std::max(candyland[i], candyland[i+1] + 1);
        }
    }
    
    return std::accumulate(candyland.begin(), candyland.end(), 0);
}