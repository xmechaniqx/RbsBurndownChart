// var currentLocation = window.location;
// var Url=window.location.href;
// console.log(Url)
// fnRequest(Url) 
var request=document.addEventListener("DOMContentLoaded", XMLHttpRequest(login));

const ctx = document.getElementById('myChart');
const date = new window.Date();  
const month = date.toLocaleString('default', { month: 'short' });
const day = date.toLocaleString('default', { day: "numeric" });
console.log(day)
var labels = [];

//how many days in month
function daysInMonth (month, year) {
    return new Date(year, month, 0).getDate();
}
  var today = new Date();
var dd = String(today.getDate()).padStart(1,'0');
console.log(dd)
for (i=3;i<31;i++) {
    labels.push(i+' '+month)
    } 
//weekend today or not    
function isWeekend(date = new Date.now()) {
        return date.getDay() === 6 || date.getDay() === 0;
      }
      console.log(isWeekend(date),date);
      const d1 = new Date('2022-09-24');
      console.log("today",d1,date); 
      console.log(d1.getDay()); 



var chart=new Chart(ctx, {
  type: 'line',
  data: {
    labels: labels,
    datasets: [{
      label: 'Antonenko Roman',
      data: [48,35,29,15,5],
      borderWidth: 2,
      borderColor: 'blue',
      backgroundColor: [
        'rgba(55, 119, 196, 1)'],
      tension:0,
      fill: false
    },{
        label: 'task "Burndown chart"',
        data: [48, 42, 36, 30, 24, 18,12,6,0],
        borderWidth: 2,
        borderColor: 'red',
        backgroundColor: [
          'rgba(252, 86, 86, 1)'],
        tension:0,
        fill: false
      }
]
    
  },
  options: {
    responsive: true,
    plugins: { 
        colors: {
        enabled: false
          },
      legend: {
       
        position: 'top',
      },
      title: {
        display: true,
        text: 'Chart.js Line Chart'
      }
    }
  },
});

function fnRequest(url) {
  var req = new XMLHttpRequest();
  req.addEventListener("load", renderResponse());
  req.open("GET", url);
  req.send();
}
//Вспомогательная функция для fnRequest(), парсинг полученной страницы формата JSON и получение необходимых параметров
function renderResponse() {
  var resp = JSON.parse(this.response);
  console.log(resp.Date)
  // var ul = document.getElementById('result');
  //Очистка тела UL таблицы как способ обновления страницы для отображения содержимого директорий при переходе
  // ul.innerHTML = '';
  //
  // console.log(resp)
  // resp.VFSNode_struct.forEach(function (element) {
      // var li = document.createElement("li");
      // li.setAttribute("vfs_path", element.path);
      // if (element.stat == "dir") {
      //     li.innerHTML = '<span ><div class="results"><img src="/static/img/folder.png" width="1%">' + element.path + '</div></span>';
      // }
      // if (element.stat == "file") {
      //     li.innerHTML = '<span ><div class="results"><img src="/static/img/file.png" width="1%">' + element.path + '</div></span>';
      // }
      //Записываем новый путь перехода по директории в переменную newUrl
      // var newUrl = window.location.href + 'flag?root=' + element.path + '/';
      // var root = resp.root;
      // curPath(root);
      // if (element.stat == "dir") {
      //     li.addEventListener("click", function () {
      //         fnRequest(newUrl);
      //     }, false);
      // }
      // ul.appendChild(li);
      // backURL = removeLastDirectoryPartOf(root);
      // backURL = removeLastDirectoryPartOf(backURL);
      // backURL = 'flag?root=' + backURL + '/';
  // });
}