import http from "k6/http";

export let options = {
    stages: [
        { duration: "10s", target: 100 },
    ],
    thresholds: {
        http_req_failed: ['rate<0.01'],
        http_req_duration: ['p(95)<200'],
    },
};

export default function () {
    let response = http.batch([
        {
            method: "POST",
            url: "http://app:8080/sum",
            body: JSON.stringify([1, 2, 3, 4, 5]),
            params: {
                headers: { 'Content-Type': 'application/json' },
            },
        },
    ]);
}
