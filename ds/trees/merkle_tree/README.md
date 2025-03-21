# MERKLE TREE


## INFO
Дерево Меркла (Merkle tree) в криптографии — это структура данных, 
используемая для эффективного хранения и проверки целостности больших наборов данных. 
Оно названо в честь его изобретателя, Ральфа Меркла.


## ХАРАКТЕРИСТИКИ
1. Иерархическая структура. 
Дерево Меркла представляет собой бинарное дерево, где каждый узел (кроме листьев) является хешем от двух дочерних узлов. 
Листья дерева содержат хеши от отдельных блоков данных.

2. Хеширование.
Каждый узел дерева Меркла содержит хеш-значение, полученное путем применения хеш-функции к данным или к комбинации хешей дочерних узлов. 
Обычно используются криптографические хеш-функции, такие как SHA-256.

3. Проверка целостности.
Дерево Меркла позволяет быстро проверить целостность данных. 
Если хеш корневого узла (Merkle root) совпадает с ожидаемым значением, то можно быть уверенным, что все данные в дереве не были изменены.

4. Эффективность. 
Проверка целостности данных с использованием дерева Меркла требует значительно меньше вычислительных ресурсов, 
чем проверка каждого отдельного блока данных. Это особенно важно для больших наборов данных.


## EXAMPLE
Предположим, у нас есть четыре блока данных: A, B, C и D.
1. Листья дерева.
Каждый блок данных хешируется, и результаты становятся листьями дерева:  
- H(A)
- H(B)
- H(C)
- H(D)

2. Промежуточные узлы.
Хеши листьев комбинируются и хешируются снова:  
- H1 = H(H(A) + H(B))
- H2 = H(H(C) + H(D))

3. Корень дерева.
Наконец, хеши промежуточных узлов комбинируются и хешируются, чтобы получить корень дерева Меркла:  
Merkle Root = H(H1 + H2)

Если какой-либо блок данных изменится, хеш соответствующего листа изменится, 
что приведет к изменению всех родительских узлов и, в конечном итоге, корня дерева Меркла. 
Таким образом, проверка корня дерева Меркла позволяет быстро определить, были ли изменены данные.