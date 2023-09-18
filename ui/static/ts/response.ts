//fnRequest() - функция создает запрос на заданный адрес
  export function fnRequest(someUrl) {
    let req = new XMLHttpRequest();
    req.addEventListener("load", renderResponse);
    req.open("GET", someUrl);
    req.send();
  }
 
  let structResult={
    name: "",  // под ключом "name" хранится значение "John"
    position: "",
    nameProject: "",
    burnDownDays: [0],
    sprintDate:"",
    color: generateNewColor()
  }
  export {structResult} 
  //Вспомогательная функция для fnRequest(), парсинг полученной страницы формата JSON и получение необходимых параметров
function renderResponse() {
    
    
    let allStruct = [structResult]
    let resp = JSON.parse(this.response);
  //В цикле формируем объекты данных из полученной исходной структуры
  let j
    for (j=0;j<resp.Charts.length;j++) {
      let sprintDate = resp.Charts[j].Project.SprintStartDate
      let tasksCostSum = resp.Charts[j].TasksCostSum
      let hours =  new Array
      let burnDownDays = new Array
    
      burnDownDays.push(tasksCostSum)
      let value 
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
        let i 
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
      
      structResult = {     // объект
        name: resp.Developer.FullName,  // под ключом "name" хранится значение "John"
        position: resp.Developer.Position,
        nameProject: resp.Charts[j].Project.Name,
        burnDownDays: burnDownDays,
        sprintDate:sprintDate,
        color: generateNewColor()
      };
      allStruct.push(structResult)
    }  
      const backName=document.getElementById('Name')as HTMLInputElement;
      backName.innerHTML = allStruct[0].name;  

      const backPosition=document.getElementById("Position")as HTMLInputElement;
      backPosition.innerHTML = allStruct[0].position;
      var i:number
      let allProjectsForDeveloper!: string[]
    for(i=0;i<allStruct.length;i++){ 
      allProjectsForDeveloper.push(' '+allStruct[i].nameProject)
    }

      const backProject=document.getElementById("Project")as HTMLInputElement;
      [backProject.innerHTML] = allProjectsForDeveloper;
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
