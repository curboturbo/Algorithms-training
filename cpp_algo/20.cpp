#include <iostream>
#include <vector>


// algortihm KMP


int main(){
    std::string input = "sadbutsad";
    std::string a = "sad";
    int n = input.length();
    std::vector<int> pi(n,0);
    int j = 0;
    int i = 1;
    while (i < a.length()){
        if (a[j] == a[i]){
            pi[i] = j+1;
            i +=1;
            j +=1;
        }
        else{
            if (j == 0){
                pi[i] = 0;
                i += 1;
            }
            else{
                j = pi[j-1];
            }
        }
    }
    i = 0;
    j = 0;
    while (i<input.length()){
        if (input[i] == a[j]){
            i +=1;
            j +=1;
            if (j == a.length()){
                return i-j;
                std::cout<<i-j<<std::endl;
                break;
            }
        }
        else{
            if (j > 0){
                j = pi[j-1];
            }
            else{
                i += 1;
            }
        }

    }
    if (i==n){

        return -1;
        //std::cout<<-1<<std::endl;
    }


}