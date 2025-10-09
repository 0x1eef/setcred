package setcred

type setcred struct {
	sc_uid            uint32
	sc_ruid           uint32
	sc_svuid          uint32
	sc_gid            uint32
	sc_rgid           uint32
	sc_svgid          uint32
	sc_pad            uint32
	sc_supp_groups_nb uint32
	sc_supp_groups    uintptr
	sc_label          uintptr
}

func New(opts ...Option) (*setcred, uint) {
	creds := &setcred{
		sc_uid:            0,
		sc_ruid:           0,
		sc_svuid:          0,
		sc_gid:            0,
		sc_rgid:           0,
		sc_svgid:          0,
		sc_label:          0,
		sc_pad:            0,
		sc_supp_groups_nb: 0,
		sc_supp_groups:    0,
	}
	flags := uint(0)
	for _, set := range opts {
		set(creds, &flags)
	}
	return creds, flags
}
