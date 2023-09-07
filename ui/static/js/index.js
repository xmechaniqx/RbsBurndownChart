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
      data: [48,35,29,15],
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