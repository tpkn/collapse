# Collapse

A small module that helps to collapse a sorted slice of numbers or dates into `{{start, end}, {...}}` format.

## Numbers

```go
collapse.Numbers([]int{1,2,3,6,8,9})
```

### arg1

**Type**: `[]int`   
Sorted slice of integers

### @return

**Type**: `[][]int`


## Dates


```go
collapse.DatesOnly([]string{"2025-01-01", "2025-01-02", "2025-01-03", "2025-01-04"})
```

### arg1

**Type**: `[]string`   
Sorted slice of strings

### @return

**Type**: `[][]string`   

