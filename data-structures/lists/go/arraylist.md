ArrayLists have a couple of interesting properties. First, the ArrayList implements the List interface, which means each element in the list has a corresponding index. 
Lists allow for duplicate elements and implement the following methods
.get(int index)
.add(E element)
.size()
.remove(int index)
.set(int index, E element)
.clear()

There are other methods we may wish to implement (such as addAll) but these are the bare minimum as I'm thinking about it right at this moment.

Secondly, ArrayLists are growable, meaning their memory allocation increases as necessary alongside the size of the data structure.