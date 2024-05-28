
functions
- Add numbers
- Create string of numbers added
- Add to dictionary
- Remove numbers from numbers to be added array


Structure

{
    NumberOfDigits (int as string): {
        Value (int as string): [
            combinations (ints as strings)
        ]
    }
}


{
    2 digit combinations: {
        with the value of 2: [ array possible combinations ],
        with the value of 3: [ array possible combinations ]
    },
        3 digit combinations: {
        with the value of 6: [ array possible combinations ],
        with the value of 7: [ array possible combinations ]
    }
}

{
    "3": { "2": ["12"]},
    "4": { "2": ["13"]},
    "5": { "2": ["14","23"]},
    "6": { "2": ["15","24"]}
}


[[4 5 6 7 8 9 9 9 9] [4 5 6 7 8 9 9 9 9] [4 5 6 7 8 9 9 9 9]] want [[1 2 3 4 5 6 7 8 9] [2 3 4 5 6 7 8 9] [3 4 5 6 7 8 9]]


Loop Depth and number ranges for 3 digit number

i = outter loop, x = first inner, y = second inner

For each loop of i, x runs 8 times ( numbers 2 - 9 )
For each loop of x, y runs 7 times ( number 3 - 9 )


nesting idea
i and x run to get list of 2 digit combinations and values
if number combinations are stored as []int can then run another loop to get 3 digit numbers
    this would need to dynamically find duplicates


loop would go
    i = 1 and 2 = 2
        remove 1 and 2 from list 3
        add 1,2,3 -> 9
    i = 1 and x = 3
        remove 3 from list 3


findDuplicates would remove 2


instead of performing removals, loops could increase the starting index