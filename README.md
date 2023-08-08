## Usage

```
 $ txgen \
     --data <csv file path> \
     --template <message template file path> \
     --out <output directory>
```

## Input file example

### Data file

- The first column must be a mail address.
- After second column is free and variable.

```
1@example.com,John,Professor Stone,78%
2@example.com,Catherine,Professor Stone,90%
3@example.com,Mike,Professor Roberts,98%
```

### Template file

- This file will be content of generated files.
- You can use placeholder that will be replaced by related column in the data file.
- But, "$0" is prohibited. please start at "$1".

```
Hello, $1. I'm $2.
Your attendance rate is $3.
Best regards.
```

## Output file

- Any files will be generated. the file number equals a record number in a data file.
- Each files name will be "<mail address>.txt" 

### 1@example.txt

```
Hello, John. I'm Professor Stone.
Your attendance rate is 78%.
Best regards.
```

### 2@example.txt

```
Hello, Catherine. I'm Professor Stone.
Your attendance rate is 90%.
Best regards.
```

### 3@example.txt

```
Hello, Mike. I'm Professor Roberts.
Your attendance rate is 98%.
Best regards.
```
