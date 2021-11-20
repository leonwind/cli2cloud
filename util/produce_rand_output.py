import random
import string
import time

def get_random_string(length):
    # choose from all lowercase letter
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for _ in range(length))


def main():
    while True:
        print(get_random_string(40))
        time.sleep(1)


if __name__ == "__main__":
    main()
