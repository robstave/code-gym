# Answer: Product inventory

**Difficulty: Easy**

A map lookup for a missing key returns the value type's zero value. Checking
whether that value is nonzero therefore cannot distinguish a missing product
from a product with zero items. The sample prints `Found: false`; it should
print `Found: true`.

Use the two-result map lookup:

```go
_, found := inventory[product]
return found
```
