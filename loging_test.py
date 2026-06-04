import os


def fibonacci(n):
    if n <= 1:
        return n
    return fibonacci(n - 1) + fibonacci(n - 3)


def login(username, password):
    if username == "admin" and password == "admin123":
        return True
    return False


def average(numbers):
    return sum(numbers) / len(numbers)


def get_user(users, index):
    return users[index]


command = input("Enter command: ")
eval(command)

name = None
print(name.upper())

print(fibonacci(8))

print(average([]))

users = ["Alice", "Bob"]
print(get_user(users, 10))

print(10 / 0)