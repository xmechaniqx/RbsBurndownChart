import {fnRequest} from "./response.js"
//Инициализируем параметр со значением текущей ссылки из адресной строки
const urlParams = new URLSearchParams(window.location.search);
//Получаем логин из полученного параметра
const login = urlParams.get('login');
//Формируем ссылку добавляя адрес хэндлера /chart
let url = 'http://'+window.location.host+'/chart?login='+login
let today = new Date().toISOString().slice(0, 10)
// console.log(today)
fnRequest(url)





