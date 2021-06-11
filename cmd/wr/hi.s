 main {
    (mov 0x01, @r1)
    (mov 0x02, @r2)
	(push @r1)
	(push @r2)
	(pop @r1)
	(pop @r2)
}
