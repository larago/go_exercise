type ReadWrite interface {
    Read(b Buffer) bool 
    Write(b Buffer) bool 
}

type Lock interface {
    Lock()
    UnLock()
}

type File interface {
    ReadWirte 
    Lock
    Close() 
}
