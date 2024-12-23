const API_URL = 'http://localhost:8080';

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

async function vote() {
    const participantId = document.getElementById('participant-select').value;
    const userAnswer = parseInt(document.getElementById('captcha-answer').value);

    if (userAnswer !== captchaAnswer) {
        alert('CAPTCHA incorreto. Por favor, tente novamente.');
        generateFakeCaptcha();
        return;
    }

    try {
        const response = await fetch(`${API_URL}/vote`, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                participant_id: participantId,
                voter_id: `voter_${Date.now()}`,
                timestamp: new Date().toISOString(),
                ip_address: '127.0.0.1',
                user_agent: navigator.userAgent,
                region: 'BR-SP',
                device_type: 'web',
                captcha_completed: true
            }),
        });

        if (!response.ok) {
            throw new Error('Falha ao registrar o voto');
        }

        alert('Voto registrado com sucesso!');
        debouncedGetResults();
        getResults();
        generateFakeCaptcha();
        document.getElementById('captcha-answer').value = '';
    } catch (error) {
        console.error('Erro ao votar:', error);
        alert('Ocorreu um erro ao tentar votar. Por favor, tente novamente.');
    }
}

async function getResults(retries = 3) {
    try {
        const response = await fetch(`${API_URL}/result`);
        if (!response.ok) {
            throw new Error('Falha ao obter resultados');
        }
        const data = await response.json();
        displayResults(data);
    } catch (error) {
        console.error('Erro ao obter resultados:', error);
        if (retries > 0) {
            console.log(`Tentando novamente. Tentativas restantes: ${retries - 1}`);
            setTimeout(() => getResults(retries - 1), 2000);
        } else {
            alert('Ocorreu um erro ao tentar obter os resultados. Por favor, tente novamente mais tarde.');
        }
    }
}

function displayResults(data) {
    const totalVotesElement = document.getElementById('total-votes');
    const resultsContainer = document.getElementById('results-container');
    const votesByHourTable = document.getElementById('votes-by-hour-table').getElementsByTagName('tbody')[0];

    totalVotesElement.innerHTML = `<h3>Total Geral de Votos: ${data.total_votes}</h3>`;

    resultsContainer.innerHTML = '';
    votesByHourTable.innerHTML = '';

    data.participant_results.forEach(result => {
        const participantName = participants.find(p => p.id === result.participant_id)?.name || result.participant_id;
        const percentage = (result.vote_count / data.total_votes) * 100;

        const div = document.createElement('div');
        div.className = 'result-bar';
        div.innerHTML = `
            <div class="participant-name">${participantName}</div>
            <div class="progress-bar">
                <div class="progress" style="width: ${percentage}%"></div>
            </div>
            <div class="vote-info">${result.vote_count} votos (${percentage.toFixed(2)}%)</div>
        `;
        resultsContainer.appendChild(div);
    });

    const sortedHours = Object.entries(data.votes_by_hour).sort((a, b) => a[0].localeCompare(b[0]));
    sortedHours.forEach(([hour, count]) => {
        const row = votesByHourTable.insertRow();
        row.insertCell(0).textContent = formatHour(hour);
        row.insertCell(1).textContent = count;
    });
}

function debounce(func, wait) {
    let timeout;
    return function executedFunction(...args) {
        const later = () => {
            clearTimeout(timeout);
            func(...args);
        };
        clearTimeout(timeout);
        timeout = setTimeout(later, wait);
    };
}

const debouncedGetResults = debounce(getResults, 1000);

function formatHour(hourString) {
    const year = hourString.slice(0, 4);
    const month = hourString.slice(4, 6);
    const day = hourString.slice(6, 8);
    const hour = hourString.slice(8, 10);
    return `${day}/${month}/${year} ${hour}:00`;
}

document.addEventListener('DOMContentLoaded', () => {
    displayParticipants();
    generateFakeCaptcha();
    getResults();
    document.getElementById('vote-button').addEventListener('click', vote);
});