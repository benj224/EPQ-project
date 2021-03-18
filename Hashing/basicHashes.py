import hashlib

#SHA256 HASH
print(hashlib.sha256(b"Hello World!").hexdigest())
print()

#MD5 HASH
print(hashlib.md5(b"Hello World!").hexdigest())
print()


#SHA224 HASH
print(hashlib.sha224(b"Hello World!").hexdigest())
print()


#HASHING A VARIABLE
item = "Hello World!"
print(hashlib.sha256(item.encode()).hexdigest())
print()


#UPDATING A HASH
item1 = hashlib.sha256()
item1.update(b"New String!")
print("Start: ", item1.hexdigest())
item1.update(b"Changed String!")
print("End  : ", item1.hexdigest())
print()

#DIGEST SIZE
item2 = hashlib.sha256()
item2.update(b"Hello World!")
print("Digest Size: ", item2.digest_size)

#BLOCK SIZE
item3 = hashlib.sha256()
item3.update(b"Hello World!")
print("Block Size: ", item3.block_size)

#COPYING
item4 = hashlib.sha256()
item4.update(b"Hello World!")
print("Original: ", item4.hexdigest())
item5 = item4.copy()
print("Copied  : ", item5.hexdigest())
