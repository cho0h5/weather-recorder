const dht22_Humi = document.getElementById("dht22_Humi");
const dht22_Temp = document.getElementById("dht22_Temp");
const bmp180_Temp = document.getElementById("bmp180_Temp");
const bmp180_Pres = document.getElementById("bmp180_Pres");

const dht22_Humi_progress = document.getElementById("dht22_Humi_progress");

const header_section = document.getElementById("header-section");
const subtitle = document.getElementById("subtitle");

function updateDashboard(data) {
  dht22_Humi.innerHTML = data[0].dht22_Humi + "%";
  dht22_Temp.innerHTML = data[0].dht22_Temp + "C";
  bmp180_Temp.innerHTML = data[0].bmp180_Temp + "C";
  bmp180_Pres.innerHTML = data[0].bmp180_Pres + "Pa";

  dht22_Humi_progress.setAttribute("value", data[0].dht22_Humi);

  if (data[0].isWorking) {
    dht22_Humi_progress.classList.add("is-info");
    dht22_Humi_progress.classList.remove("is-danger");
  } else {
    dht22_Humi_progress.classList.add("is-danger");
    dht22_Humi_progress.classList.remove("is-info");
  }
}

function updateHeader(data) {
  if (data[0].isWorking) {
    header_section.classList.add("is-info");
    header_section.classList.remove("is-danger");

    subtitle.innerHTML = "isWorking " + data[0].datetime;
  } else {
    header_section.classList.add("is-danger");
    header_section.classList.remove("is-info");

    subtitle.innerHTML = "isNotWorking " + data[0].datetime;
  }
}

function sendRequest() {
  requestRecentData(10);
}
setInterval(sendRequest, 1000);
