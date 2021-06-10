main {
    (mul 0x02, @r2)
    ret;;
}

start {
    (psh 0x01)
    (pop @r2)
    (call $main)
}

