# Singly Linked List Implementation - Go Programming Language
Simple implementation of the Singly Linked List using Go programming language.

## Usage
For running the Project (Without Building):
```
go run SimpleLinkedList.go
```
For building the Project:
```
go build SimpleLinkedList.go
```
For running the project after building:
```
./SimpleLinkedList
```

## Output
```

!!!!!!!!!!!!!!!!!!!!!!!!!!!
Can't Delete: List is Empty
!!!!!!!!!!!!!!!!!!!!!!!!!!!

--------------------------------
No. of Element in the List: 5
--------------------------------
Head: {10 0xc000014070}
Tail: {50 <nil>}
------------------
Node 1: {Data = 10, Link (Next) = 0xc000014070}
Node 2: {Data = 20, Link (Next) = 0xc000014080}
Node 3: {Data = 30, Link (Next) = 0xc000014090}
Node 4: {Data = 40, Link (Next) = 0xc0000140a0}
Node 5: {Data = 50, Link (Next) = 0x0}

!!!!!!!!!!!!!!!!!!!!!!!!!!!
Can't Delete: List is Empty
!!!!!!!!!!!!!!!!!!!!!!!!!!!

--------------------------------
No. of Element in the List: 0
--------------------------------

!!!!!!!!!!!!!
List is Empty
!!!!!!!!!!!!!

--------------------------------
No. of Element in the List: 8
--------------------------------
Head: {1 0xc0000140e0}
Tail: {8 <nil>}
------------------
Node 1: {Data = 1, Link (Next) = 0xc0000140e0}
Node 2: {Data = 2, Link (Next) = 0xc0000140f0}
Node 3: {Data = 3, Link (Next) = 0xc000014100}
Node 4: {Data = 4, Link (Next) = 0xc000014110}
Node 5: {Data = 5, Link (Next) = 0xc000014120}
Node 6: {Data = 6, Link (Next) = 0xc000014130}
Node 7: {Data = 7, Link (Next) = 0xc000014140}
Node 8: {Data = 8, Link (Next) = 0x0}
--------------------------------
No. of Element in the List: 3
--------------------------------
Head: {6 0xc000014130}
Tail: {8 <nil>}
------------------
Node 1: {Data = 6, Link (Next) = 0xc000014130}
Node 2: {Data = 7, Link (Next) = 0xc000014140}
Node 3: {Data = 8, Link (Next) = 0x0}
6 = Invalid Key!
--------------------------------
No. of Element in the List: 4
--------------------------------
Head: {15 0xc0000140d0}
Tail: {8 <nil>}
------------------
Node 1: {Data = 15, Link (Next) = 0xc0000140d0}
Node 2: {Data = 6, Link (Next) = 0xc000014130}
Node 3: {Data = 7, Link (Next) = 0xc000014140}
Node 4: {Data = 8, Link (Next) = 0x0}
--------------------------------
No. of Element in the List: 9
--------------------------------
Head: {50 0xc000014200}
Tail: {8 <nil>}
------------------
Node 1: {Data = 50, Link (Next) = 0xc000014200}
Node 2: {Data = 40, Link (Next) = 0xc0000141f0}
Node 3: {Data = 30, Link (Next) = 0xc0000141e0}
Node 4: {Data = 20, Link (Next) = 0xc0000141d0}
Node 5: {Data = 10, Link (Next) = 0xc0000141a0}
Node 6: {Data = 15, Link (Next) = 0xc0000140d0}
Node 7: {Data = 6, Link (Next) = 0xc000014130}
Node 8: {Data = 7, Link (Next) = 0xc000014140}
Node 9: {Data = 8, Link (Next) = 0x0}
--------------------------------
No. of Element in the List: 10
--------------------------------
Head: {50 0xc000014200}
Tail: {8 <nil>}
------------------
Node 1: {Data = 50, Link (Next) = 0xc000014200}
Node 2: {Data = 40, Link (Next) = 0xc0000141f0}
Node 3: {Data = 30, Link (Next) = 0xc0000141e0}
Node 4: {Data = 20, Link (Next) = 0xc0000141d0}
Node 5: {Data = 10, Link (Next) = 0xc0000141a0}
Node 6: {Data = 15, Link (Next) = 0xc0000140d0}
Node 7: {Data = 6, Link (Next) = 0xc000014130}
Node 8: {Data = 7, Link (Next) = 0xc000014240}
Node 9: {Data = 12, Link (Next) = 0xc000014140}
Node 10: {Data = 8, Link (Next) = 0x0}
--------------------------------
No. of Element in the List: 11
--------------------------------
Head: {50 0xc000014200}
Tail: {8 <nil>}
------------------
Node 1: {Data = 50, Link (Next) = 0xc000014200}
Node 2: {Data = 40, Link (Next) = 0xc0000141f0}
Node 3: {Data = 30, Link (Next) = 0xc0000141e0}
Node 4: {Data = 20, Link (Next) = 0xc0000141d0}
Node 5: {Data = 10, Link (Next) = 0xc000014270}
Node 6: {Data = 13, Link (Next) = 0xc0000141a0}
Node 7: {Data = 15, Link (Next) = 0xc0000140d0}
Node 8: {Data = 6, Link (Next) = 0xc000014130}
Node 9: {Data = 7, Link (Next) = 0xc000014240}
Node 10: {Data = 12, Link (Next) = 0xc000014140}
Node 11: {Data = 8, Link (Next) = 0x0}
--------------------------------
No. of Element in the List: 10
--------------------------------
Head: {50 0xc000014200}
Tail: {8 <nil>}
------------------
Node 1: {Data = 50, Link (Next) = 0xc000014200}
Node 2: {Data = 40, Link (Next) = 0xc0000141f0}
Node 3: {Data = 30, Link (Next) = 0xc0000141e0}
Node 4: {Data = 20, Link (Next) = 0xc0000141d0}
Node 5: {Data = 10, Link (Next) = 0xc0000141a0}
Node 6: {Data = 15, Link (Next) = 0xc0000140d0}
Node 7: {Data = 6, Link (Next) = 0xc000014130}
Node 8: {Data = 7, Link (Next) = 0xc000014240}
Node 9: {Data = 12, Link (Next) = 0xc000014140}
Node 10: {Data = 8, Link (Next) = 0x0}
--------------------------------
No. of Element in the List: 9
--------------------------------
Head: {50 0xc000014200}
Tail: {12 <nil>}
------------------
Node 1: {Data = 50, Link (Next) = 0xc000014200}
Node 2: {Data = 40, Link (Next) = 0xc0000141f0}
Node 3: {Data = 30, Link (Next) = 0xc0000141e0}
Node 4: {Data = 20, Link (Next) = 0xc0000141d0}
Node 5: {Data = 10, Link (Next) = 0xc0000141a0}
Node 6: {Data = 15, Link (Next) = 0xc0000140d0}
Node 7: {Data = 6, Link (Next) = 0xc000014130}
Node 8: {Data = 7, Link (Next) = 0xc000014240}
Node 9: {Data = 12, Link (Next) = 0x0}
--------------------------------
No. of Element in the List: 8
--------------------------------
Head: {40 0xc0000141f0}
Tail: {12 <nil>}
------------------
Node 1: {Data = 40, Link (Next) = 0xc0000141f0}
Node 2: {Data = 30, Link (Next) = 0xc0000141e0}
Node 3: {Data = 20, Link (Next) = 0xc0000141d0}
Node 4: {Data = 10, Link (Next) = 0xc0000141a0}
Node 5: {Data = 15, Link (Next) = 0xc0000140d0}
Node 6: {Data = 6, Link (Next) = 0xc000014130}
Node 7: {Data = 7, Link (Next) = 0xc000014240}
Node 8: {Data = 12, Link (Next) = 0x0}
0 = Invalid Key!
--------------------------------
No. of Element in the List: 8
--------------------------------
Head: {40 0xc0000141f0}
Tail: {12 <nil>}
------------------
Node 1: {Data = 40, Link (Next) = 0xc0000141f0}
Node 2: {Data = 30, Link (Next) = 0xc0000141e0}
Node 3: {Data = 20, Link (Next) = 0xc0000141d0}
Node 4: {Data = 10, Link (Next) = 0xc0000141a0}
Node 5: {Data = 15, Link (Next) = 0xc0000140d0}
Node 6: {Data = 6, Link (Next) = 0xc000014130}
Node 7: {Data = 7, Link (Next) = 0xc000014240}
Node 8: {Data = 12, Link (Next) = 0x0}
```

```
ⓘ Note: For the above output, requesting to refer the main() function of the code
```
