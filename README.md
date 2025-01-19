# Collapse Numbers

A small module that helps to collapse a sorted slice of numbers into `{ {start, end}, {...} }` format.

## API

```go
collapse.Numbers([]int{1,2,3,6,8,9})
```

### arg1

**Type**: `[]int`   
Sorted slice of integers

### @return

**Type**: `[][]int`   

