#include <iostream>
#include <vector>
#include <stack>
#include <algorithm>

std::string simplifyPath(std::string path) {
    std::stack<std::string> stack = {};
    std::string cur = "";
    path += "/";
    for (int i=0;i<path.length();i++){
        if (path[i]=='/' && cur!=""){
            if (cur == "."){
                cur = "";
                continue;
            }
            else if (cur==".."){
                cur = "";
                if (!stack.empty()){
                    stack.pop();
                }else{continue;} // no way upper than the root dir
            }
            else{
                stack.push(cur);
            }
            cur = "";

        }
        else if (path[i]!='/'){
            cur += path[i];
        }
    }

    while (stack.size() > 0){
        cur = '/'+stack.top() + cur;
        stack.pop();
    }
    return cur==""?"/":cur;
}
int main(){
    std::string ass = "/.../a/../b/c/../d/./";
    std::cout<<simplifyPath(ass)<<std::endl;
}