

#blue color code
possible = 0x00
imposiible = 0x80
div = 0xff

class Parser:
  #초기화시 첫 좌표로 적절한 상태인지 판단 
  def __init__(self, img, i=0, j=0):
    self.img = img
    self.i = i
    self.j = j
    self.map = []

    self.map.append([])

    v = self.getValue(self.img[i][j][2])
    if v >= 0:
      self.map[0].append(v)
      self.parseRight(0)
      self.parseRight(0)
      self.parseRight(0)
      self.parseRight(0)
      self.parseRight(0)
      self.parseRight(0)
    else:
      print("ERROR!!!! check img state")

  def getValue(self, bludCode):
    if bludCode == possible:
      return 1
    elif bludCode == imposiible:
      return 0
    elif bludCode == div:
      return -1 #나누는 경계영역
  
  def parseRight(self, i):
    #경계지점 다음에 블록이 시작되는 지점을 찾고 저장
    #0: 초기상태 1: 경계지점
    state = 0
    
    while(state<2):
      if state == 0 and self.img[i][self.j][2] == imposiible:
        print(self.j)
        state = 1
        
      elif state == 1 and self.img[i][self.j][2] == possible:
        state = 2
        v = self.getValue(self.img[i][self.j][2])
        print("DEBUG ", v)
        self.map[i].append(v)
        return
      else:
        self.j += 1
    # def Parse(self):

  
  def getMap(self):
    return self.map