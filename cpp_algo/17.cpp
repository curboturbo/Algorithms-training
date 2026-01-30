#include <iostream>
#include <vector>
#include <algorithm>
#include <numeric>


int main(){
    std::string word = "Asswhole your nigga";
    int ans = 0;
    for (int i = word.length()-1;i>=0;i--){{
        if (word[i] != ' '){
            ans +=1;
        }
        else{
            std::cout<<ans<<std::endl;
            return ans;
        }
    }
        
    }
    

}


class Solution {
public:
    int lengthOfLastWord(std::string s) {
        int end = s.length() - 1;

        while (end >= 0 && s[end] == ' ') {
            end--;
        }

        int start = end;
        while (start >= 0 && s[start] != ' ') {
            start--;
        }

        return end - start;        
    }
};