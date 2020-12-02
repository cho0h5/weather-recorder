ws = new WebSocket("ws://localhost:8080/ws"); // Todo: 호스트이름 자동 변경으로

ws.onopen = (event) => {
  let data = { event: "enter" }; // Todo: 식별기능 추가(id)
  ws.send(JSON.stringify(data));

  requestRecentData(1);
};

ws.onmessage = (event) => {
  const data = JSON.parse(event.data);
  console.log(data); // 삭제
  updateDashboard(data);
};

function requestRecentData(n) {
  const query = { event: "recentData", n: n };
  ws.send(JSON.stringify(query));
}
