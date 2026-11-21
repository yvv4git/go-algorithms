# B. Самораспространяющийся вирус в "Ахинтек"

## Info

Нужно построить стратегию распространения вируса по сети компьютеров с учётом рабочих часов устройств.
Каждая активная копия один раз рассылает копии не более чем $k$ соседям.
Копия выживает, только если успевает распаковаться до выключения устройства.
Ответ - не оптимальная стратегия, а любая с как можно большим числом успешных закреплений.

## Level - hard

## Task

Inside the internal network of the huge corporation "Ahintek" there are $N$ computers.
The network is isolated from the internet, and the computers are connected by bidirectional channels.
At hour $x$ of day zero a pre-installed self-replicating virus activates on computer $p$.
The virus packs, encrypts, and sends its copies over the channels.

To become active, a copy must be delivered to a device and immediately start unpacking and decrypting.
On computer $i$ this takes $t_i$ hours, because different devices have different protection and response systems.
The virus can work (unpack, decrypt, send) only while the computer is on.
Once unpacked, the virus enters the "blind zone" of the protection system.
There it encrypts and packs new copies and sends them to directly reachable devices.
Encryption, packing, and sending happen instantly as soon as a copy becomes active.

Each active copy can send its new packed and encrypted copies to at most $k$ different adjacent devices in a single dispatch.
Copies cannot postpone their dispatch: the protection systems would detect and eliminate them.
A single copy cannot send two copies to the same device in one dispatch.

Every computer is on only during its working hours.
Device $i$ works from hour $a_i$ for $b_i$ hours every day; days are cyclic, so the working period may cross midnight.
A copy can start unpacking only on an on and free device (one where no other copy is unpacking).
A copy sent to a switched-off device dies; a copy also dies if the computer switches off before the copy finishes unpacking.

The virus wants to spread and activate as many copies as possible.

### Formal rules

- Time is measured in whole hours from midnight of day zero; moment $T$ corresponds to hour $T \bmod 24$ of the day.
- All events (delivery, start and finish of unpacking, dispatch) happen only at integer moments.
- Computer $i$ is on during the segments $[a_i + 24d, a_i + 24d + b_i]$ for all integers $d$, including both ends.
- Since $b_i \le 23$, there is always a gap between neighbouring segments when the computer is off.
- At integer moment $T$ computer $i$ is on exactly when $(T - a_i) \bmod 24 \le b_i$.
- The initial copy is active on computer $p$ at moment $x$ and counts as the first anchoring; at the same moment it performs its dispatch.
- Delivery is instant: a copy sent at moment $T$ arrives at the receiver at the same moment $T$.
- A copy arriving at moment $s$ on device $i$ starts unpacking if at moment $s$ the device is on and free.
- If several copies arrive at the same moment at a free and on device, exactly one starts unpacking and the others die at once.
- A copy arriving at an off or busy device dies.
- A copy that started unpacking at moment $s$ on device $i$ occupies it on $[s, s + t_i)$ and finishes at $s + t_i$.
- Unpacking succeeds if the device is on during the whole segment $[s, s + t_i]$; switching off exactly at $s + t_i$ is allowed.
- This is equivalent to $(s - a_i) \bmod 24 + t_i \le b_i$.
- If the device switches off strictly before $s + t_i$, the copy dies at the switch-off moment and no longer occupies the device.
- At moment $s + t_i$ the unpacked copy becomes active and counts as an anchoring.
- At the same moment the copy immediately dispatches to at most $k$ adjacent devices.
- Starting from moment $s + t_i$ device $i$ is free again, and a new copy may arrive at that same moment.
- The process lasts $m$ hours: events may happen only at moments $x, x + 1, \ldots, x + m - 1$.
- Nothing happens at moment $x + m$ or later; a copy counts only if it finishes unpacking no later than $x + m - 1$.
- The same device may be infected many times at different moments, and every anchoring counts separately.

Write a program that builds a virus propagation strategy: whom and when to send copies to.
The goal is as many total anchorings as possible, including the initial copy on device $p$.
Finding an optimal strategy is not required: the more anchorings, the more points.

### Input

The first line contains three integers $N$, $p$, and $k$.
Here $N$ is the number of computers, $p$ is the initially infected device, and $k$ is the dispatch limit.
Here $1 \le N \le 10^5$, $1 \le p \le N$, and $1 \le k \le N$.

The second line contains two integers $x$ and $m$: the activation hour and the number of hours allotted to the process.
Here $0 \le x \le 23$, $1 \le m \le 10^6$, and it is guaranteed that $N \cdot m \le 10^6$.
It is guaranteed that device $p$ is on at moment $x$.

The next $N$ lines contain two integers $a_i$ and $b_i$, where $0 \le a_i \le 23$ and $1 \le b_i \le 23$.
Device $i$ is on during the segments $[a_i + 24d, a_i + 24d + b_i]$ for all integers $d$.

