function hexToDec(letter) {
    if (letter == "1") {
        return 1
    }
    else if (letter == "2") {
        return 2
    }
    else if (letter == "3") {
        return 3
    }
    else if (letter == "4") {
        return 4
    }
    else if (letter == "5") {
        return 5
    }
    else if (letter == "6") {
        return 6
    }
    else if (letter == "7") {
        return 7
    }
    else if (letter == "8") {
        return 8
    }
    else if (letter == "9") {
        return 9
    }
    else if (letter == "a") {
        return 10
    }
    else if (letter == "b") {
        return 11
    }
    else if (letter == "c") {
        return 12
    }
    else if (letter == "d") {
        return 13
    }
    else if (letter == "e") {
        return 14
    }
    else if (letter == "f") {
        return 15
    }
}

function squareThing() {
    var random_string = "";
    var alpha = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

    for (i = 0; i < 12; i++) {
        random_string += alpha.charAt(Math.floor(Math.random() * alpha.length));
    }

    console.log("here");
    var digest = CryptoJS.SHA256(random_string);
    var hexdigest = "\nBase64:\t" + CryptoJS.enc.Base64.parse(String(digest));

    var testHashList = [];
    var hexList = hexdigest.split("");

    for (i = 0; i < hexList.length; i++) {
        testHashList.push(hexList[i]);
    }



    for (i = 0; i < 8; i++) {
        for (j = 0; j < 8; j++) {
            if (i == 0 && j == 0 && hexToDec(testHashList[0] == 1)) {
                color = (0, 255, 0);
            }
            else if (i == 1 && j == 0 && hexToDec(testHashList[1] == 2)) {
                color = (0, 255, 0)
            }
            else if (i == 2 && j == 0 && hexToDec(testHashList[2] == 3)) {
                color = (0, 255, 0)
            }
            else if (i == 3 && j == 0 && hexToDec(testHashList[3] == 4)) {
                color = (0, 255, 0)
            }
            else {
                var num = hexToDec(testHashList[8 * i + j]) * 16;
                var color = (num, 0, 255 - num);
            }
            var coords = [40, 160, 280, 400, 520, 640, 760];
            var coords_length = coords.length;
            var c = document.getElementById("myCanvas");
            var ctx = c.getContext("2d");
            ctx.beginPath();
            ctx.rect(coords[i], 40, 100, 100);
            ctx.rect(coords[i], 160, 100, 100);
            ctx.rect(coords[i], 280, 100, 100);
            ctx.rect(coords[i], 400, 100, 100);
            ctx.rect(coords[i], 520, 100, 100);
            ctx.rect(coords[i], 640, 100, 100);
            ctx.rect(coords[i], 760, 100, 100);
            ctx.fillStyle = color;
            ctx.fill();

        }
    }
}


for(i = 0; i < 10; i++){
    squareThing();
}



