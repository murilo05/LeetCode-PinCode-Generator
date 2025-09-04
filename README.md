# LeetCode Challenge - PinCode Generator

This repository contains the implementation of a technical challenge that came up for me in an interview some time ago.  
Unfortunately, I did not advance in the process, but I decided to improve the solution on my own as a **study resource** and make it available for other developers to learn and practice.

---

## 📌 The Challenge

The problem consisted of:

- Generating a list of **4-digit pinCodes**.  
- The user can choose **how many pinCodes** they want to generate.  
- The list can be represented as either `int` or `string`.  
  - In this solution, I chose to convert them to **string**.  
- It is necessary to check if there are **duplicate pinCodes** (no duplicates allowed).  
- Create **unit tests** to ensure the implementation is valid.  

---

## 🚀 Solution

The solution was implemented in **Go (Golang)**, with the following key points:

1. **Generate a 4-digit pinCode**  
   - Each pinCode consists of random numbers from `0` to `9`.

2. **Convert to string**  
   - Example: `[0 9 8 7]` → `"0987"`.  
   - This makes it easier to check for duplicates using a `map[string]bool`.

3. **Generate the full list without duplicates**  
   - The user specifies the desired quantity.  
   - The program ensures there are no repeated values.  

4. **Unit Tests**  
   - Validate the number of pinCodes generated.  
   - Ensure there are no duplicates.  
   - Check that all pinCodes have **exactly 4 digits**.  

---

## 🧑‍💻 Usage Example

```go
//go run main.go
func main() {
    input := 10 //Number of PIN codes that will be generated
    pinCodes := generatePinCodeList(input)
    fmt.Println(pinCodes)
}
