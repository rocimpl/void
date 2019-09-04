# COOKBOOK

## Database concent

```
CREATE TABLE logs
(
    timestamp   DateTime,
    nanoseconds Int16,
    label       Array(String),
    source      String,
    hostname    String,
    level       FixedString(1),
    message     String,
    addition    String
)
    ENGINE = MergeTree
        PARTITION BY (toYYYYMM(timestamp), source, level)
        ORDER BY (timestamp, nanoseconds)
        SETTINGS index_granularity = 16384;
```