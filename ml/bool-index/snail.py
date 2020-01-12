"""
Draw a green (color=[0, 0xFF, 0]) square on the snail image
- line width=5
- top left:     (217,  42)
- bottom right: (525, 275)
"""

import matplotlib.pyplot as plt

img = plt.imread('snail.jpg').copy()
plt.imshow(img)

tl_x, tl_y = 217, 42  # x = row, y = column
br_x, br_y = 525, 275
width = 5

color = [0, 0xFF, 0]  # green
img[tl_x:tl_x+width, tl_y:br_y] = color
img[br_x:br_x+width, tl_y:br_y] = color
img[tl_x:br_x, tl_y:tl_y+width] = color
img[tl_x:br_x, br_y:br_y+width] = color

plt.imshow(img)
plt.show()
