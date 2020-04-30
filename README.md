# EmailSender

Create an application using Golang (using any database or in-memory storage) which sends out the Good Morning emails to all the users at 8 AM according to the different time zones of the users.

#Follow The Guidelines To setup the [EmailSender] Software

## Installation

###Clone the Repository

```bash
git clone https://github.com/mavensingh/emailsender
```

### Setup The database

```bash
mysql -u username -p
```

```bash
create database emailsender
```

```bash
use emailsender
```

Provide the directory location of database file

```bash
source '/home/tutree/Desktop/Assignment/mysql/emailsender.sql'
```

### Configure the Systemd File

Change [emailsender.service] file Configuration.

```bash
[Unit]
Description=emailsender

[Service]
Type=simple
Restart=always
RestartSec=5s
Environment=FROM="YOUR MAIL ID"
Environment=PASS="YOUR MAIL PASSWORD"
Environment=DBHOST="DATABASE HOST"
Environment=DBUSER="DATABASE USER "
Environment=DBPASS="DATABASE PASSWORD"
Environment=DBPORT="3306"
ExecStart=/home/tutree/Desktop/Assignment/app

[Install]
WantedBy=multi-user.target
```

### Open Terminal and Run the command

```bash
sudo cp emailsender.service ../../../../lib/systemd/system/
```

OR

```bash
sudo mv emailsender.service ../../../../lib/systemd/system/
```

### Open New Tab and Run the command

```bash
go build -o app
```

#### Then Run this Command to start emailsender as Service

```bash
systemctl start emailsender
```

#### Check The status Service Running or not

```bash
systemctl status emailsender
```

#### Note: Run [systemctl status emailsender] to check Service Running in background or not
