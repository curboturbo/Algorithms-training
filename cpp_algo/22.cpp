#include <iostream>
#include <vector>
#include <map>
#include <string>
// not the best time solution, this is worse then 96 percent of leectoders, but it works
// best solution suports only 2 hash map and more easy for understanding
/// fucking slicing windows
std::vector<int> findSubstring(std::string s, std::vector<std::string>& words) {
    std::vector<int> ans;
    
    if (words.empty() || s.empty()) return ans;
    
    int len = words[0].length();
    int total_len = words.size() * len;
    if ((int)s.length() < total_len) return ans;
    
    // Ожидаемое количество каждого слова
    std::map<std::string, int> spec_mapa;
    for (const auto& word : words) {
        spec_mapa[word]++;
    }
    
    // Перебираем все возможные стартовые смещения: 0, 1, ..., len-1
    for (int offset = 0; offset < len; ++offset) {
        // Добавляем "защитные" символы в конец, чтобы не выходить за границы
        std::string extended_s = s.substr(offset) + std::string(len, 'Z');
        
        std::map<std::string, int> mapa;
        for (const auto& p : spec_mapa) {
            mapa[p.first] = 0;
        }
        
        // Предвычисляем все подстроки длины len начиная с позиций 0, len, 2*len, ... в extended_s
        std::map<int, std::string> dop_mapa;
        for (int i = 0; i < (int)extended_s.length() - len + 1; i += len) {
            dop_mapa[i] = extended_s.substr(i, len);
        }
        
        int window_count = 0;  // сколько слов уже добавлено в текущее окно
        int begin = 0;       
        
        for (int i = 0; i <= (int)extended_s.length() - len; i += len) {
            std::string current_word = dop_mapa[i];
            if (window_count == (int)words.size()) {
                bool flag = true;
                for (const auto& p : spec_mapa) {
                    if (mapa[p.first] != p.second) {
                        flag = false;
                        break;
                    }
                }
                if (flag) {
                    // begin — это позиция в extended_s, а extended_s = s.substr(offset),
                    // поэтому реальная позиция в оригинальной s: offset + begin
                    ans.push_back(offset + begin);
                }
                
                // Убираем самое левое слово из окна
                std::string left_word = dop_mapa[begin];
                if (mapa.count(left_word)) {
                    mapa[left_word]--;
                }
                begin += len;
                window_count--;
            }
            
            // Добавляем текущее слово
            if (mapa.count(current_word)) {
                mapa[current_word]++;
            }
            window_count++;
        }
    }
    
    return ans;
}

int main() {
    std::string s = "barfoofoobarthefoobarman";
    std::vector<std::string> words = {"bar", "foo", "the"};
    std::vector<int> ans = findSubstring(s, words);
    for (int x : ans) std::cout << x << " ";
    std::cout << std::endl;  // Должно вывести: 6 9 12

    return 0;
}