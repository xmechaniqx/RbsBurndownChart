// var Url="http://localhost:4000/chart?login=gchernetsov"
// var Url = window.location.href

const urlParams = new URLSearchParams(window.location.search);
const login = urlParams.get('login');
let Url = 'http://'+window.location.host+'/chart?login='+login
console.log(Url,login)
let allStruct = []
let datasetsArr = []
fnRequest(Url)

function daysInMonth (month, year) {
    return new Date(year, month, 0).getDate();
}
function isWeekend(date = new Date.now()) {
        return date.getDay() === 6 || date.getDay() === 0;
      }


function renderChart(structsEntered) {

const ctx = document.getElementById('myChart');
const date = new window.Date();  
const month = date.toLocaleString('default', { month: 'short' });
const day = date.toLocaleString('default', { day: "numeric" });
// console.log(day)
var labels = [];

//how many days in month

  var today = new Date();
var dd = String(today.getDate()).padStart(1,'0');
// console.log(dd)
for (i=3;i<31;i++) {
    labels.push(i+' '+month)
    } 
//weekend today or not    
      // console.log(isWeekend(date),date);
      const d1 = new Date('2022-09-24');
      // console.log("today",d1,date); 
      // console.log(d1.getDay()); 

    //   [{
    //     label: struct.name,
    //     data: struct.burnDownDays,
    // // data: [48,35,29,15,5],
    //     borderWidth: 2,
    //     borderColor: struct.color,
    //     backgroundColor: [struct.color],
    //     tension:0,
    //     fill: false
    //     }]



console.log("entry",structsEntered.length);
for (i=0; i<structsEntered.length; i++) {
  struct=structsEntered[i]
let dataset = {
    label: struct.nameProject,
    data: struct.burnDownDays,
    borderWidth: 2,
    borderColor: struct.color,
    backgroundColor: [struct.color],
    tension:0,
    fill: false
  }
datasetsArr.push(dataset)
}
new Chart(ctx, {
    type: 'line',
    data: {
          labels: labels,
          datasets: datasetsArr
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
  }
  })  
// new Chart(ctx, {
//   type: 'line',
//   data: {
//         labels: labels,
//         datasets: [{
//           label: struct.name,
//           data: struct.burnDownDays,
//       // data: [48,35,29,15,5],
//           borderWidth: 2,
//           borderColor: struct.color,
//           backgroundColor: [struct.color],
//           tension:0,
//           fill: false
//           }]
//   },
//   options: {
//           responsive: true,
//           plugins: { 
//             colors: {
//               enabled: false
//             },
//             legend: {
//               position: 'top',
//             },
//             title: {
//               display: true,
//               text: 'Chart.js Line Chart'
//             }
//           }
// }
// })
}


function fnRequest(url) {
  var req = new XMLHttpRequest();
  req.addEventListener("load", renderResponse);
  req.open("GET", url);
  req.send();
}
//Вспомогательная функция для fnRequest(), парсинг полученной страницы формата JSON и получение необходимых параметров
function renderResponse() {
  var resp = JSON.parse(this.response);
  
  // var charts = parse(resp.Charts)
  for (j=0;j<resp.Charts.length;j++) {
    // var j = 0
  var tasksCostSum = resp.Charts[j].TasksCostSum
  // console.log("dlina",resp.Charts.length)
  var hours =  new Array
  var burnDownDays = new Array
  
  burnDownDays.push(tasksCostSum)
  for (value in resp.Charts[j].WorkingHours) {
  hours.push(resp.Charts[j].WorkingHours[value]);
}
// console.log(hours)

while (tasksCostSum>0) {
for (i=0;i<hours.length;i++){
  tasksCostSum=tasksCostSum-hours[i]
  if (tasksCostSum<=0){
    tasksCostSum=0
    burnDownDays.push(tasksCostSum)
    // console.log("break")
    break
  }else{
  burnDownDays.push(tasksCostSum)
  }
}
}
// console.log(burnDownDays)

// console.log("odna est",j) 
let structResult = {     // объект
  name: resp.Developer.FullName,  // под ключом "name" хранится значение "John"
  position: resp.Developer.Position,
  nameProject: resp.Charts[j].Project.Name,
  burnDownDays: burnDownDays,
  color: generateNewColor()
};
console.log("structura",structResult) 

allStruct.push(structResult)


}  
renderChart(allStruct)
}
console.log("allStruct",allStruct)




const hexCharacters = [0,1,2,3,4,5,6,7,8,9,"A","B","C","D","E","F"]

function getCharacter(index) {
	return hexCharacters[index]
}

function generateNewColor() {
	let hexColorRep = "#"

	for (let index = 0; index < 6; index++){
		const randomPosition = Math.floor ( Math.random() * hexCharacters.length ) 
    	hexColorRep += getCharacter( randomPosition )
	}
	
	return hexColorRep
}

