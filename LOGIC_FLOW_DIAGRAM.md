# Sudoku Calculator Go Code Logic Flow Diagram

## Application Startup Flow
```
┌─────────────────────────────────────────────────────────────────┐
│                        Application Start                        │
└─────────────────────┬───────────────────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                    main() Function                              │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 1. Get CALC_PORT environment variable                      │ │
│  │    - If empty or invalid → default to "8080"               │ │
│  │    - Validate port is numeric                              │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 2. Initialize Gin router                                    │ │
│  │    - Set up static file serving (/static)                  │ │
│  │    - Load HTML templates                                    │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 3. Register Routes:                                       │ │
│  │    - GET  / → serve index.html                             │ │
│  │    - GET  /dataset → getDataset()                          │ │
│  │    - POST /dataset → queryDataSet()                        │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 4. Start HTTP server on configured port                     │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

## Data Generation Flow (Startup)
```
┌─────────────────────────────────────────────────────────────────┐
│                Global Variable Initialization                   │
│                    var ds = GenerateDataSet()                   │
└─────────────────────┬───────────────────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                  GenerateDataSet() Function                     │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 1. Initialize empty DataSet structure                      │ │
│  │    - Map[int]map[int][]string (digits → value → combinations)│ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 2. For each digit count (2-9):                              │ │
│  │    - Generate all possible number combinations              │ │
│  │    - Calculate sum for each combination                     │ │
│  │    - Store in DataSet[digits][sum] = combinations          │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

## API Request Flow
```
┌─────────────────────────────────────────────────────────────────┐
│                    HTTP Request Received                        │
└─────────────────────┬───────────────────────────────────────────┘
                      │
                      ▼
┌─────────────────────────────────────────────────────────────────┐
│                    Route Handler                                │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ GET /dataset → getDataset()                                │ │
│  │    ↓                                                        │ │
│  │    Return entire dataset as JSON                           │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ POST /dataset → queryDataSet()                             │ │
│  │    ↓                                                        │ │
│  │    Process query request                                    │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

## Query Processing Flow
```
┌─────────────────────────────────────────────────────────────────┐
│                    queryDataSet() Function                      │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 1. Parse JSON request into DataSetQuery struct             │ │
│  │    - numberOfDigits, value, excludedNumbers, etc.          │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 2. Call ds.Query(dsq)                                      │ │
│  │    ↓                                                        │ │
│  │    Process the query                                        │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 3. Return JSON response or error                           │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

## Dataset Query Processing Flow
```
┌─────────────────────────────────────────────────────────────────┐
│                    ds.Query() Function                         │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 1. Initialize empty DataSetQueryResponse                   │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 2. Validate query parameters (dsq.Validate())              │ │
│  │    - Check numberOfDigits (2-9)                           │ │
│  │    - Check value is within valid range for digit count     │ │
│  │    - Validate included/excluded numbers (1-9)              │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 3. Get base combinations from dataset                      │ │
│  │    combinations = ds[numberOfDigits][value]                 │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 4. Filter combinations (GetValidCombinations)              │ │
│  │    - Apply excluded numbers filter                          │ │
│  │    - Apply included numbers filter                          │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 5. Check if analysis requested                              │ │
│  │    - GetNumbersNotPresent?                                 │ │
│  │    - GetNumbersPresentAllCombinations?                     │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 6. If analysis requested:                                  │ │
│  │    - Call AnalyzeCombinations()                            │ │
│  │    - Populate analysis fields in response                  │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 7. Return DataSetQueryResponse                             │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

## Combination Analysis Flow
```
┌─────────────────────────────────────────────────────────────────┐
│                AnalyzeCombinations() Function                   │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 1. Initialize CombinationAnalysis struct                   │ │
│  │    - NumbersNotPresent: []string                            │ │
│  │    - NumbersInAllCombinations: []string                     │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 2. Count number occurrences across all combinations        │ │
│  │    - Create map[string]int for digit counts                │ │
│  │    - Iterate through each combination                       │ │
│  │    - Count each digit occurrence                            │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 3. Find numbers not present (count == 0)                   │ │
│  │    - Check digits 1-9                                       │ │
│  │    - Add to NumbersNotPresent if count == 0                 │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 4. Find numbers in all combinations (count == total)      │ │
│  │    - Check digits 1-9                                       │ │
│  │    - Add to NumbersInAllCombinations if count == total      │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ 5. Return CombinationAnalysis                               │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

## Data Structures
```
┌─────────────────────────────────────────────────────────────────┐
│                        Data Structures                          │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ DataSetQuery:                                               │ │
│  │ - NumberOfDigits: int                                      │ │
│  │ - Value: int                                                │ │
│  │ - NumbersToExclude: []string                               │ │
│  │ - NumbersToInclude: []string                               │ │
│  │ - GetNumbersNotPresent: bool                               │ │
│  │ - GetNumbersPresentAllCombinations: bool                  │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ DataSetQueryResponse:                                      │ │
│  │ - Combinations: []string                                   │ │
│  │ - DigitsInAllCombinations: []string                        │ │
│  │ - DigitsAbsentFromAllCombinations: []string               │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ DataSet: map[int]map[int][]string                          │ │
│  │ - Outer key: number of digits (2-9)                        │ │
│  │ - Inner key: target sum value                               │ │
│  │ - Value: list of valid combinations                        │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

## Error Handling Flow
```
┌─────────────────────────────────────────────────────────────────┐
│                      Error Handling                             │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ JSON Parsing Errors:                                        │ │
│  │ - Return HTTP 400 with "Invalid JSON format"               │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ Validation Errors:                                          │ │
│  │ - Invalid digit count (not 2-9)                           │ │
│  │ - Invalid sum value (outside range)                        │ │
│  │ - Invalid included/excluded numbers (not 1-9)              │ │
│  │ - Return HTTP 400 with specific error message              │ │
│  └─────────────────────────────────────────────────────────────┘ │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ No Results Found:                                           │ │
│  │ - Return HTTP 400 with "no combinations found"            │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```

## Key Functions Summary
```
┌─────────────────────────────────────────────────────────────────┐
│                        Key Functions                             │
│  ┌─────────────────────────────────────────────────────────────┐ │
│  │ main() - Application entry point, server setup              │ │
│  │ getDataset() - Returns entire dataset                       │ │
│  │ queryDataSet() - Processes query requests                   │ │
│  │ GenerateDataSet() - Pre-computes all combinations           │ │
│  │ ds.Query() - Main query processing logic                    │ │
│  │ dsq.Validate() - Input validation                          │ │
│  │ GetValidCombinations() - Filters combinations               │ │
│  │ AnalyzeCombinations() - Performs analysis                  │ │
│  │ calculateMinSum() - Calculates minimum possible sum         │ │
│  │ calculateMaxSum() - Calculates maximum possible sum        │ │
│  └─────────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────────┘
```
