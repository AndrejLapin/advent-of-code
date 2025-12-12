
how to solve now?
we probably need to choose buttons that increment the most and then choose some
that will help reach end...

(0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4) {7,5,12,7,2}
4 can only be pressed twice
0 and 3 are pressed 7 times
1 is pressed 5 times
2 is 12 times

I think we need to work with subtraction (kinda again)
and equalize the indecies
the biggest one is 12 at 2
we have (0,2,3,4), (2,3), (0,1,2) and (1,2,3,4)
we kinda need the smalles one
that would be (2, 3) which is only -2

so I guess there are weights per index then?
But it should try to equalize first

find highest indecies maybe?
12 - 2
7 - 0
7 - 3
5 - 1
2 - 4

so best first is (2, 3), how?
because 2 and 0 would also have 1
and 2 and 3 is exists
!press!

11 - 2
7 - 0
6 - 3
5 - 1
2 - 4

so we get 11, and try to find it with 7, it goes only with 5
then we try to get 11 with 6, 6 is more than 5, so we proceed with (2, 3) again

10 - 2
7 - 0
5 - 1
5 - 3
2 - 4

10 with 7 and 5, could resolve a tie based on number count? Don't need, probably even better if higher?
(0,1,2)

9 - 2
6 - 0
5 - 3
4 - 1
2 - 4

9 with 6 and 4 vs 9 with 5 5 > 4
(2, 3)

8 - 2
6 - 0
4 - 1
4 - 3
2 - 4

8, 6, 4 (0, 1, 2)

7 - 2
5 - 0
4 - 3
3 - 1
2 - 4


rule is find the longest range with the **bigger** *smallest* number

7, 4
(2, 3)

6 - 2
5 - 0
3 - 1
3 - 3
2 - 4

(0, 1, 2)

5 - 2
4 - 0
3 - 3
2 - 1
2 - 4

(2, 3)

4 - 0
4 - 2
2 - 1
2 - 3
2 - 4

(0,2,3,4) and (1,2,3,4)
4, 4, 2, 2     2, 4, 2, 2 (this tie is based on sum or on the highest number count, sum is easier)
maybe we will get some mistakes in this part

so we choose based on sum
(0,2,3,4)

3 - 0
3 - 2
2 - 1
1 - 3
1 - 4

(0,2,3,4) and (1,2,3,4)
 3,3,1,1       2,3,1,1

(0,2,3,4)

2 - 0
2 - 2
2 - 1
0 - 3
0 - 4

so also can't choose with 0
only got (0, 1, 2) left twice,
what's the resutl?

(2, 3) - 5 times
(0, 1, 2) - 5 times
(0,2,3,4) - 2 times

so the initial formula did not work
the one that was looking at the smallestHighest number
then biggest length
then sum

maybe divergent paths should be created also because there were instances where buttons had identical scores

but lets try to solve the first one manually, should get 10

(3) (1,3) (2) (2,3) (0,2) (0,1) {3,5,4,7}

(7) (5,7) (4) (4,7) (3,4) (3,5)

press button (3)

{3,5,4,6}

6 - 3
5 - 1
4 - 2
3 - 0

(3)
so the algo is wrong!
sum based?

let's try again
{3,5,4,7}

7 - 3
4 - 2
5 - 1
3 - 0

sum based

(3) (1,3) (2) (2,3) (0,2) (0,1)

(7) (12)  (4) (11)  (7)   (8)

(1, 3)

6 - 3
4 - 2
4 - 1
3 - 0

(6) (10) (4)  (10)  (7)  (7)

(1, 3) should there be a divergent path here or something?

5 - 3
4 - 2
3 - 1
3 - 0

longest, then sum, or no? so far relied only on sum

(5)  (8)  (4)  (9)  (7)  (6)

(2, 3)

4 - 3
3 - 2
3 - 1
3 - 0

(4)  (7) (3) (7) (6) (6)

(1, 3)

3 - 3 
3 - 2
2 - 1
3 - 0

(3) (1,3) (2) (2,3) (0,2) (0,1)
(3) (5)   (3) (6)   (6)   (5)

(2, 3)

2 - 3
2 - 2
2 - 1
3 - 0

(3) (1,3) (2) (2,3) (0,2) (0,1)
(2)  (4)  (2)  (4)   (5)   (5)

(0,2)

2 - 3
1 - 2
2 - 1
2 - 0

(3) (1,3) (2) (2,3) (0,2) (0,1)
2     4    2    3      3    4

(1, 3)

1 - 3
1 - 2
1 - 1
2 - 0

(3) (1,3) (2) (2,3) (0,2) (0,1)
 1    2    1    2     3     3

(0, 2)

1 - 3
0 - 2
1 - 1
1 - 0

(3) (1,3) (2) (2,3) (0,2) (0,1)
 1    2    x    x     x     2

(1, 3)

0 - 3
0 - 2
0 - 1
1 - 0

uh oh! but kinda solvable with 1 I guess
this is a mess

I guess there should be diverging paths, for more accurate solutions

so probably this whole thing has to be done through math?
we just need to somehow determine how many times to subtract
but I mean without search states it could also kinda work? maybe search states are needed at the end or something?

but probably it should work linearly

I don't think it's possible to know which move is the best, is it?
or the priority has to shift?

length -> sum -> lowest highest?
maybe then search state is required?

For machine 0 might've been solved in 45 moves
with 29 mult
For machine 1 maybe in 77 wtih mult 20
For machine 2 maybe 46 mult 0
For machine 3 maybe 58 mult 30


Current counter state [0 1 0 0 0 15 0], machineId - 1 op count - 15817


