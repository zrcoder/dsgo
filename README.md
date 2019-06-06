# dsGo
data structures impletioned with Go
### `base` or `safe`?
we give 2 versions of each data structrue, thread safe or not safe, api is just the same. for example:
```
// this is a base set
github.com/zrcoder/dsGo/base/set

// a thread safe set
github.com/zrcoder/dsGo/safe/set
```
the base version usually performances better than the safe version <br>
and the safe version usually used for concurrent scenes
### our data structures
[queue](base/queue)
```
A queue gives you a FIFO or first-in firs-out order.
```
[stack](base/stack)
```
A stack gives you a LIFO or last-in first-out order.
```
[set](base/set)
```
A set can store unique values, without any particular order.
```
