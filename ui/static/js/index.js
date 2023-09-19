document.addEventListener("DOMContentLoaded", function() {
  fnRequest(url);
});

//Инициализируем параметр со значением текущей ссылки из адресной строки
const urlParams = new URLSearchParams(window.location.search);
//Получаем логин из полученного параметра
const login = urlParams.get('login');
//Формируем ссылку добавляя адрес хэндлера /chart
let url = 'http://'+window.location.host+'/chart?login='+login;
//Объявляем массивы для заполнения итоговой структуры графика
let allProjectsForDeveloper = [];  //массив состоящий из всех проектов для конкретного пользователя
let nameProject = [];              //массив наименований проектов
let structResult = new Object;     //объект собирающий все исходные параметры
let datasetsArr = [];              //массив объектов dataset
let arrOfProjects = [];            //массив объектов structResult

//getDaysInMonth() - функция возвращает количество дней в текущем месяце
function getDaysInMonth(year, month) {
  return new Date(year, month, 0).getDate();
}
const date = new Date();
const currentYear = date.getFullYear();
const currentMonth = date.getMonth() + 1;
const daysInCurrentMonth = getDaysInMonth(currentYear,currentMonth);

//renderChart() - функция формирует и выводит график на странице на основании полученной структуры
function renderChart(structsEntered) {
  const ctx = document.getElementById('myChart');
  let daysOfWeek = ["Пн","Вт","Ср","Чт","Пт","Сб","Вс"];
  let labels = [];
  let looper = 0;

  //Формируем массив дней месяца для актуального месяца
  for (i = 0; i < daysInCurrentMonth; i++) {
    //looper сбрасывается каждый раз когда массив заполнен всеми днями недели
    if (looper>6){
      looper=0;
    }
    labels.push(daysOfWeek[looper])
    looper++;
  } 
  //Заполняем структуру dataset из полученного объекта
  structsEntered.projects.forEach(struct => {
    let dataset = new Object;
    dataset = {
      label: struct.nameProject,
    data: struct.burnDownDays,
      borderWidth: 2,
      borderColor: struct.color,
      backgroundColor: [struct.color],
      tension: 0,
      fill: false
    };
    //Заполняем массив полученными структурами
    datasetsArr.push(dataset)
  });

  //Создаем новую функцию отображения графика
  new Chart(ctx, {
    type: 'line',
    data: {
          labels: labels,
          datasets: datasetsArr
          },
    options: {
          responsive: true,
          plugins: {colors: {enabled: false},legend: {position: 'top'},
          title: {display: true, text: 'Chart.js Line Chart'}
          }
    }
  });
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
  //В случае получения параметра ошибки 1 выводим на странице сообщение иначе продолжаем обработку
  if (resp.Status==1){
      document.write('<b><font size="5">'+resp.Error+'</font></b>');
      return
  }
  //В цикле формируем объекты данных из полученной исходной структуры
  resp.BurndownChart.Charts.forEach(element => {
    let tasksCostSum = element.TasksCostSum;
    let hours =  new Array;
    //burnDownDays - массив убывания исходного времени для проекта в зависимости от производительности команды
    let burnDownDays = new Array;

    burnDownDays.push(tasksCostSum);
    //В ходе цикла получаем отображения рабочего времени по дням недели
    for (value in element.WorkingHours) {
      hours.push(element.WorkingHours['Понедельник']);
      hours.push(element.WorkingHours['Вторник']);
      hours.push(element.WorkingHours['Среда']);
      hours.push(element.WorkingHours['Четверг']);
      hours.push(element.WorkingHours['Пятница']);
      hours.push(element.WorkingHours['Суббота']);
      hours.push(element.WorkingHours['Воскресенье']);
    }
    while (tasksCostSum > 0) {
      for (i = 0; i < hours.length; i++) {
        tasksCostSum = tasksCostSum - hours[i];
          if (tasksCostSum <= 0) {
            tasksCostSum = 0
            burnDownDays.push(tasksCostSum);
            break
          } else {
            burnDownDays.push(tasksCostSum);
          }
      }
    }
    project = {
      nameProject: element.Project.Name,
      burnDownDays: burnDownDays,
      sprintDate: element.Project.SprintStartDate,
      color: generateNewColor()
    }
    arrOfProjects.push(project); 
  });
  structResult = {
    name: resp.BurndownChart.Developer.FullName,
    position: resp.BurndownChart.Developer.Position,
    projects: arrOfProjects,
    
  };
  document.getElementById("Name").innerHTML = structResult.name;
  document.getElementById("Position").innerHTML = structResult.position;
  structResult.projects.forEach(name => {
    allProjectsForDeveloper.push(' ' + name.nameProject);
  });
  document.getElementById("Project").innerHTML = allProjectsForDeveloper;
  renderChart(structResult);
};
//generateNewColor() - функция позволяет случайным образом сгенерировать цвет
const hexCharacters = [0,1,2,3,4,5,6,7,8,9,"A","B","C","D","E","F"]
function generateNewColor() {
	let hexColorRep = "#";
	for (let index = 0; index < 6; index++) {
		const randomPosition = Math.floor ( Math.random() * hexCharacters.length );
    hexColorRep += hexCharacters[randomPosition];
	}
	return hexColorRep
};

