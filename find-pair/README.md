#Challenge 2

##Dependencies
Golang 1.7 or later.<br />
**Note:** This program uses only Golang standard library

##Setup
1. Clone the repo.
2. Setup go workspace appropriately (if not already setup)
3. `cd $GOPATH/github.com/paxos/find-pair
4. `go build`
5. `./find-pair <file> <target>`<br/>
**NOTE:** To run the bonus for this challenge, use `./find-pair <file> <target> -bonus`

##Success Examples
```bash
cat utxos1.txt
   abcdef 5
   48a92b 7
   e478ab 10
   13474a 14
   a84739 20
./find-pair utxos1.txt 25
   a84739   20 ,  abcdef   5

cat utxos2.txt
   abcdef 1
   e478ab 2
   74738a 2
   a84739 22

./find-pair utxos2.txt 25
   Not possible

cat utxos3.txt
   147bce 2
   abcdef 6
   e478ab 20
   a84739 24

./find-pair utxos3.txt 25
   a84739   24 ,  147bce   2
```

##Success Examples Bonus
```bash
cat utxos4.txt
   abcdef 1
   e478ab 1
   74738a 1
   oooooo 3
   a84739 22

./find-pair utxos4.txt 3 -bonus
   abcdef 1 e478ab 1 74738a 1
./find-pair utxos4.txt 3 -bonus
   oooooo 3
```

##Errored Examples
```bash
cat utxos5.txt
    147bce 2
    abcdef 6
    e478ab
    a84739 24

./find-pair utxos5.txt 3
    File contains bad input.

./find-pair utxos5.txt 3 -bonus
    File contains bad input.

 cat utxos6.txt

./find-pair utxos6.txt 5
    File contains bad input.

./find-pair utxos6.txt 5 -bonus
    File contains bad input.
 ```



