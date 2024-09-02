# 🏔️ Aurelius

Aurelius is an embeddable log-structured storage engine, built for educational purposes. Its usage in production environment is not recommended without testing.

## Architecture and Implementation Details

Aurelius is much simpler compared to other alternatives and lacks many features. The main goal of the project is to understand how log-structured storage engines work and to build a simple version from scratch. You can find the core components and implementation details below.

### Memtable

A memtable is an in-memory data structure that holds keys and values before they are written to disk. Aurelius uses a skiplist-based memtable implementation.

The memtable is configured to have a specific size in bytes. If this size threshold is exceeded after a write operation, the memtable becomes immutable and is swapped with a new one. Eventually, the immutable memtable is flushed to disk in a table format and then destroyed.

The memtable implementation does not support key deletion; instead, a special token is assigned as the value.