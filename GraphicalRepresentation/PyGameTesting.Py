import pygame
import random
import hashlib

pygame.init()
pygame.font.init()

myfont = pygame.font.SysFont('Comic Sans MS', 30)
surface = pygame.display.set_mode((0, 0), pygame.FULLSCREEN)

running = True

info = pygame.display.Info()
height = info.current_h
width = info.current_w

blockHeight = height/12
blockWidth = blockHeight

placementGrid = []

sideMarginWidth = (width - height)/2

def hexToDec(letter):
    if letter == "1":
        return 1
    elif letter == "2":
        return 2
    elif letter == "3":
        return 3
    elif letter == "4":
        return 4
    elif letter == "5":
        return 5
    elif letter == "6":
        return 6
    elif letter == "7":
        return 7
    elif letter == "8":
        return 8
    elif letter == "9":
        return 9
    elif letter == "a":
        return 10
    elif letter == "b":
        return 11
    elif letter == "c":
        return 12
    elif letter == "d":
        return 13
    elif letter == "e":
        return 14
    elif letter == "f":
        return 15
    elif letter == "0":
        return 0

# def hexToRGB(hexStr):
#     redHex = hexStr[0:2]
#     blueHex = hexStr[2:]
#     red = hexToDec(redHex[0])*16+hexToDec(redHex[1])
#     blue = hexToDec(blueHex[0])*16+hexToDec(blueHex[1])
#     return red, blue



for i in range(8):
    placementRow = []
    for j in range(8):
        placementRow.append([sideMarginWidth+((i+2)*blockHeight), (j+2)*blockHeight])
    placementGrid.append(placementRow)

while running:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

    pygame.time.Clock.tick(1) ##### fix this

    randNo = random.randint(0, 1000000)
    testHash = hashlib.sha256(bytes(randNo)).hexdigest()
    testHashList = []
    for n in range(len(testHash)):
        testHashList.append(testHash[n])


    surface.fill((255, 255, 255))
    for i in range(8):
        for j in range(8):
            if i == 0 and j == 0 and hexToDec(testHashList[0]) == 1:
                color = (0, 255, 0)
            elif i == 1 and j == 0 and hexToDec(testHashList[1]) == 2:
                color = (0, 255, 0)
            elif i == 2 and j == 0 and hexToDec(testHashList[2]) == 3:
                color = (0, 255, 0)
            elif i == 3 and j == 0 and hexToDec(testHashList[3]) == 4:
                color = (0, 255, 0)
            else:
                color = (hexToDec(testHashList[8*i+j])*16, 0, 0)
            pygame.draw.rect(surface, color, pygame.Rect(placementGrid[i][j][0], placementGrid[i][j][1], 60, 60))

    

    pygame.display.flip()