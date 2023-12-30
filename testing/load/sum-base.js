import http from "k6/http";

export let options = {
  insecureSkipTLSVerify: true,
  noConnectionReuse: false,
  stages: [{ duration: "10s", target: 100 }],
};

export default function () {
  const responses = http.batch([
    {
      method: "POST",
      url: `${__ENV.API_BASE_URL}/sum`,
      body: [1, 2, 3, 4, 5],
    },
  ]);

  check(responses[0], {
    "sum status 200": (res) => res.status === 200,
    "sum gives correct answear": (res) => res.body["result"] === "15",
  });
}
