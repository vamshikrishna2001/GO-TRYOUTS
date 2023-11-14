const fs = require('fs');


function print1() {
    fs.readFile('h.txt','utf-8',function (err,data){
        console.log(data)

    })

}

function print2() {
    console.log("2")
}

function print3(){
    setTimeout(()=>{
        console.log("3")
        print2()
        print1()
    },10)
}


// print1()
// print2()
print3()


