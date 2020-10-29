// go run linux/mksysnum.go -Wall -Werror -static -I/tmp/include -fsigned-char /tmp/include/asm/unistd.h
// Code generated by the command above; see README.md. DO NOT EDIT.

// +build s390x,linux

package unix

const (
	SYS_EXIT                   = 1
	SYS_FORK                   = 2
	SYS_READ                   = 3
	SYS_WRITE                  = 4
	SYS_OPEN                   = 5
	SYS_CLOSE                  = 6
	SYS_RESTART_SYSCALL        = 7
	SYS_CREAT                  = 8
	SYS_LINK                   = 9
	SYS_UNLINK                 = 10
	SYS_EXECVE                 = 11
	SYS_CHDIR                  = 12
	SYS_MKNOD                  = 14
	SYS_CHMOD                  = 15
	SYS_LSEEK                  = 19
	SYS_GETPID                 = 20
	SYS_MOUNT                  = 21
	SYS_UMOUNT                 = 22
	SYS_PTRACE                 = 26
	SYS_ALARM                  = 27
	SYS_PAUSE                  = 29
	SYS_UTIME                  = 30
	SYS_ACCESS                 = 33
	SYS_NICE                   = 34
	SYS_SYNC                   = 36
	SYS_KILL                   = 37
	SYS_RENAME                 = 38
	SYS_MKDIR                  = 39
	SYS_RMDIR                  = 40
	SYS_DUP                    = 41
	SYS_PIPE                   = 42
	SYS_TIMES                  = 43
	SYS_BRK                    = 45
	SYS_SIGNAL                 = 48
	SYS_ACCT                   = 51
	SYS_UMOUNT2                = 52
	SYS_IOCTL                  = 54
	SYS_FCNTL                  = 55
	SYS_SETPGID                = 57
	SYS_UMASK                  = 60
	SYS_CHROOT                 = 61
	SYS_USTAT                  = 62
	SYS_DUP2                   = 63
	SYS_GETPPID                = 64
	SYS_GETPGRP                = 65
	SYS_SETSID                 = 66
	SYS_SIGACTION              = 67
	SYS_SIGSUSPEND             = 72
	SYS_SIGPENDING             = 73
	SYS_SPCHOSTNAME            = 74
	SYS_SETRLIMIT              = 75
	SYS_GETRUSAGE              = 77
	SYS_GETTIMEOFDAY           = 78
	SYS_SETTIMEOFDAY           = 79
	SYS_SYMLINK                = 83
	SYS_READLINK               = 85
	SYS_USELIB                 = 86
	SYS_SWAPON                 = 87
	SYS_REBOOT                 = 88
	SYS_READDIR                = 89
	SYS_MMAP                   = 90
	SYS_MUNMAP                 = 91
	SYS_TRUNCATE               = 92
	SYS_FTRUNCATE              = 93
	SYS_FCHMOD                 = 94
	SYS_GETPRIORITY            = 96
	SYS_SETPRIORITY            = 97
	SYS_STATFS                 = 99
	SYS_FSTATFS                = 100
	SYS_SOCKETCALL             = 102
	SYS_SYSLOG                 = 103
	SYS_SETITIMER              = 104
	SYS_GETITIMER              = 105
	SYS_STAT                   = 106
	SYS_LSTAT                  = 107
	SYS_FSTAT                  = 108
	SYS_LOOKUP_DCOOKIE         = 110
	SYS_VHANGUP                = 111
	SYS_IDLE                   = 112
	SYS_WAIT4                  = 114
	SYS_SWAPOFF                = 115
	SYS_SYSINFO                = 116
	SYS_IPC                    = 117
	SYS_FSYNC                  = 118
	SYS_SIGRETURN              = 119
	SYS_CLONE                  = 120
	SYS_SETDOMAINNAME          = 121
	SYS_UNAME                  = 122
	SYS_ADJTIMEX               = 124
	SYS_MPROTECT               = 125
	SYS_SIGPROCMASK            = 126
	SYS_CREATE_MODULE          = 127
	SYS_INIT_MODULE            = 128
	SYS_DELETE_MODULE          = 129
	SYS_GET_KERNEL_SYMS        = 130
	SYS_QUOTACTL               = 131
	SYS_GETPGID                = 132
	SYS_FCHDIR                 = 133
	SYS_BDFLUSH                = 134
	SYS_SYSFS                  = 135
	SYS_PERSONALITY            = 136
	SYS_AFS_SYSCALL            = 137
	SYS_GETDENTS               = 141
	SYS_SELECT                 = 142
	SYS_FLOCK                  = 143
	SYS_MSYNC                  = 144
	SYS_READV                  = 145
	SYS_WRITEV                 = 146
	SYS_GETSID                 = 147
	SYS_FDATASYNC              = 148
	SYS__SYSCTL                = 149
	SYS_MLOCK                  = 150
	SYS_MUNLOCK                = 151
	SYS_MLOCKALL               = 152
	SYS_MUNLOCKALL             = 153
	SYS_SCHED_SETPARAM         = 154
	SYS_SCHED_GETPARAM         = 155
	SYS_SCHED_SETSCHEDULER     = 156
	SYS_SCHED_GETSCHEDULER     = 157
	SYS_SCHED_YIELD            = 158
	SYS_SCHED_GET_PRIORITY_MAX = 159
	SYS_SCHED_GET_PRIORITY_MIN = 160
	SYS_SCHED_RR_GET_INTERVAL  = 161
	SYS_NANOSLEEP              = 162
	SYS_MREMAP                 = 163
	SYS_QUERY_MODULE           = 167
	SYS_POLL                   = 168
	SYS_NFSSERVCTL             = 169
	SYS_PRCTL                  = 172
	SYS_RT_SIGRETURN           = 173
	SYS_RT_SIGACTION           = 174
	SYS_RT_SIGPROCMASK         = 175
	SYS_RT_SIGPENDING          = 176
	SYS_RT_SIGTIMEDWAIT        = 177
	SYS_RT_SIGQUEUEINFO        = 178
	SYS_RT_SIGSUSPEND          = 179
	SYS_PREAD64                = 180
	SYS_PWRITE64               = 181
	SYS_GETCWD                 = 183
	SYS_CAPGET                 = 184
	SYS_CAPSET                 = 185
	SYS_SIGALTSTACK            = 186
	SYS_SENDFILE               = 187
	SYS_GETPMSG                = 188
	SYS_PUTPMSG                = 189
	SYS_VFORK                  = 190
	SYS_GETRLIMIT              = 191
	SYS_LCHOWN                 = 198
	SYS_GETUID                 = 199
	SYS_GETGID                 = 200
	SYS_GETEUID                = 201
	SYS_GETEGID                = 202
	SYS_SETREUID               = 203
	SYS_SETREGID               = 204
	SYS_GETGROUPS              = 205
	SYS_SETGROUPS              = 206
	SYS_FCHOWN                 = 207
	SYS_SETRESUID              = 208
	SYS_GETRESUID              = 209
	SYS_SETRESGID              = 210
	SYS_GETRESGID              = 211
	SYS_CHOWN                  = 212
	SYS_SETUID                 = 213
	SYS_SETGID                 = 214
	SYS_SETFSUID               = 215
	SYS_SETFSGID               = 216
	SYS_PIVOT_ROOT             = 217
	SYS_MINCORE                = 218
	SYS_MADVISE                = 219
	SYS_GETDENTS64             = 220
	SYS_READAHEAD              = 222
	SYS_SETXATTR               = 224
	SYS_LSETXATTR              = 225
	SYS_FSETXATTR              = 226
	SYS_GETXATTR               = 227
	SYS_LGETXATTR              = 228
	SYS_FGETXATTR              = 229
	SYS_LISTXATTR              = 230
	SYS_LLISTXATTR             = 231
	SYS_FLISTXATTR             = 232
	SYS_REMOVEXATTR            = 233
	SYS_LREMOVEXATTR           = 234
	SYS_FREMOVEXATTR           = 235
	SYS_GETTID                 = 236
	SYS_TKILL                  = 237
	SYS_FUTEX                  = 238
	SYS_SCHED_SETAFFINITY      = 239
	SYS_SCHED_GETAFFINITY      = 240
	SYS_TGKILL                 = 241
	SYS_IO_SETUP               = 243
	SYS_IO_DESTROY             = 244
	SYS_IO_GETEVENTS           = 245
	SYS_IO_SUBMIT              = 246
	SYS_IO_CANCEL              = 247
	SYS_EXIT_GROUP             = 248
	SYS_EPOLL_CREATE           = 249
	SYS_EPOLL_CTL              = 250
	SYS_EPOLL_WAIT             = 251
	SYS_SET_TID_ADDRESS        = 252
	SYS_FADVISE64              = 253
	SYS_TIMER_CREATE           = 254
	SYS_TIMER_SETTIME          = 255
	SYS_TIMER_GETTIME          = 256
	SYS_TIMER_GETOVERRUN       = 257
	SYS_TIMER_DELETE           = 258
	SYS_CLOCK_SETTIME          = 259
	SYS_CLOCK_GETTIME          = 260
	SYS_CLOCK_GETRES           = 261
	SYS_CLOCK_NANOSLEEP        = 262
	SYS_STATFS64               = 265
	SYS_FSTATFS64              = 266
	SYS_REMAP_FILE_PAGES       = 267
	SYS_MBIND                  = 268
	SYS_GET_MEMPOLICY          = 269
	SYS_SET_MEMPOLICY          = 270
	SYS_MQ_OPEN                = 271
	SYS_MQ_UNLINK              = 272
	SYS_MQ_TIMEDSEND           = 273
	SYS_MQ_TIMEDRECEIVE        = 274
	SYS_MQ_NOTIFY              = 275
	SYS_MQ_GETSETATTR          = 276
	SYS_KEXEC_LOAD             = 277
	SYS_ADD_KEY                = 278
	SYS_REQUEST_KEY            = 279
	SYS_KEYCTL                 = 280
	SYS_WAITID                 = 281
	SYS_IOPRIO_SET             = 282
	SYS_IOPRIO_GET             = 283
	SYS_INOTIFY_INIT           = 284
	SYS_INOTIFY_ADD_WATCH      = 285
	SYS_INOTIFY_RM_WATCH       = 286
	SYS_MIGRATE_PAGES          = 287
	SYS_OPENAT                 = 288
	SYS_MKDIRAT                = 289
	SYS_MKNODAT                = 290
	SYS_FCHOWNAT               = 291
	SYS_FUTIMESAT              = 292
	SYS_NEWFSTATAT             = 293
	SYS_UNLINKAT               = 294
	SYS_RENAMEAT               = 295
	SYS_LINKAT                 = 296
	SYS_SYMLINKAT              = 297
	SYS_READLINKAT             = 298
	SYS_FCHMODAT               = 299
	SYS_FACCESSAT              = 300
	SYS_PSELECT6               = 301
	SYS_PPOLL                  = 302
	SYS_UNSHARE                = 303
	SYS_SET_ROBUST_LIST        = 304
	SYS_GET_ROBUST_LIST        = 305
	SYS_SPLICE                 = 306
	SYS_SYNC_FILE_RANGE        = 307
	SYS_TEE                    = 308
	SYS_VMSPLICE               = 309
	SYS_MOVE_PAGES             = 310
	SYS_GETCPU                 = 311
	SYS_EPOLL_PWAIT            = 312
	SYS_UTIMES                 = 313
	SYS_FALLOCATE              = 314
	SYS_UTIMENSAT              = 315
	SYS_SIGNALFD               = 316
	SYS_TIMERFD                = 317
	SYS_EVENTFD                = 318
	SYS_TIMERFD_CREATE         = 319
	SYS_TIMERFD_SETTIME        = 320
	SYS_TIMERFD_GETTIME        = 321
	SYS_SIGNALFD4              = 322
	SYS_EVENTFD2               = 323
	SYS_INOTIFY_INIT1          = 324
	SYS_PIPE2                  = 325
	SYS_DUP3                   = 326
	SYS_EPOLL_CREATE1          = 327
	SYS_PREADV                 = 328
	SYS_PWRITEV                = 329
	SYS_RT_TGSIGQUEUEINFO      = 330
	SYS_PERF_EVENT_OPEN        = 331
	SYS_FANOTIFY_INIT          = 332
	SYS_FANOTIFY_MARK          = 333
	SYS_PRLIMIT64              = 334
	SYS_NAME_TO_HANDLE_AT      = 335
	SYS_OPEN_BY_HANDLE_AT      = 336
	SYS_CLOCK_ADJTIME          = 337
	SYS_SYNCFS                 = 338
	SYS_SETNS                  = 339
	SYS_PROCESS_VM_READV       = 340
	SYS_PROCESS_VM_WRITEV      = 341
	SYS_S390_RUNTIME_INSTR     = 342
	SYS_KCMP                   = 343
	SYS_FINIT_MODULE           = 344
	SYS_SCHED_SETATTR          = 345
	SYS_SCHED_GETATTR          = 346
	SYS_RENAMEAT2              = 347
	SYS_SECCOMP                = 348
	SYS_GETRANDOM              = 349
	SYS_MEMFD_CREATE           = 350
	SYS_BPF                    = 351
	SYS_S390_PCI_MMIO_WRITE    = 352
	SYS_S390_PCI_MMIO_READ     = 353
	SYS_EXECVEAT               = 354
	SYS_USERFAULTFD            = 355
	SYS_MEMBARRIER             = 356
	SYS_RECVMMSG               = 357
	SYS_SENDMMSG               = 358
	SYS_SOCKET                 = 359
	SYS_SOCKETPAIR             = 360
	SYS_BIND                   = 361
	SYS_CONNECT                = 362
	SYS_LISTEN                 = 363
	SYS_ACCEPT4                = 364
	SYS_GETSOCKOPT             = 365
	SYS_SETSOCKOPT             = 366
	SYS_GETSOCKNAME            = 367
	SYS_GETPEERNAME            = 368
	SYS_SENDTO                 = 369
	SYS_SENDMSG                = 370
	SYS_RECVFROM               = 371
	SYS_RECVMSG                = 372
	SYS_SHUTDOWN               = 373
	SYS_MLOCK2                 = 374
	SYS_COPY_FILE_RANGE        = 375
	SYS_PREADV2                = 376
	SYS_PWRITEV2               = 377
	SYS_S390_GUARDED_STORAGE   = 378
	SYS_STATX                  = 379
	SYS_S390_STHYI             = 380
	SYS_KEXEC_FILE_LOAD        = 381
	SYS_IO_PGETEVENTS          = 382
	SYS_RSEQ                   = 383
)
