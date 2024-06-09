# 433. Minimum Genetic Mutation


## Level - medium


## Task
A gene string can be represented by an 8-character long string, with choices from 'A', 'C', 'G', and 'T'.

Suppose we need to investigate a mutation from a gene string startGene to a gene string endGene where one mutation is defined as one single character changed in the gene string.
- For example, "AACCGGTT" --> "AACCGGTA" is one mutation.

There is also a gene bank bank that records all the valid gene mutations. A gene must be in bank to make it a valid gene string.

Given the two gene strings startGene and endGene and the gene bank bank, return the minimum number of mutations needed to mutate from startGene to endGene. If there is no such a mutation, return -1.

Note that the starting point is assumed to be valid, so it might not be included in the bank.


## Объяснение
У вас есть четыре вида генов: A, C, G и T. Вам даны две строки - начальная и конечная последовательности генов, 
а также массив строк, представляющих все допустимые мутации.

Цель задачи - найти минимальное количество мутаций, необходимых для преобразования начальной последовательности генов в конечную. 
Мутация - это изменение одного гена на другой.

Например, если начальная последовательность генов - "AACCGGTT", конечная - "AACCGGTA", а допустимые мутации - ["AACCGGTA"], 
то минимальное количество мутаций будет 1.

Наша задача - реализовать функцию, которая будет принимать начальную и конечную последовательности генов, 
а также массив допустимых мутаций, и возвращать минимальное количество мутаций, 
необходимых для преобразования начальной последовательности генов в конечную. 
Если преобразование невозможно, верните -1.


## Example 1:
````
Input: startGene = "AACCGGTT", endGene = "AACCGGTA", bank = ["AACCGGTA"]
Output: 1
````

## Example 2:
````
Input: startGene = "AACCGGTT", endGene = "AAACGGTA", bank = ["AACCGGTA","AACCGCTA","AAACGGTA"]
Output: 2
````

## Constraints:
- 0 <= bank.length <= 10
- startGene.length == endGene.length == bank[i].length == 8
- startGene, endGene, and bank[i] consist of only the characters ['A', 'C', 'G', 'T'].