# 1106. Parsing A Boolean Expression


## Level - hard


## Task
A boolean expression is an expression that evaluates to either true or false. 
It can be in one of the following shapes:
- 't' that evaluates to true.
- 'f' that evaluates to false.
- '!(subExpr)' that evaluates to the logical NOT of the inner expression subExpr.
- '&(subExpr1, subExpr2, ..., subExprn)' that evaluates to the logical AND of the inner expressions subExpr1, subExpr2, ..., subExprn where n >= 1.
- '|(subExpr1, subExpr2, ..., subExprn)' that evaluates to the logical OR of the inner expressions subExpr1, subExpr2, ..., subExprn where n >= 1.

Given a string expression that represents a boolean expression, return the evaluation of that expression.
It is guaranteed that the given expression is valid and follows the given rules.


## Объяснение
Это задача на парсинг булевых выражений, которые могут содержать логические операторы (!, &, |) и скобки. 
Выражения могут быть вложенными и содержать несколько операторов.

Например, выражение !(f) будет интерпретировано как NOT f, а выражение &(t,f) будет интерпретировано как AND(true, false).

Задача требует написать функцию, которая принимает на вход строку, представляющую булево выражение, и возвращает результат вычисления этого выражения.


## Решение с помощью стека
Одним из подходов к решению этой задачи является использование стека для отслеживания вложенных выражений и операторов. 
Алгоритм будет проходить по строке справа налево и выполнять следующие действия:
1. Если текущий символ является закрывающей скобкой ), то он помещается в стек.
2. Если текущий символ является открывающей скобкой (, то из стека извлекаются выражения и операнды, 
пока не будет достигнута соответствующая закрывающая скобка. 
Затем вычисляется результат выражения и результат помещается обратно в стек.
3. Если текущий символ является логическим оператором (!, &, |), то он также помещается в стек.
4. Если текущий символ является булевым значением (t или f), то он преобразуется в соответствующее булево значение и помещается в стек.

В результате, стек будет использоваться для отслеживания вложенных выражений и их результатов, 
и в конце алгоритма будет получен результат вычисления булевого выражения.


## Example 1:
````
Input: expression = "&(|(f))"
Output: false
Explanation:
First, evaluate |(f) --> f. The expression is now "&(f)".
Then, evaluate &(f) --> f. The expression is now "f".
Finally, return false.
````


## Example 2:
````
Input: expression = "|(f,f,f,t)"
Output: true
Explanation: The evaluation of (false OR false OR false OR true) is true.
````


## Example 3:
````
Input: expression = "!(&(f,t))"
Output: true
Explanation:
First, evaluate &(f,t) --> (false AND true) --> false --> f. The expression is now "!(f)".
Then, evaluate !(f) --> NOT false --> true. We return true.
````


## Constraints:
- 1 <= expression.length <= 2 * 10^4
- expression[i] is one following characters: '(', ')', '&', '|', '!', 't', 'f', and ','.