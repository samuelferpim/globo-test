const API_URL = '/api';

const participants = [
    { id: '1', name: 'Goku', image: '../static/images/goku.jpg' },
    { id: '2', name: 'Naruto', image: '../static/images/naruto.jpg' },
    { id: '3', name: 'Luffy', image: '../static/images/luffy.jpg' },
];

let captchaAnswer;

function displayParticipants() {
    const container = document.getElementById('participants');
    participants.forEach(participant => {
        const div = document.createElement('div');
        div.className = 'participant';
        div.innerHTML = `
            <img src="${participant.image}" alt="${participant.name}">
            <h3>${participant.name}</h3>
        `;
        container.appendChild(div);
    });

    const select = document.getElementById('participant-select');
    participants.forEach(participant => {
        const option = document.createElement('option');
        option.value = participant.id;
        option.textContent = participant.name;
        select.appendChild(option);
    });
}

function generateFakeCaptcha() {
    const num1 = Math.floor(Math.random() * 10);
    const num2 = Math.floor(Math.random() * 10);
    captchaAnswer = num1 + num2;

    document.getElementById('num1').textContent = num1;
    document.getElementById('num2').textContent = num2;
}

function vote() {
    const participantId = document.getElementById('participant-select').value;
    const userAnswer = parseInt(document.getElementById('captcha-answer').value);

    if (userAnswer !== captchaAnswer) {
        alert('CAPTCHA incorreto. Por favor, tente novamente.');
        generateFakeCaptcha();
        return;
    }

    // Simular o armazenamento do voto
    let votes = JSON.parse(localStorage.getItem('votes')) || {};
    votes[participantId] = (votes[participantId] || 0) + 1;
    localStorage.setItem('votes', JSON.stringify(votes));

    alert('Voto registrado com sucesso!');
    getResults();
    generateFakeCaptcha();
    document.getElementById('captcha-answer').value = '';
}

function getResults() {
    const votes = JSON.parse(localStorage.getItem('votes')) || {};
    const results = participants.map(participant => ({
        participant_id: participant.name,
        vote_count: votes[participant.id] || 0
    }));

    const totalVotes = results.reduce((sum, result) => sum + result.vote_count, 0);
    results.forEach(result => {
        result.percentage = totalVotes > 0 ? (result.vote_count / totalVotes) * 100 : 0;
    });

    displayResults(results);
}

function displayResults(results) {
    const container = document.getElementById('results-container');
    container.innerHTML = '';
    results.forEach(result => {
        const div = document.createElement('div');
        div.className = 'result-bar';
        div.innerHTML = `
            <div class="result-fill" style="width: ${result.percentage}%">
                ${result.participant_id}: ${result.vote_count} votos (${result.percentage.toFixed(2)}%)
            </div>
        `;
        container.appendChild(div);
    });
}

document.addEventListener('DOMContentLoaded', () => {
    displayParticipants();
    generateFakeCaptcha();
    getResults();
    document.getElementById('vote-button').addEventListener('click', vote);
});