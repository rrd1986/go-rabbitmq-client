# go-rabbitmq-client 
A go client wrapper on top of github.com/streadway/amqp

# Run the unit test
```
$ make test
  >  locally deploying rabbitmq as docker cointainer.....
b300d8ecd8c48a627af2fc8618aa0d643210977bbd7d476f57344b4b7a53c3ed
  >  Running tests...
?       github.com/rrd1986/go-rabbitmq-client   [no test files]
ok      github.com/rrd1986/go-rabbitmq-client/client    1.069s  coverage: 37.7% of statements
ok      github.com/rrd1986/go-rabbitmq-client/connection        1.678s  coverage: 30.0% of statements
?       github.com/rrd1986/go-rabbitmq-client/logs      [no test files]
ok      github.com/rrd1986/go-rabbitmq-client/notification      0.653s  coverage: 46.7% of statements
ok      github.com/rrd1986/go-rabbitmq-client/producer  1.353s  coverage: 36.7% of statements
?       github.com/rrd1986/go-rabbitmq-client/utils     [no test files]
  >  removing locally deployed rabbitmq as docker cointainer.....
rabbitmq
```