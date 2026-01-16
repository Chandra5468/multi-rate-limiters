# Trying different kind of rate limiter implementation

### 1. Token bucket limit
    ---- How it works
        Tokens are added to a bucket at a fixed rate.
        A request consumes a token; if none are available, its delayed or rejected
    
    ---- Best for
        Allowing controlled bursts of traffic (e.g., a user makes several rapid requests)
        while maintaining avg rate.
### 2. Leaky bucket
    ---- How it works
        Requests fill a bucket, but leak out (are processed) at a constant, fixed rate, smooting out traffic.
    
    ---- Best for
        Enforcing a strict, steady output rate, ideal for traffic shaping

### 3. Fixed window
    ---- How it works
        Divides time into fixed intervals (eg 1 minute) and counts requests
        within that window. If the limit is hit, subsequent requests in that window are
        blocked.
    
    ---- Best for
        Simplicity, but can allow double the requests at window boundaries (eg end of one window/start of next)

### 4. Sliding window
    ---- How it works
        Tracks requests over a moving time window. The log keeps every timestamp, while the
        counter uses aggregated counts for sub-windows for better performance.
    ---- Best for
        Fairness and accuracy by adapting to recent traffic, solving the "double burst" issue of fixed window, through memory-intensive for the log version.


### Whichone to use when

1. For API Gateways : ***Token Bucket*** is very popular for general API protection.

2. For Network Traffic Smoothing : ***Leaky Bucket*** is excellent for controlling packet flow.

3. For fair usage : ***Sliding window counter*** offers a good balance of fairness and efficiency for user-specific limits.