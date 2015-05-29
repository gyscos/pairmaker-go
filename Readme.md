Pairmaker
=========

Pairmaker is a small tool designed to create schedules for pair programming sessions.

For `n` names, pairmaker generates `n-1` combinations, and add a 'reversed' version of each (for asymmetrical pairs: one leader, one follower).

It outputs one line per session, with the flat list of names. Take them two by two to make the pairs.

```
./pairmaker A,B,C,D
A,B,C,D
B,A,D,C
A,C,D,B
C,A,B,D
A,D,B,C
D,A,C,B
```

```
./pairmaker Alex,Greg,John,Marc,Neil,Saul
Alex,Greg,John,Saul,Marc,Neil
Greg,Alex,Saul,John,Neil,Marc
Alex,John,Marc,Greg,Neil,Saul
John,Alex,Greg,Marc,Saul,Neil
Alex,Marc,Neil,John,Saul,Greg
Marc,Alex,John,Neil,Greg,Saul
Alex,Neil,Saul,Marc,Greg,John
Neil,Alex,Marc,Saul,John,Greg
Alex,Saul,Greg,Neil,John,Marc
Saul,Alex,Neil,Greg,Marc,John
```
