## Lists

The Redis database internally stores List as a linked list. This linked list has a head and tail. Whenever we insert a new element, we can insert it either at the head or tail. The "head" of the list is considered as the left-most element of the list and the "tail" is considered as right-most element of the list.

LPUSH command is used to insert a value at the head of the list. 
LRANGE command takes the start and end index as input, which is used to specify which elements from the list are displayed.If we are not aware of the elements count then we can provide -1 in the end index. The following command will show all the elements in the list.
The RPUSH command is used to insert a value at the tail of the list. 
The RPOP command is used to remove the element from the tail of the list.

The LLEN command is used to find the length of the list
The LINDEX command is used to find the element at a particular index in the list.
The LSET command is used to update the value at a given index.
The lpushx command adds an element to the head of the list if the list exists.
linsert command is useful if we need to insert an element before a particular element