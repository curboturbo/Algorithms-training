#include <stack>
#include <iostream>
#include <vector>
#include <string>
#include <map>

int define(int x, int y, std::string oper){
    if (oper == "/"){return x/y;}
    else if (oper=="*"){return x*y;}
    else if (oper=="+"){return x+y;}
    else{return x-y;}
}

int evalRPN(std::vector<std::string>& tokens) {
    std::stack<int> stack = {};
    std::map<std::string,int> mapa;
    mapa["+"] = 0;
    mapa["-"] = 0;
    mapa["/"] = 0;
    mapa["*"] = 0;
    int result = 0;

    for (auto& element: tokens){
        if (mapa.count(element)){
            int a = stack.top();
            stack.pop();
            int b = stack.top();
            stack.pop();
            stack.push(define(b,a,element));
        }else{stack.push(stoi(element));}
    }
    return stack.top();
}
int main(){
    std::vector<std::string> vec = {"4","13","5","/","+"};
    std::cout<<evalRPN(vec)<<std::endl;
}