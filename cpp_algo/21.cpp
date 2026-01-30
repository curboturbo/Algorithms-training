#include <iostream>
#include <vector>

// common prefix on vector of strings
std::string longestCommonPrefix(std::vector<std::string>& strs){
    std::string first = strs[0];
    std::string longest_prefix = "";
    for (int i =1;i<strs.size();i++){
        if (first.length() >= strs[i].length()){
            int cnt = 0;
            while (cnt < strs[i].length()){
                if (first[cnt] == strs[i][cnt]){
                    longest_prefix+=first[cnt];
                    cnt +=1;
                }
                else{break;}
            }
        }
        else{
            int cnt = 0;
            while (cnt < first.length()){
                if (first[cnt] == strs[i][cnt]){
                    longest_prefix+=first[cnt];
                    cnt +=1;
                }
                else{break;}
            }
        }
        first = longest_prefix;
        longest_prefix = "";
    }
    return first;
}


int main(){
    std::vector<std::string> pref = {"flower","flow","zZlight"};
    std::cout<<longestCommonPrefix(pref)<<std::endl;
}