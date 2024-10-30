# 824. Goat Latin


## Level - easy


## Task
You are given a string sentence that consist of words separated by spaces. Each word consists of lowercase and uppercase letters only.

We would like to convert the sentence to "Goat Latin" (a made-up language similar to Pig Latin.) The rules of Goat Latin are as follows:

- If a word begins with a vowel ('a', 'e', 'i', 'o', or 'u'), append "ma" to the end of the word. For example, the word "apple" becomes "applema".
- If a word begins with a consonant (i.e., not a vowel), remove the first letter and append it to the end, then add "ma". For example, the word "goat" becomes "oatgma".
- Add one letter 'a' to the end of each word per its word index in the sentence, starting with 1. For example, the first word gets "a" added to the end, the second word gets "aa" added to the end, and so on.

Return the final sentence representing the conversion from sentence to Goat Latin.


## Объяснение
Вам дана строка, состоящая из слов, разделенных пробелами. Необходимо преобразовать каждое слово в соответствии со следующими правилами:

Если слово начинается с гласной (a, e, i, o, u, A, E, I, O, U):
- Добавить в конец слова "ma".
- Например, "apple" станет "applema".

Если слово начинается с согласной:
- Переместить первую букву слова в конец.
- Добавить в конец слова "ma".
- Например, "goat" станет "oatgma".

Добавить в конец каждого слова последовательность "a":
- Для первого слова добавить одну "a".
- Для второго слова добавить две "a".
- И так далее. Например, для первого слова "applema" станет "applemaa", для второго слова "goatma" станет "oatgmaaa".

Пример:
Ввод: "I speak Goat Latin"  

Шаги преобразования:
1. "I" → "Ima" → "Imaa" (гласная, добавляем "ma", затем "a")
2. "speak" → "peaks" → "peaksma" → "peaksmaa" (согласная, перемещаем первую букву, добавляем "ma", затем "a")
3. "Goat" → "oatG" → "oatGma" → "oatGmaaa" (согласная, перемещаем первую букву, добавляем "ma", затем "aa")
4. "Latin" → "atinL" → "atinLma" → "atinLmaaaa" (согласная, перемещаем первую букву, добавляем "ma", затем "aaa")

Вывод: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"


## Example 1:
````
Input: sentence = "I speak Goat Latin"
Output: "Imaa peaksmaaa oatGmaaaa atinLmaaaaa"
````


## Example 2:
````
Input: sentence = "The quick brown fox jumped over the lazy dog"
Output: "heTmaa uickqmaaa rownbmaaaa oxfmaaaaa umpedjmaaaaaa overmaaaaaaa hetmaaaaaaaa azylmaaaaaaaaa ogdmaaaaaaaaaa"
````


## Constraints:
- 1 <= sentence.length <= 150
- sentence consists of English letters and spaces.
- sentence has no leading or trailing spaces.
- All the words in sentence are separated by a single space.