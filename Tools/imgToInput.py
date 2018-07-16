from PIL import Image
import numpy as np
import imgParser

im = Image.open("sample.bmp")



img = np.array(im)
# print(img)


val = 0xd3
for idx, i in enumerate(img):
  for jdx, j in enumerate(i):
    r,g,b = 0xff, 0x00, 0x00

    if img[idx][jdx][1] > img[idx][jdx][2]:
      r,g,b = 255,255,0x80
    elif img[idx][jdx][0] <= val and img[idx][jdx][1] <= val and img[idx][jdx][2] <= val:
      r,g,b = 0x00, 0x00, 0xff
  
    else:
      r,g,b = 0,0,0
    img[idx][jdx] = r,g,b


parser = imgParser.Parser(img)
print(parser.getMap())
im2 = Image.fromarray(img)
im2.save("test.bmp")