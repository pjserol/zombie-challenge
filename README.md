# Challenge Zombie Apocalypse

The world consists of an n x n grid on which zombies and creatures live.

At the beginning of the program, a single zombie awakes and begins to move around the
grid.

If a zombie moves on the same location as a creature, the creature is
transformed into another zombie and zombies score one point.

The zombie continues moving and infecting creatures until has performed all its moves.
Once it has completed its movement, the first newly created zombie moves using the same
sequence as the original zombie.

Until all the newly created zombie moves using the same sequence as the original zombie.

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
    "creatures": [
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
    "moves": "DLUURR"
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
