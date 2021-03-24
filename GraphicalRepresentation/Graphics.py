import pygame
pygame.init()

screen = pygame.display.set_mode([600, 600])
running = True
dificulty = "1234"

def pixelsFromCoordinates(x, y): ##quick function to turn x and y grid coordinates to pixel coordinates
    return 30+60*x, 30+60*y

while running:
    for event in pygame.event.get():
        if event.type == pygame.QUIT:
            running = False

    screen.fill((255, 255, 255))
    color = (255, 0, 0)
    pygame.draw.rect(screen, color, pygame.Rect(30, 30, 60, 60))
    pygame.display.flip()