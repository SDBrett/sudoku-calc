document.addEventListener('DOMContentLoaded', function() {
    const form = document.getElementById('sudokuForm');
    const resultsContainer = document.getElementById('results');
    const errorContainer = document.getElementById('error');
    const resultsContent = document.getElementById('resultsContent');
    const errorContent = document.getElementById('errorContent');

    form.addEventListener('submit', async function(e) {
        e.preventDefault();
        
        // Hide previous results/errors
        resultsContainer.style.display = 'none';
        errorContainer.style.display = 'none';
        
        // Show loading state
        resultsContent.innerHTML = '<div class="loading">Calculating combinations...</div>';
        resultsContainer.style.display = 'block';
        
        try {
            // Collect form data
            const formData = new FormData(form);
            const data = {
                numberOfDigits: parseInt(formData.get('numberOfDigits')),
                value: parseInt(formData.get('value')),
                excludedNumbers: formData.getAll('excludedNumbers'),
                includedNumbers: formData.getAll('includedNumbers'),
                getNumbersNotPresent: formData.has('getNumbersNotPresent'),
                getNumbersPresentAllCombinations: formData.has('getNumbersPresentAllCombinations')
            };

            console.log('Sending data:', data);

            // Make API request
            const fetchResponse = await fetch('/dataset', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                },
                body: JSON.stringify(data)
            });

            const result = await fetchResponse.json();
            
            console.log('Received response:', result);
            
            if (!fetchResponse.ok) {
                throw new Error(result.error || `HTTP error! status: ${fetchResponse.status}`);
            }

            // Display results
            displayResults(result, data);
            
        } catch (error) {
            console.error('Error:', error);
            displayError(error.message);
        }
    });

    function displayResults(response, queryData) {
        let html = '';
        
        if (!response.combinations || response.combinations.length === 0) {
            html = '<p>No combinations found matching your criteria.</p>';
        } else {
            html = `<p><strong>Found ${response.combinations.length} combination(s):</strong></p>`;
            html += '<div class="combination-list">';
            
            response.combinations.forEach(combination => {
                html += `<div class="combination-item">${combination}</div>`;
            });
            
            html += '</div>';
            
            // Add analysis results if available
            if (response.digitsInAllCombinations && Array.isArray(response.digitsInAllCombinations) && response.digitsInAllCombinations.length > 0) {
                html += '<div style="margin-top: 20px; padding: 15px; background: #d4edda; border-radius: 6px;">';
                html += '<h3>Numbers Present in All Combinations:</h3>';
                html += `<p>${response.digitsInAllCombinations.join(', ')}</p>`;
                html += '</div>';
            }
            
            if (response.digitsAbsentFromAllCombinations && Array.isArray(response.digitsAbsentFromAllCombinations) && response.digitsAbsentFromAllCombinations.length > 0) {
                html += '<div style="margin-top: 20px; padding: 15px; background: #f8d7da; border-radius: 6px;">';
                html += '<h3>Numbers Not Present in Any Combination:</h3>';
                html += `<p>${response.digitsAbsentFromAllCombinations.join(', ')}</p>`;
                html += '</div>';
            }
            
            // Add summary information
            html += '<div style="margin-top: 20px; padding: 15px; background: #e9ecef; border-radius: 6px;">';
            html += '<h3>Query Summary:</h3>';
            html += `<p><strong>Number of digits:</strong> ${queryData.numberOfDigits}</p>`;
            html += `<p><strong>Target sum:</strong> ${queryData.value}</p>`;
            
            if (queryData.includedNumbers.length > 0) {
                html += `<p><strong>Must include:</strong> ${queryData.includedNumbers.join(', ')}</p>`;
            }
            
            if (queryData.excludedNumbers.length > 0) {
                html += `<p><strong>Must exclude:</strong> ${queryData.excludedNumbers.join(', ')}</p>`;
            }
            
            html += '</div>';
        }
        
        resultsContent.innerHTML = html;
        resultsContainer.style.display = 'block';
    }

    function displayError(message) {
        errorContent.innerHTML = `<p>Error: ${message}</p>`;
        errorContainer.style.display = 'block';
    }

    // Add some helpful interactions
    const numberOfDigitsSelect = document.getElementById('numberOfDigits');
    const valueInput = document.getElementById('value');
    
    // Update value input placeholder based on number of digits
    numberOfDigitsSelect.addEventListener('change', function() {
        const digits = parseInt(this.value);
        if (digits) {
            const minValue = calculateMinSum(digits);
            const maxValue = calculateMaxSum(digits);
            valueInput.min = minValue;
            valueInput.max = maxValue;
            valueInput.placeholder = `Enter target sum (${minValue}-${maxValue})`;
        }
    });

    // Prevent selecting the same number in both include and exclude
    const includeCheckboxes = document.querySelectorAll('input[name="includedNumbers"]');
    const excludeCheckboxes = document.querySelectorAll('input[name="excludedNumbers"]');
    
    includeCheckboxes.forEach(checkbox => {
        checkbox.addEventListener('change', function() {
            if (this.checked) {
                const excludeCheckbox = document.querySelector(`input[name="excludedNumbers"][value="${this.value}"]`);
                if (excludeCheckbox) {
                    excludeCheckbox.checked = false;
                }
            }
        });
    });
    
    excludeCheckboxes.forEach(checkbox => {
        checkbox.addEventListener('change', function() {
            if (this.checked) {
                const includeCheckbox = document.querySelector(`input[name="includedNumbers"][value="${this.value}"]`);
                if (includeCheckbox) {
                    includeCheckbox.checked = false;
                }
            }
        });
    });
});

// Helper functions to calculate valid sum ranges
function calculateMinSum(numberOfDigits) {
    let sum = 0;
    for (let i = 1; i <= numberOfDigits; i++) {
        sum += i;
    }
    return sum;
}

function calculateMaxSum(numberOfDigits) {
    let sum = 0;
    for (let i = 9 - numberOfDigits + 1; i <= 9; i++) {
        sum += i;
    }
    return sum;
}
