
kubectl run -it --rm --image=mysql:5.6 -n kubecon --restart=Never mysql-client -- mysql -h mysql -ppassword

CREATE DATABASE health;

use health;

create table steps(
   steps_id INT NOT NULL AUTO_INCREMENT,
   steps INT NOT NULL,
   name VARCHAR(400),
   activity_date DATE,
   PRIMARY KEY ( steps_id )
);

create table id(
   id INT NOT NULL AUTO_INCREMENT,
   name VARCHAR(400),
   activity_date DATE,
   PRIMARY KEY ( id )
);

create table activity(
   id INT NOT NULL AUTO_INCREMENT,
   activity_name VARCHAR(400) NOT NULL,
   activity_duration INT NOT NULL,
   name VARCHAR(400),
   activity_date DATE,
   PRIMARY KEY ( id )
);
