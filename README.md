# dsGo
data structues impletioned with Go
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
