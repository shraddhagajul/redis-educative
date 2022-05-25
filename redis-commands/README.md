Redis is a key-value store, but that doesn't mean that it stores only string keys and string values. Redis supports different types of data structures as values. The key in Redis is a binary-safe String, with a max size of 512 MB, but you should always consider creating shorter keys.

A binary-safe string is a string that can contain any kind of data, e.g., a JPEG image or a serialized Java object


## String
In Redis, we have the advantage of storing both strings and collections of strings as values. A string value cannot exceed 512 MB of text or binary data. However, it can store any type of data, like text, integers, floats, videos, images, and audio files.

Memcached is also an open-source distributed memory caching system. Like Redis, it also stores data in key-value pairs, but it only supports only String type data.

Redis String can be used to store session IDs, static HTML pages, configuration XML, JSON, etc. It can also be used to work as a counter if integers are stored.

## List
If we need to store a collection of strings in Redis, then we can use the List type. If we use List in Redis, the elements are stored in a linked list. The benefit of this is the quick insertion and removal of the element from the head. If we need to insert an element in a List with 500 records, then it will take about the same amount of time as adding the element in a list of 50,000 records.

The downside is that if we need to access an element, the entire list is scanned, and it becomes a time-consuming operation. Since the List uses a linked list, the elements are sorted on the basis of the insertion order.

<!-- The list should be stored in those cases where the order of insertion matters and where the write speed matters as compared to the read speed. One such case is storing logs. -->

## Set
The Set value type is similar to List. The only difference is that the set doesn't allow duplicates. The elements are not sorted in any order.

<!-- Set offers constant time performance for adding and removing operations. We can use set to store data where uniqueness matters, e.g., storing the number of unique visitors on our website. -->

## Sorted set
If we need our elements to be sorted, we can use Sorted Set as the value type. Each element in the sorted set is associated with a number, called a score. The elements are stored in the Set based on their score. Let's say we have a key called fruits. We need to store apple and banana as the value. Let's say the score of apple is 10, and the score of banana is 15. As we can see, scoreapple < scorebanana
scoreapple<scorebanana
, so the order will be apple, followed by banana.

If the score of two elements is the same, then we check which String is lexicographically bigger. The two strings cannot be the same, as this is a Set.

Lexicographic order is dictionary order, except that all the uppercase letters precede all the lowercase letters.

## Hash
The hash value type is a field-value pair. Let's say we need to store the information about the marks scored by students. In this case, the subject can be the key. The value can be a field-value pair, where the field is the student name, and the value is the marks obtained.