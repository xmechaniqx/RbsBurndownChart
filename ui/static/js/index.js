const ctx = document.getElementById('myChart');
const date = new window.Date();  // 2009-11-10
const month = date.toLocaleString('default', { month: 'short' });
const day = date.toLocaleString('default', { day: "numeric" });
console.log(day)
var labels = [];

//how many days in month
function daysInMonth (month, year) {
    return new Date(year, month, 0).getDate();
}
for (i=0;i<daysInMonth(9,2023);i++) {
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
      label: 'RBS Burndown chart plan',
      data: [45,0],
      borderWidth: 2,
      borderColor: 'blue',
      backgroundColor: [
        'rgba(55, 119, 196, 1)'],
      tension:0,
      fill: false
    },{
        label: 'RBS Burndown chart',
        data: [12, 19, 3, 5, 2, 3],
        borderWidth: 2,
        borderColor: 'blue',
        backgroundColor: [
          'rgba(100, 50, 230, 1)'],
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