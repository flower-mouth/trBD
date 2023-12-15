document.addEventListener('DOMContentLoaded', function() {
    document.getElementById('pointsParticipant1').addEventListener('change', function() {
        const points1 = parseFloat(this.value);
        const points2Select = document.getElementById('pointsParticipant2');

        if (points1 === 0.0) {
            points2Select.value = '1.0';
        } else if (points1 === 1.0) {
            points2Select.value = '0.0';
        } else if (points1 === 0.5) {
            points2Select.value = '0.5';
        }
    });

    document.getElementById('pointsParticipant2').addEventListener('change', function() {
        const points2 = parseFloat(this.value);
        const points1Select = document.getElementById('pointsParticipant1');

        if (points2 === 0.0) {
            points1Select.value = '1.0';
        } else if (points2 === 1.0) {
            points1Select.value = '0.0';
        } else if (points2 === 0.5) {
            points1Select.value = '0.5';
        }
    });
});
