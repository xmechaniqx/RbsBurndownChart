//Инициализируем параметр со значением текущей ссылки из адресной строки
const urlParams = new URLSearchParams(window.location.search);
//Получаем логин из полученного параметра
const login = urlParams.get('login');
//Формируем ссылку добавляя адрес хэндлера /chart
let url = 'http://'+window.location.host+'/chart?login='+login
let today = new Date().toISOString().slice(0, 10)
// console.log(today)
let allProjectsForDeveloper = []
let allStruct = []
let datasetsArr = []
fnRequest(url)

function daysInMonth (month, year) {
    return new Date(year, month, 0).getDate();
}
function isWeekend(date = new Date.now()) {
        return date.getDay() === 6 || date.getDay() === 0;
}

//renderChart() - функция формирует и выводит график на странице на основании полученной структуры
function renderChart(structsEntered) {
const ctx = document.getElementById('myChart');
const date = new window.Date();  
const month = date.toLocaleString('default', { month: 'short' });
// const day = date.toLocaleString('default', { day: "numeric" });
let labels = [];
//формируем массив дней месяца для актуального месяца
for (i=1;i<31;i++) {
    labels.push(i+' '+month)
    } 

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
      scales: {
        x: {
            min: struct.SprintStartDate,
        }
      },
        responsive: true,
        plugins: {colors: {enabled: false},legend: {position: 'top'},
        title: {display: true, text: 'Chart.js Line Chart'}
        }
    }
  }) 
}

//fnRequest() - функция создает запрос на заданный адрес
function fnRequest(someUrl) {
  let req = new XMLHttpRequest();
  req.addEventListener("load", renderResponse);
  req.open("GET", someUrl);
  req.send();
}
//Вспомогательная функция для fnRequest(), парсинг полученной страницы формата JSON и получение необходимых параметров
function renderResponse() {
  let resp = JSON.parse(this.response);
//В цикле формируем объекты данных из полученной исходной структуры
  for (j=0;j<resp.Charts.length;j++) {
    let sprintDate = resp.Charts[j].Project.SprintStartDate
    let tasksCostSum = resp.Charts[j].TasksCostSum
    let hours =  new Array
    let burnDownDays = new Array
  
    burnDownDays.push(tasksCostSum)
    for (value in resp.Charts[j].WorkingHours) {
      hours.push(resp.Charts[j].WorkingHours['Понедельник']);
      hours.push(resp.Charts[j].WorkingHours['Вторник']);
      hours.push(resp.Charts[j].WorkingHours['Среда']);
      hours.push(resp.Charts[j].WorkingHours['Четверг']);
      hours.push(resp.Charts[j].WorkingHours['Пятница']);
      hours.push(resp.Charts[j].WorkingHours['Суббота']);
      hours.push(resp.Charts[j].WorkingHours['Воскресенье']);
    }

    while (tasksCostSum>0) {
      for (i=0;i<hours.length;i++){
        tasksCostSum=tasksCostSum-hours[i]
          if (tasksCostSum<=0){
            tasksCostSum=0
              burnDownDays.push(tasksCostSum)
                break
          }else{
            burnDownDays.push(tasksCostSum)
          }
      }
    }

    let structResult = {     // объект
      name: resp.Developer.FullName,  // под ключом "name" хранится значение "John"
      position: resp.Developer.Position,
      nameProject: resp.Charts[j].Project.Name,
      burnDownDays: burnDownDays,
      sprintDate:sprintDate,
      color: generateNewColor()
    };
    allStruct.push(structResult)
  }  
    document.getElementById("Name").innerHTML = allStruct[0].name;
    document.getElementById("Position").innerHTML = allStruct[0].position;
  for(i=0;i<allStruct.length;i++){ 
    allProjectsForDeveloper.push(' '+allStruct[i].nameProject)
  }
    document.getElementById("Project").innerHTML = allProjectsForDeveloper
    renderChart(allStruct)
}

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

