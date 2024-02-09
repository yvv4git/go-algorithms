num = int(input())

print(num)

while num != 1:
    if num % 2 == 0:
        num = int(num / 2)
        print(num, end=" ")
    else:
        num *= 3
        num += 1
        print(num, end=" ")