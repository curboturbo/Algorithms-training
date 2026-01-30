#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>
//6 zigzag conversation
std::string main(){
    std::string s = "АВ";
    int row = 1;
    std::vector<std::string> d = {};
    for (int i = 0;i<row;i++){
        d.push_back("");
    }
    if (row == 1){
        return s;
    }
    int direction = -1;
    int start = 0;
    row = row - 1;
    for (int i = 0;i<s.length();i++){
        if (direction == -1){
            if (start < row){
                d[start] += s[i];
                start += 1;
            }
            else if (start == row){
                d[start] += s[i];
                direction = 1;
                start -=1;
            }
        }
        else{
            if (start > 0){
                d[start] += s[i];
                start -=1;
            }
            else if (start == 0){
                d[start] += s[i];
                direction = -1;
                start += 1;
            }
        }
    }
    std::string result = "";
    for (int i =0;i<d.size();i++){
        result += d[i];
    }
    return result;
}