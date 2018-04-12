# Grouping by two fields


## Have database

|cnt| gr| id| name|
|---|---|---|--|
 4| g2| 5538e1a1-5209-4111-8831-46004ddb683b |bbbb4
 4|g2|2f37f7a2-3838-4b48-822c-63037f09fdaa|bbbb4
 4|g2|9ec3455f-d237-4e7f-b7c0-9fc3ab818c8b|bbbb4
 2|g2|eeabc2c6-410e-486f-aa48-0968320776c2|bbbb2
 4|g1|f085cc2e-7b11-43a8-b637-f49b6fcb07e5|ppp3
 1|g1|cd7648ab-0e55-4481-af15-2d9f5a33fedc|ppp
 2|g1|c085560d-b687-4960-8757-2ffaae95e4eb|ppp2
 5|g2|e6bff60d-990b-469e-b54c-4092a7d5a715|bbbb


## Program group with sum
```javascript

r.db("system").table("Temp")
  .group("gr","name")
  .ungroup()
  .map({"pp":  r.row("reduction")("name").nth(0), 
        "pp2": r.row("reduction")("cnt").sum()})

```


## Results

|p |pp|
|--|--|
|2 | ppp 
|1 | ppp2
|2 | ppp3
|4 | bbbb2
|2 | bbbb3
|12 | bbbb4
