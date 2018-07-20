









#0 RIGHT
#1 DOWN
#2 LEFT
#3 UP
dy = (0, 1, 0, -1)
dx = (1, 0, -1, 0)

def cpmap(m):
  cm = [] #copy
  for i in m:
    cm.append(i)
  return cm

def isValid(m, w, h, i, j):
  if 0 <= i < h:
    if 0 <= j <= w:
      if m[i][j] == '.':
        return True
  return False

def remove(m, i, j, target):
  queue = [(i, j)]
  while(len(queue)):
    i, j = queue[0]
    queue = [queue[1:]]
    
    
    
    
  

def game(width, height, board):
  print(width, height, board)
  m = [board[width*i:width*(i+1)] for i in range(height)]
  
  for i in m:
    print(i)
  

  

game(11, 7, ".cccdddcbd....cdcdbdd....b.ccbdb....b..........b..........b..................")