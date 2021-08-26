# Birthday Notify Bot
### **Premise:**
For an NGO I work for, we loved having the slackbot notify us of specific birthdays. Recently, that service ended and the creators have requested $2 per user per month for the bot. For our team of over 100+, that'd adds up to $2400 a year, which is about half the profit we make annually. I have a free alternative that is relatively easy to set up for less techy people that want a slack birthday bot 

### **Setup:**
- [ ]  Complete setup guide

### **Tech/Design choices:**
<details>
<summary>Ignore if you're uninterested in the tech</summary>
<br>
The purpose of the project is that it is not intended to scale because then it would go over the AWS free tier and would seize to be free.

**Golang**: I used and wanted to write a project in it because I've previously worked with it heavily in server side applications. The Go iteractions with Google cloud services are very fast. Also the many post requests made to send slack chats benefit from concurrenncy

**Google sheets**: Very easy to use and set up with a google form. Also, can be used as a database for birthdays and other public info in order

**AWS Lambda**: Serverless application so I don't have to worry about running an EC2 server and all the management that comes with that. It only needs to run once a day or when people submit extra messages to the user

**AWS DynamoDB**: The interactive messages allow a user to send a url/message to the birthday person and stores it in the DB. The key<>document database makes it easy to pull up the info for the specific birthday person and ssend the information. Ideally, would use AWS Keyspaces but the free trial is monthly for Keyspaces instead of the flat 25 GB for AWS DynamoDB

**Terraform**: Makes it easy to partition resources quickly for new users for the cloud services

</details>

### **TODO:**
#### Personal stuff for Kevin to organize and show what's currently accocmplished in the project
<details>
<summary>Project TODOs</summary>
<br>

**Admin stuff/documentation**
- [ ]  Complete ReadMe
- [x]  Remove my environment variables
- [ ]  Document environment variables
- [x]  Set up AWS Credentials
- [x]  Set up google cloud credenntials
- [ ]  Finish Google form template/Google sheets template
- [ ]  Robustness thorough and document errors page

**Code stuff**
- [x] Randoomize birthday messages, including multiple BDAY messages
- [x] Enable connection and reading to a google sheets as a database that can be used by non-coders
- [x] Implement fibonnaci heap for faster access and reduce search times
- [x] Set up slack notifications with auth token and an app
- [x] Use user email to search for slackID
- [x] Enable private messaging and pre-birthday private messages
- [x] Set up public messaging and the creatioon of a "Birthday" channel
- [x] Enable the @ feature to @ the special bday people, and also if the credentials are wrong, use a filler name 
- [x] Set up template for the slackbot (auth token and permissions) 
- [ ] Enable Google sheets writes
- [x] Use terraform to set up DB
- [ ] Use terraform to set up lambda/cloudwatchs
- [ ] Set up Docker to make it easily reproduceable
- [ ] Set up interactive slack messages
- [ ] Set up DynamoDB for interactive messages
</details>