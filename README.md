# Getting Started

```
$make dep
$make run
```

and access to graghql playground : `http://localhost:8082/`

# Sample query

```
query{
  greet(persons:[
    {
      name: "Pera"
    },
    {
      name: "Nikov"
    },
  ]){
    message
  }
}
```
