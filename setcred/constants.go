package setcred

const SYSCALL = int(591)
const SETUID = uint(1) << 0
const SETRUID = uint(1) << 1
const SETSVUID = uint(1) << 2
const SETGID = uint(1) << 3
const SETRGID = uint(1) << 4
const SETSVGID = uint(1) << 5
const SETSUPPGROUPS = uint(1) << 6
const SETMACLABEL = uint(1) << 7
