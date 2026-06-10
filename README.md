# Weather Station
## Learning
> [!Note]
> 1. Coverting string to integer
> - strconv.ParseInt(s, 10, 64) (ParseInt(s string, base int, bitSize ibnt)) **fastest**
> - strconv.Atoi(s) **still very fast**
> - fmt.Sscanf(s, "id:%5d", &i) *parsing custom strings holding a number* .exp:
> ```
> s := "id:00123"
> var i int
> if _, err := fmt.Sscanf(s, "id:%5d", &i); err == nil {
   > fmt.Println(i) // Outputs 123
> }
> ```

