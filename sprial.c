#include <math.h>
#include <stdio.h>
#include <stdlib.h>


struct point {
    int x;
    int y;
};
typedef struct point point_t;

int sprialine_getpoint(int n, point_t *p);

int main(int argc, char ** argv) {
    int n;
    point_t p;

    printf("Input Number:");
    scanf("%d", &n);

    sprialine_getpoint(n, &p);

    printf ("return result: %d,%d\n", p.x, p.y);
    return 0;
}

void point_move_x_axis_left(point_t from, point_t *to, int len)
{
    to->y = from.y;
    to->x = from.x - len;
}

void point_move_x_axis_right(point_t from, point_t *to, int len) 
{
    to->y = from.y;
    to->x = from.x + len;

}

void point_move_y_axis_up(point_t from, point_t *to, int len)
{
    to->x = from.x;
    to->y = from.y + len;
}

void point_move_y_axis_down(point_t from, point_t *to, int len)
{
    to->x = from.x;
    to->y = from.y - len;
}

void get_odd_square_point(int a, point_t *p)
{
    int k = (a - 1)/2; //2k + 1 = a;
    p->x = -k;
    p->y = -p->x;
    printf("%d at location (%d, %d)\n", a*a, p->x, p->y);
}

void get_even_square_point(int a, point_t *p)
{
    int k = (a)/2; //2k=a;
    p->x = k;
    p->y = 1 - p->x;
    printf("%d at location (%d, %d)\n", a*a, p->x, p->y);
}

void get_odd_cross_point(point_t pa, point_t pb, point_t *p)
{
    p->x = pb.x;
    p->y = pa.y + 1;
    printf("cross point: (%d, %d)\n", p->x, p->y);
}

void get_even_cross_point(point_t pa, point_t pb, point_t *p) 
{
    p->x = pb.x;
    p->y = pa.y - 1;
    printf("cross point: (%d, %d)\n",  p->x, p->y);
}


int sprialine_getpoint(int n, point_t *p) 
{
    int a, b; // a^2 < n < (b)^2; b=a+1;
    int len, mid;
    point_t pa, pb, pm; //the a, b position and the cross point

    a = (int)floor(sqrt(n));
    if (a == (int)ceil(sqrt(n))) {
        if (a%2) get_odd_square_point(a, p);
        else get_even_square_point(a, p);
        return 0;
    }

    mid = a*a + a + 1; //mid the cross point of two line
    len = abs(n - mid);
    printf("%d is between %d^2 and %d^2, mid is %d, len is %d\n", n, a, a + 1, mid, len);
    
    if (a % 2 == 1) {
        get_odd_square_point(a, &pa);
        get_even_square_point(a + 1, &pb);
        get_odd_cross_point(pa, pb, &pm);
        if (n <= mid) {
            point_move_x_axis_left(pm, p, len);
        } else {
            point_move_y_axis_down(pm, p, len);
        }
    } else {
        get_even_square_point(a, &pa);
        get_odd_square_point(a + 1, &pb);
        get_even_cross_point(pa, pb, &pm);
        if (n <= mid) {
            point_move_x_axis_right(pm, p, len);
        } else {
            point_move_y_axis_up(pm, p, len);
        }
    }
    
    return 0;
}