The next $N$ lines contain one integer $t_i$ each: the unpacking time in hours on device $i$.
Here $1 \le t_i \le 23$.

Then follow $N$ blocks, each terminated by the number $0$.
Block $i$ contains the numbers of devices directly connected to device $i$ (devices are numbered from one, and $0$ is only a terminator).
Every number, including the terminating $0$, is placed on a separate line.
A block for a device without neighbours consists of a single number $0$.

The network is an undirected graph: if device $j$ appears in the block of device $i$, then device $i$ appears in the block of device $j$.
No device is connected to itself, and numbers inside one block do not repeat.
The number of channels (edges) does not exceed $5 \cdot 10^5$, so the total number of device numbers in all blocks does not exceed $10^6$.

### Output

In the first line print an integer $A$ ($A \ge 1$): the number of anchorings in your strategy, including the initial copy on device $p$.

Then print $A - 1$ lines, one per anchoring except the initial one.
Each line contains three integers $u$, $v$, and $s$.
At moment $s$ the active copy on device $u$ sent a copy to device $v$, and it started unpacking at $s$ and succeeded.
Lines may be printed in any order.
Dead copies must not be printed.

## Объяснение

В задаче есть сеть из $N$ компьютеров и вирус, который распространяется по каналам.
Каждая активная копия может один раз отправить копии не более чем $k$ соседям, и это происходит мгновенно.
Копия выживает, только если попала на включённое свободное устройство и успела распаковаться, пока устройство не выключилось.
Распаковка на устройстве $i$ занимает $t_i$ часов и должна целиком поместиться в рабочий интервал устройства.

Устройство $i$ включено на отрезках $[a_i + 24d, a_i + 24d + b_i]$ для всех целых $d$.
Поэтому в момент $s$ копия может начать распаковку, только если $(s - a_i) \bmod 24 + t_i \le b_i$.
Если условие выполнено, копия занимает устройство до момента $s + t_i$ и затем становится активной.
В момент активации она засчитывается как закрепление и сразу делает свою рассылку.

Важно, что одно устройство можно заражать много раз в разные моменты.
Поэтому выгодно не только продвигаться в новые вершины, но и возвращаться в уже посещённые, чтобы получить дополнительные закрепления.

Простая стратегия - жадно моделировать процесс по часам.
Будем хранить очередь активных копий и для каждой копии выбирать до $k$ соседей, на которых распаковка успеет завершиться.
Если копий на одно устройство в один момент несколько, оставляем одну, остальные считаем погибшими.
Такое моделирование даёт корректную стратегию, но не обязательно максимальное число закреплений.

Более аккуратный подход - событийное моделирование с приоритетной очередью.
События (активация копии, освобождение устройства) обрабатываются в порядке времени.
При рассылке выбираются соседи с наиболее ранним подходящим окном работы, чтобы успешно завершить распаковку.
Это уменьшает число погибших копий и увеличивает итоговое число закреплений.

Оптимизированная версия событийного моделирования хранит события в плоских массивах по времени, а не в приоритетной очереди.
Соседи каждого устройства заранее сортируются counting sort по времени распаковки, а для каждого часа хранится битовая маска.
Быстрый ввод и вывод уменьшают накладные расходы на больших тестах.

Возможные подходы:

- Greedy hour-by-hour simulation (жадное моделирование по часам). Перебирать моменты времени и рассылать копии в подходящих соседей.
- Event-driven simulation (событийное моделирование). Хранить события в приоритетной очереди и обрабатывать их по времени.
- Optimized event-driven simulation (оптимизированное событийное моделирование). Хранить события в плоских массивах по времени.

## Example 1

```text
Input:
4 1 1
10 13
10 9
6 11
9 4
6 4
2
2
1
1
2
0
1
3
0
2
4
0
3
0

Output:
5
1 2 10
2 3 12
3 2 13
2 1 15
```

Explanation: the initial copy is anchored on device 1 at moment 10.
It sends a copy to device 2, which unpacks and anchors at moment 12.
Then device 2 sends to device 3, and device 3 sends back to device 2.
Finally device 2 sends to device 1 again, giving five anchorings in total.

## Constraints

- $1 \le N \le 10^5$
- $1 \le p \le N$
- $1 \le k \le N$
- $0 \le x \le 23$
- $1 \le m \le 10^6$
- $N \cdot m \le 10^6$
- $0 \le a_i \le 23$
- $1 \le b_i \le 23$
- $1 \le t_i \le 23$
- The number of edges does not exceed $5 \cdot 10^5$.
- Device $p$ is on at moment $x$.

## См. также

- [Подходы и их сравнение](approachs.md)
- [Greedy hour-by-hour simulation](solve1_greedy_hour_by_hour)
- [Event-driven simulation](solve2_event_driven)
- [Optimized event-driven simulation](solve3_optimized_event_driven)
