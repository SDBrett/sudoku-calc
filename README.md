# sudoku-calc
Learning project to create a sudoku calculator

## Web Interface

This application now includes a modern web interface for easy interaction with the sudoku calculator.

### Running the Application

1. Start the server:
   ```bash
   go run main.go
   ```

2. Open your web browser and navigate to:
   ```
   http://localhost:8080
   ```

### Features

- **Interactive Form**: Easy-to-use web form for specifying calculation parameters
- **Real-time Results**: Instant display of number combinations
- **Constraint Support**: Include/exclude specific numbers
- **Responsive Design**: Works on desktop and mobile devices
- **Modern UI**: Clean, professional interface with smooth animations

### API Endpoints

- `GET /` - Web interface
- `GET /dataset` - Get the complete dataset
- `POST /dataset` - Query the dataset with specific parameters

### Example API Usage

```bash
curl -X POST http://localhost:8080/dataset \
  -H "Content-Type: application/json" \
  -d '{
    "numberOfDigits": 3,
    "value": 15,
    "excludedNumbers": ["1"],
    "includedNumbers": ["5"]
  }'
```
