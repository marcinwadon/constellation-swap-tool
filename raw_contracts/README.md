 ### Requirements:
 - abigen
 
 ### Compile to go lang
 
 1. Create .bin using smart contract binary code and .abi using json abi.
 2. Compile using following commands 
 
 `abigen --bin=Token.bin --abi=Token.abi --pkg=token --out=Token.go`

 `abigen --bin=Swap.bin --abi=Swap.abi --pkg=token --out=Swap.go`
