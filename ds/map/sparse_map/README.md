# Sparse map


## Info
SparseMap — это структура данных, которая используется для хранения пар "ключ-значение", где ключи представляют собой целые числа. 
Она оптимизирована для случаев, когда ключи не образуют непрерывного диапазона, то есть многие ключи могут отсутствовать. 
Это позволяет экономить память, так как неиспользуемые ключи не занимают места в памяти.


## Почему эффективна?
- SparseMap не выделяет память для ключей, которые не используются. 
Это делает её более эффективной по сравнению с обычными массивами или HashMap, если ключи разбросаны (sparse).

- В отличие от HashMap, который может работать с любыми типами ключей, SparseMap обычно используется для целочисленных ключей.

- Время доступа к элементам в SparseMap обычно быстрее, чем у HashMap, 
так как она использует более простую структуру данных для хранения ключей и значений.