# BASIC KRIS KRINGLE BOT

```
usage: GMAIL_ADDR=<Gmail Address> GMAIL_PW=<GOOGLE APP PASSWORD> ./bot -t <template file location> -e <json document of emails, names>
```
* This application requires a Gmail Account with an App Password Configured. These are specified by the environment variables GMAIL_ADDR and GMAIL_PW
* The template.html file is a html email template. The code will replace occurences of {{from}} and {{to}}. 
* {{from}} is the person that will be buying the present
* {{to}} is the person that will be receiving a present.
* The email.json file stores a json array of people who will be participating in the kris kringle for example
```
[
    {
        "name": "Bob",
        "email": "bob@example.com"
    },
    {
        "name": "Alice",
        "email": "alice@example.com"
    },
    {
        "name": "Belle",
        "email": "belle@example.com"
    }
]
```
