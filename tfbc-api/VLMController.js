var express = require('express');
var router = express.Router();
var bodyParser = require('body-parser');

router.use(bodyParser.urlencoded({ extended: true }));
router.use(bodyParser.json());

var VLM = require("./FabricHelper")


// Request LC
router.post('/addVehicle', function (req, res) {

VLM.addVehicle(req, res);

});

// Issue LC
router.post('/transferVehicle', function (req, res) {

    VLM.transferVehicle(req, res);
    
});

router.post('/getVehicle', function (req, res) {

    VLM.getVehicle(req, res);
    
});



module.exports = router;
