inp = document.querySelector('.text');
btn =  document.querySelectorAll('button')
    for (let i of btn) {
           i.disabled = true
        }
       
inp.oninput = function() {
if (inp.value.length > 0) {
    for (let i of btn) {
        i.disabled = false
        }
               
} else {
    for (let i of btn) {       
        i.disabled = true
        }
                
    }
}