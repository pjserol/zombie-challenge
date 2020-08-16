# Challenge Zombie Apocalypse

## Requirements

go installed

## Run the application

go run . data/input.json

## Result

Response generated in data/response.json

## Unit test

go test .

## Example of input of data (JSON format)

- dimensions of the area (N)
- the initial position of the zombie
- a list of positions of poor creatures
- a list of moves zombies will make U (Up), D (Down), L (Left), R (Right)

```json
{
    "dimensions": 4,
    "zombie": {
        "x": 2,
        "y": 1
    },
    "creature": [
        {
            "x": 0,
            "y": 1
        },
        {
            "x": 1,
            "y": 2
        },
        {
            "x": 3,
            "y": 1
        }
    ],
    "Moves": "DLUURR"
}
```

## Example of output (JSON format)

```json
{
    "Scores": 3,
    "zombie": [
        {
            "x": 3,
            "y": 0
        },
        {
            "x": 2,
            "y": 1
        },
        {
            "x": 1,
            "y": 0
        },
        {
            "x": 0,
            "y": 0
        }
    ]
}
```
