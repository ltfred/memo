## memo

`memo`是一个终端备忘录工具

## Example

### Add a new record

```shell
memo add # or `memo a`
```

### Show record

```shell
memo show # or `memo s`
```
#### Filter

You can use `-p` or `-s` filter record.

```shell
# filter priority, the value [low, high]
memo show -p low # memo s -p low

# filter status, the value [undo, doing]
memo show -s undo # memo s -s undo
```

#### Order

You can use `-o` order.

```shell
# order by `priority`
memo show -o p # memo s -o p

# order by `date`
memo show -o d # memo s -o d
```

Of course, you can also use them together.

```shell
memo s -p high -s undo -o d
```

### Modify

```shell
memo modify -u <uuid> # memo m -u <uuid>
```

### Delete

```shell
memo delete -u <uuid> # memo d -u <uuid>
```