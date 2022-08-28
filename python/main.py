import csv

MAX = 100000
NUMBERS_FILE = "../numbers.txt"

nums = [0]*MAX

try:
    count = 0

    with open(NUMBERS_FILE, mode='r') as f:
        reader = csv.reader(f, delimiter=',')
        count = 0

        for row in reader:
            if (len(row) < MAX):
                print("data < MAX")
                exit(1)

            for i in range(len(nums)):
                nums[i] = int(row[i])
                count += 1
        
    
    print(nums)
    print("len(nums): ", count)

except FileNotFoundError:
    print("FileNotFoundError")
except IOError:
    print(IOError)
