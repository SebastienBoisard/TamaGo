# MongoDB on TamaGo


To install the latest version of MongoDB (v3.6.2) on my Ubuntu, I followed the instructions of the tutorial 
["install MongoDB on Ubuntu"](https://docs.mongodb.com/manual/tutorial/install-mongodb-on-ubuntu/) from the official 
MongoDB.



1. Import the public key used by the package management system
 

```
sudo apt-key adv --keyserver hkp://keyserver.ubuntu.com:80 --recv 2930ADAE8CAF5059EE73BB4B58712A2291FA4AD5

```

2. Create a list file for MongoDB

```
echo "deb [ arch=amd64,arm64 ] https://repo.mongodb.org/apt/ubuntu xenial/mongodb-org/3.6 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-3.6.list
```

3. Reload local package database

```
sudo apt-get update
```


4. Install the MongoDB packages

```
sudo apt-get install -y mongodb-org
```

5. Start MongoDB

```
sudo service mongod start
```

6. Verify that MongoDB has started successfully

```
grep "waiting for connections on port 27017" /var/log/mongodb/mongod.log
```

7. Install MongoDB driver for Go (aka [mgo](https://labix.org/mgo))

```
go get gopkg.in/mgo.v2
```

8. Play with a mongo shell 

```
mongo --host 127.0.0.1:27017
```