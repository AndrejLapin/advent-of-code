[.##.]
[....] (3) (1,3) (2) (2,3) (0,2) (0,1)

we need to find buttons that affect those indicators, which are 1 and 2
so we have
(3) (1,3) (2) (2,3) (0,2) (0,1)
    ^		^		^		^
we found 4 buttons that affect the right indicators

so only one button has no side effects which is (2)
buttons (1, 3), (0, 2) and (0, 1) have side effects

now how do we find combinations of the button presses?
to a human it is obvious, but to the algo, not so much,
how can an algorith determine that it's optimal to press (0, 2) and (0, 1)
since they have opposite side effects?

do we try some combinations?
okay, how about we start with a button that has no side effects and find another button that turns
on an indicator and helps our cause?

so we take (2) and then we only have left is (1, 3) and (0, 1)
if we take (1, 3) button we need to find a button that can disable the "3"
side effect
these buttons are (3), (1, 3), (2,3)
(1,3) should be eliminated because we just pressed it and it also turns off 1
(2,3) probably is okay, and we can turn on (2) after
(3) will immediately end the problem for us
so here we could try all these 3 options I guess? And the best one would be the button (3)
and we would collect the data that this took us 3 presses

then we somehow need to retry this problem with other buttons,
like (0,2) and (0,1)

maybe we need an algo that would calculate the best buttons for each press or something
and for the next one? like how close one button press gets us to the goal?

there is no way to determine whether the button press after the current one will be any better

what about button combinations? like the ones that would cancel eachother?

[.....] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4)
[...#.] - the only one we need is index 3
we have (0,2,3,4), (2, 3), (1,2,3,4)
best would obviously be (2, 3) first, but then how do we eliminate the side effect?
[..##.]
Can I even solve it myself?
What's the plan? I think to try a bunch of ideas, no? But then how do I know they are optimal
if we go with (0,2,3,4) we get
[#.###] what will (1,2,3,4) do?
[##...] then (0,1,2)
[..#..] then (2,3)
[...#.]
so it was a good plan to start from (2,3) huh? 4 moves

best one is:
(1,2,3,4) -> [.####]
(0,1,2) -> [#..##]
(0,4) -> [...#.]

we need some kind of calculation how easy it is to eliminate side effects?

so how would that work in the first one?
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1)
targets are [1,2]
side effect 3 is repeated 3 times, it is on button with target [1] and target [2]
side effect 0 is repeasted 2 times (so less likely), it is on button with target [1] and target [2]
they are even and produce even results here
but with weights calculations (2) somehow should be below all these 4 buttons
(1,3), (2,3), (0,2), (0,1) - i don't know how that would even work, because they introduce the side effects
but they cancel eachother if pressed, (2) does not have a non side effect pair, so it kinda has to be discarded

I think first we only should look at the buttons that contain targets, skip if they don't
after that count turning off the target as a side effect, and count toggling off bad ones as targets

[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4)
target is [3]
(0,2,3,4) has target [3] and side effects (0, 2, 4)
side effect 0 is repeasted 2 more times in (0,4) and (0,1,2) and has no target, but has other side effects
side effect 2 is repeated 3 more times in (2,3), (0,1,2) and (1,2,3,4), has target [3] 2 times and has other side effects
side effect 4 is repeated 2 more times in (0,4) and (1,2,3,4), has target [3] once, and has a bunch of other side effects
how many of the side effects have repeating button to toggle 0 and 2 have 1 repeating, 2 and 4 have 1 repeating but
this repeating has a [3] toggle, so maybe not so good, but (0, 1, 2) introduces only 1 new side effect
(0, 4) in this case would be a target button
so for each side effect we get the array or something
for 0 (+2, +1)
for 2 (0, +1, 0)
for 4 (+2, 0)
(so best case is +2?)
again
for 0 (0,4) and (0,1,2)
for 2 (2,3), (0,1,2) and (1,2,3,4)
for 4 (0,4) and (1,2,3,4)
for 0 (+2, +1)
for 2 (0, +1, 0)
for 4 (+2, 0)

(2,3) has target [3] and side effects are (2)
2 is repeated 3 more times in (0,2,3,4), (0,1,2) and (1,2,3,4), 2 contan target [3] (which would toggle off)
the ones that toggle off contain 4 unique side effects (maybe count the target as a side effect in this case?)
and one that has 2 side effeects, so maybe just, (3, 3, 2) ?
(-2, -1, -3) - pretty bad moves after, best case is -1

(0, 4) does not contain targets, so skipped

(0, 1, 2) no targets, skipped

(1, 2, 3, 4)
for 1 (0,1,2)
for 2 (0,2,3,4) (2,3) (0,1,2)
for 4 (0,2,3,4) (0,4)

so probably don't need to do it in a for side effect fashion, just calc the next bes move for the whole thing (I guess)

for 1 (+1)
for 2 (0, 0, +1)
for 4 (0, 0)
best case here is +1 only?
but there are also options that do not contain the target
or maybe it's possible here to try and find a multiple button combination to finish?
or we would just try all these buttons

but maybe toggling the target should have higher weight?

lets try number one again
[.##.] (3) (1,3) (2) (2,3) (0,2) (0,1)
[1, 2]
first check if anything finishes, second check second best case
maybe even save state if there are multiple best?
(1, 3) - (0) - next best cases (+1, +1, +2, 0, 0) - total +2
(2) - (+1) - next best case (-1, 0, -2, -2, 0) - total +1
(2, 3) - (0) - next best case (+1, +2, -1, -2, 0) - total +2
(0, 2) - (0) - next best case (-1, 0, -1, -2, +2) - total +2
(0, 1) - (0) - next best case (-1, -2, +1, 0, +2) - total +2

(1, 3) - check if anything finishes - (2, 3) finishes
// maybe don't try others here? because move count is 2
// and nothing will go less than 2, it does not finish with first, so we end here


[...#.] (0,2,3,4) (2,3) (0,4) (0,1,2) (1,2,3,4)
[3]
first check if anything finishes - nothing, how would this work best? Like if any of the buttons solves it?
But if we are calculating the best move, hmm
(0,2,3,4) - button press gives -2 (score) - next best (0, +2, +1, 0) - total 0
(2,3) - button press gives 0 (score) - next best (-2, -2, +1, -2) - total +1
(0,4) - button  - skip
(0,1,2) - skip
(1, 2, 3, 4) - press gives -2 (score) - (0, 0, 0, +1) - total -1 (but it's the best one)
I guess we would save this state, we have +2 and +1 options, and the +1 is the correct one
U guess the initial score would be healthy to track? Would it even make sense to save buttons which have
negative next best moves? Maybe?

if we already know the next best? (we calculated last time) do we just apply it? but there are multiple positive
with (0,2,3,4) selected targets are [0, 2, 4]
and the +2 was (0, 4) and +1 was (0, 1, 2)
save state or check how good the (0,4) next move will do? Or, do we do a check on the next best again?
confusing :D
lets try all again, calc next best for each

(2, 3) - (0) target would become [0, 3, 4] - next best (+2, +2, -1, 0) - with best case +2
(0, 4) - (+2) target would become [2] - next best (-2, 0, -1, -2) - with best case +2
(0,1,2) - (+1) target would become [1, 4] - next best (-2, -2, 0, 0) - with best case +1
(1, 2, 3, 4) - (0) - target would become [0, 1, 3] - next best (0, 0, 0, +1) - with best case +1


I wonder if we always have to save state, to return later
if we saved state for previous, but solved on next - discard
I guess we should save every time there are multiple positive options?

maybe calculating 2 forward is the way? Or maybe it does nothing
but I think we should try. But how is that the best move is actually the wors one then? Like it's options are the worst after that, maybe we are calculating best moves wrong then?

Maybe do some kind of paralel search?
We should try and implement paralel
After each branch we advance all branches by 1
But maybe that would be too much to save?
I shall try anyway
