# D. Hack the Polynomial Hash

## Info

Нужно вывести N различных строк, у которых полиномиальный хеш по модулю 128 совпадает при одном и том же неизвестном x.
Ключевая идея - собрать строку из блоков Туэ-Морса длины 64 и общего суффикса длины 7.
Итоговая сложность - O(N*455) по времени и O(N*455) по памяти.

## Level - hard

## Task

Consider a hash table with 128 cells.

The keys of this hash table are strings.
To decide which cell to search for a given string, the following hash function is used:

```python
def get_hash(string):
    result = 0
    for char in string:
        result = (result * x + ord(char)) % 128
    return result
```

The value of x is unknown to you, but you know that it is an integer from 2 to 127.

Try to hack this hash function: construct N pairwise distinct strings such that get_hash returns the same value for each of them.
The value x can be any number from 2 to 127, but it is the same for all strings.

### Input

The input contains one integer N (1 <= N <= 128).

### Output

Print N distinct strings consisting of lowercase Latin letters.
The length of each string must not exceed 5000 characters.

The hash values of all printed strings must be equal.

## Объяснение

Хеш-функция вычисляет полином от x по модулю 128, где коэффициенты - это ASCII-коды символов строки.
Модуль 128 равен 2^7, поэтому у чётных и нечётных x поведение хеша разное.
Для чётного x выполнено x^7 = 0 (mod 128), значит хеш зависит только от последних семи символов строки.
Поэтому все строки можно закончить одним и тем же суффиксом длины 7, и для чётных x они сразу станут равны.
Для нечётного x по теореме Эйлера x^32 = 1 (mod 128).
Возьмём блок A длины 64, где символы выбраны по последовательности Туэ-Морса, и блок B, который получается заменой 'a' на 'b' и наоборот.
Разность полиномов блоков A и B равна произведению (1-x)(1-x^2)(1-x^4)(1-x^8)(1-x^16)(1-x^32).
При нечётном x последний множитель равен нулю, поэтому хеши блоков A и B совпадают по модулю 128.
Соберём строку из 7 таких блоков, выбирая в каждой позиции A или B, и добавим суффикс длины 7.
Разные комбинации блоков дают 2^7 = 128 различных строк, а разность хешей любой пары равна нулю при любом x от 2 до 127.
Остаётся вывести первые N из этих строк.

Возможный подход:

- Thue-Morse blocks (блоки Туэ-Морса). Собрать N строк из блоков длины 64 и общего суффикса, время O(N*455).

## Example 1

```text
Input:
1

Output:
test
```

Explanation: for N = 1 any single lowercase string is a valid answer, for example "test".

## Constraints

- 1 <= N <= 128
- Each output string must consist of lowercase Latin letters.
- The length of each output string must not exceed 5000.
- All output strings must be pairwise distinct.
- There must exist a common x from 2 to 127 with equal hash values for all output strings.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Thue-Morse blocks](solve1_thue_morse)
