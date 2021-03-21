import base64
def sign(privatekey, data):
    return base64.b64encode(str((privatekey.sign(data, ""))[0]).encode())