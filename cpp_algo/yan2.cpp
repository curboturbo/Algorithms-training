#include <iostream>
#include <vector>
#include <string>
#include <unordered_map>

int main(){ 
    std::string s;
    std::cin >> s;
    int n = s.length();
    int all_combine = (s.length()*(s.length()-1))/2;
    std::unordered_map<int,int> mapa = {};
    for (int i =0;i<s.length();i++){
        mapa[s[i]] += 1;
    }      
    for (auto& key_value : mapa) {
        all_combine -= (key_value.second * (key_value.second - 1)) / 2;
    }
    std::cout<< all_combine + 1<<std::endl;

}



























//using namespace std;
//
//
//int main() {
//    int n, m;
//    std::cin >> n >> m;
//
//    std::string s;
//    std::cin >> s;
//
//    std::vector<std::string> parts(m);
//    for (int i = 0; i < m; ++i) {
//        std::cin >> parts[i];
//    }
//
//    int k = n / m;
//
//    std::unordered_map<std::string, std::vector<int>> part_dict;
//    for (int i = 0; i < m; ++i) {
//        part_dict[parts[i]].push_back(i + 1);
//    }
//
//    std::vector<int> result;
//    for (int i = 0; i < n; i += k) {
//        std::string segment = s.substr(i, k);
//        if (!part_dict[segment].empty()) {
//            result.push_back(part_dict[segment].back());
//            part_dict[segment].pop_back();
//        }
//    }
//
//    // Вывод результата
//    for (int i = 0; i < result.size(); ++i) {
//        if (i > 0) cout << ' ';
//        cout << result[i];
//    }
//    cout << '\n';
//
//    return 0;
//}
//
//
//
//
//
//using namespace std;
//
//int main() {
//    long long n, k;
//    if (!(cin >> n >> k)) return 0;
//
//    if (k == 0 || n % 10 == 0) {
//        cout << n << '\n';
//        return 0;
//    }
//
//    while (k > 0) {
//        int d = n % 10;
//        if (d == 0) break;
//        if (d == 2) break; 
//        n += d;
//        --k;
//    }
//
//    if (k == 0 || n % 10 == 0) {
//        cout << n << '\n';
//        return 0;
//    }
//
//    if (n % 10 == 2) {
//
//        long long fullCycles = k / 4;
//        n += fullCycles * 20;
//        int rem = (int)(k % 4);
//        int add[4] = {2, 4, 8, 6}; 
//        for (int i = 0; i < rem; ++i) n += add[i];
//        cout << n << '\n';
//        return 0;
//    } 
//    cout << n << '\n';
//    return 0;
//}
//
//
//int main() {
//    int n, k, x;
//    std::cin >> n >> k;
//
//    std::map<int, int> mapa;
//    for (int i = 0; i < n; ++i) {
//        std::cin >> x;
//        mapa[x]++;
//    }
//
//    int out = 0;
//
//
//    for (auto& key_value : mapa) {
//        std::cout << key_value.first << " ";
//        key_value.second--;
//        out++;
//        if (out == k) return 0;
//    }
//
//    for (auto& key_value : mapa) {
//        while (key_value.second > 0 && out < k) {
//            std::cout << key_value.first << " ";
//            key_value.second--;
//            out++;
//        }
//        if (out == k) break;
//    }
//
//    return 0;
//}
//