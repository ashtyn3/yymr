# Wr (Written-version)

**syntax**

```

main {
    (mul 2, @2)
    ret;;
}

start {
    (psh 0x01)
    (pop @2)
    (call $main)
}
```
