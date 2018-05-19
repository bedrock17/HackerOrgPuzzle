#include <iostream>
#include <ctime>
#include <cstdio>
#include <cstdlib>

#define WIDTH  176
#define HEIGHT 176
#define BUFFER 176

using namespace std;

class Stack {
	int x, y;
	int direction;
	int count;
public:
	Stack() {
		x = 0;
		y = 0;
		direction = 0;
		count = 0;
	}
	void init(int a, int b, int c, int d) {
		x = a;
		y = b;
		direction = c;
		count = d;
	}
	int getX() {
		return x;
	}
	int getY() {
		return y;
	}
	int getDirection() {
		return direction;
	}
	int getCount() {
		return count;
	}
};
class Queue {
	int x, y;
public:
	Queue() {
		x = 0;
		y = 0;
	}
	void init(int x, int y) {
		this->x = x;
		this->y = y;
	}
	int getX() {
		return x;
	}
	int getY() {
		return y;
	}
};
int w;
int h;
char ans[BUFFER];

int dfs(int[][WIDTH], int, int);
int bfs(int[][WIDTH]);
int move(int[][WIDTH], int, int&, int&, int&);
int possibleTile(int[][WIDTH], int, int);
void convert(char[][WIDTH], int[][WIDTH]);

int main()
{
	while (1) {
		char tmp[HEIGHT][WIDTH] = { 0 };
		int tile[HEIGHT][WIDTH] = { 0 };
		clock_t start, end;

		printf("                                 ����������������������������������������������������\n");
		printf("                                 ��                                                ��\n");
		printf("                                 ��       http://www.hacker.org/coil Solution      ��\n");
		printf("                                 ��                                                ��\n");
		printf("                                 ��          Made by YunGoon in 2016-05-06         ��\n");
		printf("                                 ��                                                ��\n");
		printf("                                 ����������������������������������������������������\n\n");
		printf(">> Please enter flashvars element of html embed tag\n>> ");

		//�Է�
		scanf("x=%d&y=%d&board=", &w, &h);
		for (int i = 0; i < h; i++)
			for (int j = 0; j < w; j++)
				scanf("%c", &tmp[i][j]);

		start = clock();

		int ansX, ansY, exit = 1;
		for (int j = 0; exit && j < w; j++) {
			for (int i = 0; exit && i < h; i++) {
				if (tmp[i][j] == 'X') continue; //���� ������ ���� ��� �ٸ� ���� ����
				convert(tmp, tile); 			//tile�� ������ tmp�� ��ȯ�Ͽ� ����
				if (dfs(tile, j, i)) {
					ansX = j;
					ansY = i;
					exit = 0;
				}
				if (exit) printf("(%2d, %2d) passed\n", j, i);
			}
		}

		end = clock();

		printf("\a\n>> ANSWER (%lf seconds)\n", (end - start) / 1000.0);
		printf("http://www.hacker.org/coil/index.php?x=%d&y=%d&qpath=%s\n", ansX, ansY, ans);

		system("pause");
		system("cls");

		char c = getchar();
	}

	return 0;
}
int dfs(int t[][WIDTH], int a, int b)
{
	//(a,b)�� ���� ��ġ

	int k = 0; //���ÿ� �����ִ� �ڷ� ����
	int tile_count = 2;
	int dy[] = { 0,1,0,-1 };
	int dx[] = { 1,0,-1,0 };
	Stack s[WIDTH * HEIGHT];

	s[k++].init(a, b, -1, tile_count);

	while (k) {
		k--;
		int x = s[k].getX();			  //������ ������ ��ǥ
		int y = s[k].getY();			  //������ ������ ��ǥ
		int drc = s[k].getDirection();	  //1������ 2�Ʒ��� 3���� 4����
		int tct = s[k].getCount();		  //Ÿ�� �ѹ�

		t[y][x] = tct; //������ ����

		int m = move(t, drc, x, y, tct); //������ �������� �ΰ��� ���� ���� ������

		//���� ��ġ���� �̵��� Ÿ���� m(m=0�� 0, m=5�� 2)�� �ִ�
		int check = 0;
		int bfs_val = bfs(t);

		if (bfs_val == 2) {
			for (int i = 0; i < h; i++) {
				for (int j = 0; j < w; j++) {
					if (t[i][j] == 1) printf("�� ");
					else if (t[i][j] == 2) printf("�� ");
					else printf("%2d ", t[i][j] - 2);
				}
				printf("\n");
			}
			printf("\n");

			return 1;
		}
		if (bfs_val && m) { //Ÿ���� �̾��� �ְ� �� Ÿ���� �ִ�
			for (int i = 0; i < 4; i++) {
				int nx = x + dx[i];
				int ny = y + dy[i];

				if ((nx >= 0 && nx < w) && (ny >= 0 && ny < h) && !t[ny][nx]) {
					if (i == 0) s[k++].init(nx, ny, 1, tct + 1);		 //������
					else if (i == 1) s[k++].init(nx, ny, 2, tct + 1);	 //�Ʒ���
					else if (i == 2) s[k++].init(nx, ny, 3, tct + 1);	 //����
					else s[k++].init(nx, ny, 4, tct + 1);				 //����
				}
			}
		}
		else {
			for (int i = 0; i < h; i++) {
				for (int j = 0; j < w; j++) {
					if (k == 0) return 0;
					if (t[i][j] >= s[k - 1].getCount()) t[i][j] = 0;
				}
			}
		}
	}
}
int bfs(int t[][WIDTH])
{
	int catI = -1, catJ = -1, exit = 0;

	for (int i = 0; !exit && i < h; i++) {
		for (int j = 0; !exit && j < w; j++) {
			if (t[i][j] == 0) {
				catI = i;
				catJ = j;
				exit = 1;
			}
		}
	}

	if (catI + catJ < 0) return 2; //����!
	//(catJ,catI)�� ���� ��ġ

	int hd = 0, rr = 0;
	int dy[] = { 0,1,0,-1 };
	int dx[] = { 1,0,-1,0 };
	int c[HEIGHT][WIDTH];
	Queue q[WIDTH * HEIGHT];

	for (int i = 0; i < h; i++)
		for (int j = 0; j < w; j++)
			c[i][j] = t[i][j];

	q[rr++].init(catJ, catI);
	c[catI][catJ] = 1; //������ ����

	while (hd != rr) {
		int x = q[hd].getX();	 //������ ������ ��ǥ
		int y = q[hd].getY();	 //������ ������ ��ǥ
		hd++;

		for (int i = 0; i < 4; i++) {
			int nx = x + dx[i];
			int ny = y + dy[i];

			if ((nx >= 0 && nx < w) && (ny >= 0 && ny < h)) {
				if (!c[ny][nx]) {
					q[rr++].init(nx, ny);
					c[ny][nx] = 1; //������ ����
				}
			}
		}
	}

	for (int i = 0; i < h; i++) {
		for (int j = 0; j < w; j++) {
			if (c[i][j] == 0) return 0; //�Ұ���
		}
	}

	return 1; //����
}
int move(int t[][WIDTH], int d, int &x, int &y, int &tct)
{
	//�̵��� ���� �� ���ۿ� ���� ��� ��� �̵�
	//�� ������ ���� ������ ���� Ż��

	int count = 1;
	while (d >= 1 && d <= 4) {
		if (d == 1) {
			while (x < w - 1 && !t[y][x + 1]) t[y][++x] = tct;	//������
			if (count) ans[tct - 3] = 'R', ans[tct - 2] = '\0', count--;
		}
		else if (d == 3) {
			while (x > 0 && !t[y][x - 1]) t[y][--x] = tct;		//����
			if (count) ans[tct - 3] = 'L', ans[tct - 2] = '\0', count--;
		}
		else if (d == 4) {
			while (y > 0 && !t[y - 1][x]) t[--y][x] = tct;		//����
			if (count) ans[tct - 3] = 'U', ans[tct - 2] = '\0', count--;
		}
		else if (d == 2) {
			while (y < h - 1 && !t[y + 1][x]) t[++y][x] = tct;	//�Ʒ���
			if (count) ans[tct - 3] = 'D', ans[tct - 2] = '\0', count--;
		}

		d = possibleTile(t, x, y);
	}

	return d;
}
void convert(char tmp[][WIDTH], int tile[][WIDTH])
{
	for (int i = 0; i < h; i++) {
		for (int j = 0; j < w; j++) {
			if (tmp[i][j] == '.') tile[i][j] = 0;	//'.'�� 0����
			else tile[i][j] = 1;					//'X'�� 1��
		}
	}
}
int possibleTile(int t[][WIDTH], int x, int y)
{
	int dy[] = { 0,1,0,-1 };
	int dx[] = { 1,0,-1,0 };
	int cnt = 0;
	int drc;

	for (int i = 0; i < 4; i++) {
		int nx = x + dx[i];
		int ny = y + dy[i];

		if ((nx >= 0 && nx < w) && (ny >= 0 && ny < h)) { //���� �˻�
			if (!t[ny][nx]) {
				cnt++; //0�̸� ī����
				drc = i + 1;
			}
		}
	}

	if (cnt == 1) return drc;
	else if (cnt == 0) return 0;
	else if (cnt == 2) return 5;
}