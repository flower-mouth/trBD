document.addEventListener("DOMContentLoaded", function() {
    var position = 1;
    var rows = document.querySelectorAll('.results-table tbody tr');

    rows.forEach(function(row) {
        var cells = row.getElementsByTagName('td');
        cells[2].textContent = position++;
        
        if (position === 2) {
            row.classList.add('gold');
        } else if (position === 3) {
            row.classList.add('silver');
        } else if (position === 4) {
            row.classList.add('bronze');
        }
    });
});