#include <iostream>
#include <string>
#include <vector>
#include <map>
std::string reverseWords(std::string s) {
    std::map<int, std::string> mapa;
    std::string test = "";
    int cnt = 0;

    for (int i = 0; i < s.length(); i++) {
        if (s[i] != ' ') {
            test += s[i];
        } else {
            if (test != "") {
                mapa[cnt] = test;
                cnt++;
                test = "";
            }
        }
    }
    if (test != "") {
        mapa[cnt] = test;
        cnt++;
    }
    std::string answer = "";
    for (int i = cnt - 1; i >= 0; i--) {
        answer += mapa[i];
        if (i > 0) {
            answer += ' ';
        }
    }
    return answer;
}

int main(){
    std::cout<<reverseWords("fuck the world")<<std::endl;

}