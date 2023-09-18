
//renderChart() - функция формирует и выводит график на странице на основании полученной структуры
function renderChart(structsEntered:object):void {
    let dataset 
    let datasetsArr = [dataset]
    const ctx = document.getElementById('myChart');
    const date = new window.Date();  
    const month = date.toLocaleString('default', { month: 'short' });
    // const day = date.toLocaleString('default', { day: "numeric" });
    let labels! : string[];
    let struct ={
      name: "",  // под ключом "name" хранится значение "John"
      position: "",
      nameProject: "",
      burnDownDays: [0],
      sprintDate:"",
      color: ""
    } 
    //формируем массив дней месяца для актуального месяца
    let i : number
    for (i=1;i<31;i++) {
        labels.push(i+' '+month)
        } 

    for (i=0; i<Object.keys(structsEntered).length; i++) {
      struct=structsEntered[i]
    dataset = {
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