var express = require("express");
var app = express();
var port = 3000;

var bodyParser = require('body-parser');
app.use(bodyParser.json());
app.use(bodyParser.urlencoded({ extended: true }));

var mongoose = require("mongoose");
mongoose.Promise = global.Promise;
mongoose.connect("mongodb://localhost:27017/testingMongo");
var myData;

var studentDetails = new mongoose.Schema({
  name: String,
  email: String,
  phone: String
});

var User = mongoose.model("User", studentDetails);

app.get("/", (req, res) => {
  res.sendFile(__dirname + "/index.html");
});

app.post("/saveDetails", (req, res) => {
  myData = new User(req.body);
  myData.save()
    .then(item => {
      //res.send("item saved to database");
      res.sendFile(__dirname+"/data.html");
    })
    .catch(err => {
      res.status(400).send("unable to save to database");
    });
});

app.post("/viewDetails", (req, res) => {
	res.send("Name: "+myData.name+", Email: "+myData.email+", Phone: "+myData.phone);
});

app.listen(port, () => {
  console.log("Server listening on port " + port);
});