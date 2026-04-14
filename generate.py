import random
import sys

names = ['arjun', 'alice', 'bob', 'charlie', 'diana', 'eve', 'frank', 'grace', 'henry', 'irene']
count = int(sys.argv[1]) if len(sys.argv) > 1 else 1000000

with open('input.txt', 'w') as f:
    for i in range(count):
        name = random.choice(names)
        f.write(f'{name} | 2026-04-13 10:00 | message {i}\n')

print(f'generated {count} lines -> input.txt')