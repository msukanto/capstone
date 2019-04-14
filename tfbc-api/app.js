var express = require('express');
var app = express();


var swaggerUi = require('swagger-ui-express');
var swaggerDocument = require('./swagger.json');


var VLMController = require('./VLMController');

app.use('/api-docs', swaggerUi.serve, swaggerUi.setup(swaggerDocument));
app.use('/vlm', VLMController);

module.exports = app;
