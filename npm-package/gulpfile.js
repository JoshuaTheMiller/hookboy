const fs = require("fs");
const gulp = require("gulp");
const gulpCopy = require("gulp-copy");

function tryMakeDir(dir) {
    if (!fs.existsSync(dir)){
        fs.mkdirSync(dir);
    }
}

gulp.task('build-go', function(done) {
    const exec = require('child_process').exec;
    
    tryMakeDir("./out/linux/");
    process.env.GOOS = "linux";
    exec('go build -o ./out/linux/ ../', function callback(error, stdout, stderr){
        done();
    });   

    tryMakeDir("./out/windows/");
    process.env.GOOS = "windows";
    exec('go build -o ./out/windows/ ../', function callback(error, stdout, stderr){
        done();
    });   
    
    tryMakeDir("./out/darwin/");
    process.env.GOOS = "darwin";
    exec('go build -o ./out/darwin/ ../', function callback(error, stdout, stderr){
        done();
    });  
});