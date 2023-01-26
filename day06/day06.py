f = open("input.txt", "r")
text = f.readline()

start = 0
diff = 1
n = 14
for k in range(len(text)):
    flag = False
    for i in range(start,start+diff):
        if text[i] == text[k]:
            start = k
            diff = 1
            flag = True
            break
    if flag == True:
        continue
    diff += 1
    if diff == n:
        print(k)
        break

