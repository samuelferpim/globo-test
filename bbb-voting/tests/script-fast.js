import http from 'k6/http';
import { check, sleep } from 'k6';

export const options = {
    scenarios: {
        morning: {
            executor: 'constant-arrival-rate',
            rate: 1000,
            timeUnit: '1s',
            duration: '20s',
            preAllocatedVUs: 100,
            maxVUs: 1000,
            exec: 'morningScenario',
            startTime: '0s',
        },
        afternoon: {
            executor: 'constant-arrival-rate',
            rate: 1000,
            timeUnit: '1s',
            duration: '20s',
            preAllocatedVUs: 100,
            maxVUs: 1000,
            exec: 'afternoonScenario',
            startTime: '20s',
        },
        evening: {
            executor: 'constant-arrival-rate',
            rate: 1000,
            timeUnit: '1s',
            duration: '20s',
            preAllocatedVUs: 100,
            maxVUs: 1000,
            exec: 'eveningScenario',
            startTime: '40s',
        },
    },
};

const BASE_URL = 'http://localhost:8080';

const participants = [
    { id: 'participant123', weight: 0.67 },
    { id: '2', weight: 0.165 },
    { id: '3', weight: 0.165 },
];

function selectParticipant() {
    const r = Math.random();
    let cumulativeWeight = 0;
    for (const participant of participants) {
        cumulativeWeight += participant.weight;
        if (r <= cumulativeWeight) {
            return participant.id;
        }
    }
    return participants[participants.length - 1].id;
}

function castVote() {
    const participantId = selectParticipant();
    const payload = JSON.stringify({
        participant_id: participantId,
        voter_id: `voter_${__VU}_${__ITER}`,
        timestamp: new Date().toISOString(),
        ip_address: '192.168.1.1',
        user_agent: 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36',
        region: 'BR-SP',
        device_type: 'desktop',
        captcha_completed: true
    });

    const headers = { 'Content-Type': 'application/json' };
    const response = http.post(`${BASE_URL}/vote`, payload, { headers: headers });

    check(response, {
        'status is 200': (r) => r.status === 200,
    });
}

export function morningScenario() {
    castVote();
    sleep(0.01);
}

export function afternoonScenario() {
    castVote();
    sleep(0.01);
}

export function eveningScenario() {
    castVote();
    sleep(0.01);
}