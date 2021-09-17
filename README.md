# Custom Terraform Plugin
This is the minimal architecture for writing custom terraform plugins

# How To Reuse This

### Step 1
Clone this repo

### Step2 
With same name create your own github repo

### Step3 
In main.go on line # 6 change import link according to your username of github

### Step4
Remove git origin of this repo and add your repo as origin

### Step5
Create terraform account using github signup

### Step6
Create a gpg key on your local system and get its public and private keys


1) Add public gpg key to your terrafrom account

2) Add private gpg in secrets of your github repo with name `GPG_PRIVATE_KEY`

### Step7
commit code  
add tag as v1.0.0  
push code and tag to your repo

### Step8
Rerun your github action and you are done!

OR

Add a comment in your code and push it!
