# Circuit Breaker

**Circuit Breaker Pattern** is a design pattern in modern software development used to detect failures and encapsulates logic of preventing a failure to reoccur constantly.

### In software:

If a service is failing (e.g. slow or down), the circuit breaker:

1. Stops calling that service temporarily
2. Returns a fallback response instead
3. Gives the failing service time to recover

### How It Works (3 States)

1. Closed (Normal)
- Requests go through normally
- If failures exceed a threshold → switch to OPEN
2. Open (Broken)
- Requests are blocked immediately
- No calls to the failing service
- Returns fallback (default response, cached data, etc.)
3. Half-Open (Testing)
- Allow a few requests through
- If success → back to CLOSED
- If fail → back to OPEN

![alt text](<Screenshot 2026-03-28 150835.png>)

### What Circuit Breaker Solves

- Prevents system overload
- Avoids waiting for timeouts
- Improves system resilience
- Enables graceful degradation (fallback)



# Hystrix

- Hystrix is a library created by Netflix to implement the circuit breaker pattern.

![alt text](<Screenshot 2026-03-28 160134.png>)

# hystrix-dashboard docker image

- https://hub.docker.com/r/steeltoeoss/hystrix-dashboard