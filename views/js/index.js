console.log("Hello World!");
let counter = 0;

const btn = document.getElementById("btn");
const cnt = document.getElementById("cnt")
btn.onclick = ()=>{
    counter += 1;
    console.log(counter);
    cnt.innerHTML = counter;
}