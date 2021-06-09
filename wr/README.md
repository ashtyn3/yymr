# Wr (Written-representation)

**syntax**

```
" Comments look like this...

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
