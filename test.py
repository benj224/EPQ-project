import random
import hashlib
import string

x = "".join(random.choice(string.ascii_lowercase) for i in range(10))
x = hashlib.sha256(x.encode()).hexdigest()
difficulty = "000"
while x[:3] != difficulty:
    x = "".join(random.choice(string.ascii_lowercase) for i in range(10))
    print("Unhashed: ", x)
    x = hashlib.sha256(x.encode()).hexdigest()
    print("Hashed: ", x)